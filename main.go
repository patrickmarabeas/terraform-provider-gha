package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/patrickmarabeas/terraform-provider-gha/gha"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: gha.Provider,
	})
}
