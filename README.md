# Terraform Provider

Terraform Provider for interacting with GitHub App Installations

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.svg)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.10.x
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

## Building the Provider

```sh
$ make build
```

## Using the Provider

### Provider

#### Example Usage

```hcl-terraform
provider "gha" {
    app_id          = "1"
    installation_id = "1"
}
```

#### Argument Reference

The following arguments are supported in the `provider` block:

- `base_url` - (Optional) This is the target GitHub base API endpoint. It can also be sourced from the `GITHUB_BASE_URL`
    environment variable. This will default to `https://api.github.com/`

- `app_id` - (Required) The GitHub App's ID. It can also be sourced from the `GITHUB_APP_ID` environment variable.

- `installation_id` - (Required) The GitHub App Installation's ID. It can also be sourced from the `GITHUB_APP_INSTALLATION_ID` 
    environment variable.

- `pem` - (Required) The GitHub App private key. It can also be sourced from the `GITHUB_APP_PEM` environment variable
    if there is only a single GitHub App Provider being used.


### Data Sources

#### Token

##### Example Usage

```hcl-terraform
data "gha_token" "test_org" {}

provider "github" {
    organization = "test_org"
    base_url     = "https://github.service.anz/api/v3/"
    token        = data.gha_token.test_org.token
}
```
