package test

import (
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	providerConfig = `provider "uxi" {
		host		  = "test.api.capenetworks.com"
		client_id     = "client_id"
		client_secret = "client_secret"
	}`
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"uxi": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)
