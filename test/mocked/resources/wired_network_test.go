/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func TestWiredNetworkResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	wiredNetwork := util.GenerateWiredNetworkResponse("id", "").Items[0]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a wired_network is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_wired_network" "my_wired_network" {
						name = "name"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)creating a wired_network is not supported; wired_networks can only be\s*imported`,
				),
			},
			// Importing a wired_network
			{
				PreConfig: func() {
					util.MockGetWiredNetwork("id", util.GenerateWiredNetworkResponse("id", ""), 2)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_wired_network" "my_wired_network" {
						name = "name"
					}

					import {
						to = hpeuxi_wired_network.my_wired_network
						id = "id"
					}`,

				Check: shared.CheckStateAgainstWiredNetwork(
					t,
					"hpeuxi_wired_network.my_wired_network",
					wiredNetwork,
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetWiredNetwork("id", util.GenerateWiredNetworkResponse("id", ""), 1)
				},
				ResourceName:      "hpeuxi_wired_network.my_wired_network",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a wired_network is not allowed
			{
				PreConfig: func() {
					util.MockGetWiredNetwork("id", util.GenerateWiredNetworkResponse("id", ""), 1)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_wired_network" "my_wired_network" {
					name = "updated_name"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)updating a wired_network is not supported; wired_networks can only be updated\s*through the dashboard`,
				),
			},
			// Deleting a wired_network is not allowed
			{
				PreConfig: func() {
					util.MockGetWiredNetwork("id", util.GenerateWiredNetworkResponse("id", ""), 2)
				},
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`(?s)deleting a wired_network is not supported; wired_networks can only removed\s*from state`,
				),
			},
			// Remove wired_network from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_wired_network.my_wired_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestWiredNetworkResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	wiredNetwork := util.GenerateWiredNetworkResponse("id", "").Items[0]
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Read
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Get(shared.WiredNetworkPath).
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetWiredNetwork("id", util.GenerateWiredNetworkResponse("id", ""), 2)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_wired_network" "my_wired_network" {
						name = "name"
					}

					import {
						to = hpeuxi_wired_network.my_wired_network
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					shared.CheckStateAgainstWiredNetwork(
						t,
						"hpeuxi_wired_network.my_wired_network",
						wiredNetwork,
					),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Cleanup
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_wired_network.my_wired_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestWiredNetworkResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					util.MockGetWiredNetwork("id", util.EmptyGetListResponse, 1)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_wired_network" "my_wired_network" {
						name = "name"
					}

					import {
						to = hpeuxi_wired_network.my_wired_network
						id = "id"
					}`,
				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
						Get(shared.WiredNetworkPath).
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_wired_network" "my_wired_network" {
						name = "name"
					}

					import {
						to = hpeuxi_wired_network.my_wired_network
						id = "id"
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
