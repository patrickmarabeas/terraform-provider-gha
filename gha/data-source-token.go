package gha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/hashicorp/terraform/helper/schema"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func dataSourceGhaToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGhaTokenRead,

		Schema: map[string]*schema.Schema{
			"token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceGhaTokenRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(ResourceProvider).BaseURL
	pem := meta.(ResourceProvider).Pem
	appID := meta.(ResourceProvider).AppID
	installID := meta.(ResourceProvider).InstallationID

	token, err := newToken(api, pem, appID, installID)
	if err != nil {
		return fmt.Errorf("error getting GitHub App Installation token: %w", err)
	}

	d.SetId(fmt.Sprintf("gha-token-%s-%s", appID, installID))
	d.Set("token", token)

	return nil
}

func newToken(api string, pem string, appId string, installId string) (string, error) {
	claims := jws.Claims{}
	claims.SetIssuedAt(time.Now())
	claims.SetExpiration(time.Now().Add(time.Duration(10) * time.Second))
	claims.SetIssuer(appId)

	pem = strings.ReplaceAll(pem, "\\n", "\n")
	rsaPrivate, err := crypto.ParseRSAPrivateKeyFromPEM([]byte(pem))
	if err != nil {
		return "", fmt.Errorf("error parsing RSA private key from PEM string: %w", err)
	}

	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)

	bearer, err := jwt.Serialize(rsaPrivate)
	if err != nil {
		return "", fmt.Errorf("error serializing JWT: %w", err)
	}

	url := fmt.Sprintf("%sapp/installations/%s/access_tokens", api, installId)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("error posting request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearer))
	req.Header.Set("Accept", "application/vnd.github.machine-man-preview+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error getting response: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return "", fmt.Errorf("status code returned (%d) is not 201", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	bodyString := string(bodyBytes)
	res := TokenResponse{}
	err = json.Unmarshal([]byte(bodyString), &res)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return res.Token, nil
}
