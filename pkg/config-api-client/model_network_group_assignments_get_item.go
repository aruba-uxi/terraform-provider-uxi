/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
HPE Aruba Networking UXI Configuration

This document outlines the API contracts for HPE Aruba Networking UXI.

API version: 5.22.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NetworkGroupAssignmentsGetItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkGroupAssignmentsGetItem{}

// NetworkGroupAssignmentsGetItem struct for NetworkGroupAssignmentsGetItem
type NetworkGroupAssignmentsGetItem struct {
	// The unique identifier of the network group assignment
	Id      string                            `json:"id"`
	Group   NetworkGroupAssignmentsGetGroup   `json:"group"`
	Network NetworkGroupAssignmentsGetNetwork `json:"network"`
	// The type of the resource.
	Type string `json:"type"`
}

type _NetworkGroupAssignmentsGetItem NetworkGroupAssignmentsGetItem

// NewNetworkGroupAssignmentsGetItem instantiates a new NetworkGroupAssignmentsGetItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGroupAssignmentsGetItem(
	id string,
	group NetworkGroupAssignmentsGetGroup,
	network NetworkGroupAssignmentsGetNetwork,
	type_ string,
) *NetworkGroupAssignmentsGetItem {
	this := NetworkGroupAssignmentsGetItem{}
	this.Id = id
	this.Group = group
	this.Network = network
	this.Type = type_
	return &this
}

// NewNetworkGroupAssignmentsGetItemWithDefaults instantiates a new NetworkGroupAssignmentsGetItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGroupAssignmentsGetItemWithDefaults() *NetworkGroupAssignmentsGetItem {
	this := NetworkGroupAssignmentsGetItem{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkGroupAssignmentsGetItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsGetItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkGroupAssignmentsGetItem) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *NetworkGroupAssignmentsGetItem) GetGroup() NetworkGroupAssignmentsGetGroup {
	if o == nil {
		var ret NetworkGroupAssignmentsGetGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsGetItem) GetGroupOk() (*NetworkGroupAssignmentsGetGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *NetworkGroupAssignmentsGetItem) SetGroup(v NetworkGroupAssignmentsGetGroup) {
	o.Group = v
}

// GetNetwork returns the Network field value
func (o *NetworkGroupAssignmentsGetItem) GetNetwork() NetworkGroupAssignmentsGetNetwork {
	if o == nil {
		var ret NetworkGroupAssignmentsGetNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsGetItem) GetNetworkOk() (*NetworkGroupAssignmentsGetNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *NetworkGroupAssignmentsGetItem) SetNetwork(v NetworkGroupAssignmentsGetNetwork) {
	o.Network = v
}

// GetType returns the Type field value
func (o *NetworkGroupAssignmentsGetItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsGetItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkGroupAssignmentsGetItem) SetType(v string) {
	o.Type = v
}

func (o NetworkGroupAssignmentsGetItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkGroupAssignmentsGetItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["network"] = o.Network
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *NetworkGroupAssignmentsGetItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"network",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNetworkGroupAssignmentsGetItem := _NetworkGroupAssignmentsGetItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkGroupAssignmentsGetItem)

	if err != nil {
		return err
	}

	*o = NetworkGroupAssignmentsGetItem(varNetworkGroupAssignmentsGetItem)

	return err
}

type NullableNetworkGroupAssignmentsGetItem struct {
	value *NetworkGroupAssignmentsGetItem
	isSet bool
}

func (v NullableNetworkGroupAssignmentsGetItem) Get() *NetworkGroupAssignmentsGetItem {
	return v.value
}

func (v *NullableNetworkGroupAssignmentsGetItem) Set(val *NetworkGroupAssignmentsGetItem) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGroupAssignmentsGetItem) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGroupAssignmentsGetItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGroupAssignmentsGetItem(
	val *NetworkGroupAssignmentsGetItem,
) *NullableNetworkGroupAssignmentsGetItem {
	return &NullableNetworkGroupAssignmentsGetItem{value: val, isSet: true}
}

func (v NullableNetworkGroupAssignmentsGetItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGroupAssignmentsGetItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
