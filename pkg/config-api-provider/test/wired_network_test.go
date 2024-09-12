package test

import (
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWiredNetworkResource(t *testing.T) {
	defer gock.Off()
	MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a wired_network is not allowed
			{
				Config: providerConfig + `
					resource "uxi_wired_network" "my_wired_network" {
						alias = "alias"
					}`,

				ExpectError: regexp.MustCompile(`(?s)creating a wired_network is not supported; wired_networks can only be\s*imported`),
			},
			// Importing a wired_network
			{
				PreConfig: func() {
					MockGetWiredNetwork(
						"uid",
						GenerateWiredNetworkPaginatedResponse(
							[]map[string]interface{}{GenerateWiredNetworkResponse("uid", "")},
						),
						2,
					)
				},
				Config: providerConfig + `
					resource "uxi_wired_network" "my_wired_network" {
						alias = "alias"
					}

					import {
						to = uxi_wired_network.my_wired_network
						id = "uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_wired_network.my_wired_network", "alias", "alias"),
					resource.TestCheckResourceAttr("uxi_wired_network.my_wired_network", "id", "uid"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					MockGetWiredNetwork(
						"uid",
						GenerateWiredNetworkPaginatedResponse(
							[]map[string]interface{}{GenerateWiredNetworkResponse("uid", "")},
						),
						1,
					)
				},
				ResourceName:      "uxi_wired_network.my_wired_network",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a wired_network is not allowed
			{
				PreConfig: func() {
					MockGetWiredNetwork(
						"uid",
						GenerateWiredNetworkPaginatedResponse(
							[]map[string]interface{}{GenerateWiredNetworkResponse("uid", "")},
						),
						1,
					)
				},
				Config: providerConfig + `
				resource "uxi_wired_network" "my_wired_network" {
					alias = "updated_alias"
				}`,
				ExpectError: regexp.MustCompile(`(?s)updating a wired_network is not supported; wired_networks can only be updated\s*through the dashboard`),
			},
			// Deleting a wired_network is not allowed
			{
				PreConfig: func() {
					MockGetWiredNetwork(
						"uid",
						GenerateWiredNetworkPaginatedResponse(
							[]map[string]interface{}{GenerateWiredNetworkResponse("uid", "")},
						),
						2,
					)
				},
				Config:      providerConfig + ``,
				ExpectError: regexp.MustCompile(`(?s)deleting a wired_network is not supported; wired_networks can only removed\s*from state`),
			},
			// Remove wired_network from state
			{
				Config: providerConfig + `
					removed {
						from = uxi_wired_network.my_wired_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
