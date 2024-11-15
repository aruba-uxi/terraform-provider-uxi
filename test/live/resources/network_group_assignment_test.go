/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNetworkGroupAssignmentResourceForWiredNetwork(t *testing.T) {
	const (
		groupName  = "tf_provider_acceptance_test_network_assignment_resource"
		group2Name = "tf_provider_acceptance_test_network_assignment_resource_two"
	)

	var (
		resourceIdBeforeRecreate string
		resourceIdAfterRecreate  string
	)

	// Test Wired Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_wired_network" "my_network" {
						name = "` + config.WiredNetworkName + `"
					}

					import {
						to = uxi_wired_network.my_network
						id = "` + config.WiredNetworkId + `"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wired_network.my_network.id
						group_id   = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WiredNetworkId,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							assert.Equal(t, group_id, util.GetGroupByName(groupName).Id)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_network_group_assignment.my_network_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIdBeforeRecreate = rs.Primary.ID
						return util.CheckStateAgainstNetworkGroupAssignment(
							t,
							"uxi_network_group_assignment.my_network_group_assignment",
							util.GetNetworkGroupAssignment(resourceIdBeforeRecreate),
						)(s)
					},
				),
			},
			// ImportState
			{
				ResourceName:      "uxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_wired_network" "my_network" {
						name = "` + config.WiredNetworkName + `"
					}

					import {
						to = uxi_wired_network.my_network
						id = "` + config.WiredNetworkId + `"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from network/group to network/group_2
					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = uxi_wired_network.my_network.id
						group_id 		 = uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WiredNetworkId,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							assert.Equal(t, group_id, util.GetGroupByName(group2Name).Id)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_network_group_assignment.my_network_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIdAfterRecreate = rs.Primary.ID
						return util.CheckStateAgainstNetworkGroupAssignment(
							t,
							"uxi_network_group_assignment.my_network_group_assignment",
							util.GetNetworkGroupAssignment(resourceIdAfterRecreate),
						)(s)
					},
					// Check that resource has been recreated
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"id",
						func(value string) error {
							assert.NotEqual(t, value, resourceIdBeforeRecreate)
							return nil
						},
					),
				),
			},
			// Delete network-group-assignments and remove networks from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wired_network.my_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			assert.Equal(t, util.GetGroupByName(groupName), nil)
			assert.Equal(t, util.GetGroupByName(group2Name), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIdBeforeRecreate), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIdAfterRecreate), nil)
			return nil
		},
	})
}

func TestNetworkGroupAssignmentResourceForWirelessNetwork(t *testing.T) {
	const (
		groupName  = "tf_provider_acceptance_test_network_assignment_test"
		group2Name = "tf_provider_acceptance_test_network_assignment_test_two"
	)

	var (
		resourceIdBeforeRecreate string
		resourceIdAfterRecreate  string
	)

	// Test Wired Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_wireless_network" "my_network" {
						name = "` + config.WirelessNetworkName + `"
					}

					import {
						to = uxi_wireless_network.my_network
						id = "` + config.WirelessNetworkId + `"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wireless_network.my_network.id
						group_id   = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WirelessNetworkId,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							assert.Equal(t, group_id, util.GetGroupByName(groupName).Id)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_network_group_assignment.my_network_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIdBeforeRecreate = rs.Primary.ID
						return util.CheckStateAgainstNetworkGroupAssignment(
							t,
							"uxi_network_group_assignment.my_network_group_assignment",
							util.GetNetworkGroupAssignment(resourceIdBeforeRecreate),
						)(s)
					},
				),
			},
			// ImportState
			{
				ResourceName:      "uxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_wireless_network" "my_network" {
						name = "` + config.WirelessNetworkName + `"
					}

					import {
						to = uxi_wireless_network.my_network
						id = "` + config.WirelessNetworkId + `"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from network/group to network/group_2
					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = uxi_wireless_network.my_network.id
						group_id 		 = uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WirelessNetworkId,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							assert.Equal(t, group_id, util.GetGroupByName(group2Name).Id)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_network_group_assignment.my_network_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIdAfterRecreate = rs.Primary.ID
						return util.CheckStateAgainstNetworkGroupAssignment(
							t,
							"uxi_network_group_assignment.my_network_group_assignment",
							util.GetNetworkGroupAssignment(resourceIdAfterRecreate),
						)(s)
					},
					// Check that resource has been recreated
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"id",
						func(value string) error {
							assert.NotEqual(t, value, resourceIdBeforeRecreate)
							return nil
						},
					),
				),
			},
			// Delete network-group-assignments and remove networks from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wireless_network.my_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			assert.Equal(t, util.GetGroupByName(groupName), nil)
			assert.Equal(t, util.GetGroupByName(group2Name), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIdBeforeRecreate), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIdAfterRecreate), nil)
			return nil
		},
	})
}
