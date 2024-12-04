/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

/*
HPE Aruba Networking UXI Configuration

This document outlines the API contracts for HPE Aruba Networking UXI.

API version: 5.19.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NetworkGroupAssignmentPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkGroupAssignmentPostRequest{}

// NetworkGroupAssignmentPostRequest struct for NetworkGroupAssignmentPostRequest
type NetworkGroupAssignmentPostRequest struct {
	GroupId   string `json:"groupId"`
	NetworkId string `json:"networkId"`
}

type _NetworkGroupAssignmentPostRequest NetworkGroupAssignmentPostRequest

// NewNetworkGroupAssignmentPostRequest instantiates a new NetworkGroupAssignmentPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGroupAssignmentPostRequest(
	groupId string,
	networkId string,
) *NetworkGroupAssignmentPostRequest {
	this := NetworkGroupAssignmentPostRequest{}
	this.GroupId = groupId
	this.NetworkId = networkId
	return &this
}

// NewNetworkGroupAssignmentPostRequestWithDefaults instantiates a new NetworkGroupAssignmentPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGroupAssignmentPostRequestWithDefaults() *NetworkGroupAssignmentPostRequest {
	this := NetworkGroupAssignmentPostRequest{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *NetworkGroupAssignmentPostRequest) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentPostRequest) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *NetworkGroupAssignmentPostRequest) SetGroupId(v string) {
	o.GroupId = v
}

// GetNetworkId returns the NetworkId field value
func (o *NetworkGroupAssignmentPostRequest) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentPostRequest) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *NetworkGroupAssignmentPostRequest) SetNetworkId(v string) {
	o.NetworkId = v
}

func (o NetworkGroupAssignmentPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkGroupAssignmentPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	toSerialize["networkId"] = o.NetworkId
	return toSerialize, nil
}

func (o *NetworkGroupAssignmentPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupId",
		"networkId",
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

	varNetworkGroupAssignmentPostRequest := _NetworkGroupAssignmentPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkGroupAssignmentPostRequest)

	if err != nil {
		return err
	}

	*o = NetworkGroupAssignmentPostRequest(varNetworkGroupAssignmentPostRequest)

	return err
}

type NullableNetworkGroupAssignmentPostRequest struct {
	value *NetworkGroupAssignmentPostRequest
	isSet bool
}

func (v NullableNetworkGroupAssignmentPostRequest) Get() *NetworkGroupAssignmentPostRequest {
	return v.value
}

func (v *NullableNetworkGroupAssignmentPostRequest) Set(val *NetworkGroupAssignmentPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGroupAssignmentPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGroupAssignmentPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGroupAssignmentPostRequest(
	val *NetworkGroupAssignmentPostRequest,
) *NullableNetworkGroupAssignmentPostRequest {
	return &NullableNetworkGroupAssignmentPostRequest{value: val, isSet: true}
}

func (v NullableNetworkGroupAssignmentPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGroupAssignmentPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
