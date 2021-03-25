package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-provider-jerakia/jerakia"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: jerakia.Provider,
	})
}
