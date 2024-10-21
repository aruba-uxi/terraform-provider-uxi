/*
Configuration Api

Nice description goes here

API version: 2.6.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NetworkGroupAssignmentsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkGroupAssignmentsPostRequest{}

// NetworkGroupAssignmentsPostRequest struct for NetworkGroupAssignmentsPostRequest
type NetworkGroupAssignmentsPostRequest struct {
	GroupId   string `json:"groupId"`
	NetworkId string `json:"networkId"`
}

type _NetworkGroupAssignmentsPostRequest NetworkGroupAssignmentsPostRequest

// NewNetworkGroupAssignmentsPostRequest instantiates a new NetworkGroupAssignmentsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkGroupAssignmentsPostRequest(groupId string, networkId string) *NetworkGroupAssignmentsPostRequest {
	this := NetworkGroupAssignmentsPostRequest{}
	this.GroupId = groupId
	this.NetworkId = networkId
	return &this
}

// NewNetworkGroupAssignmentsPostRequestWithDefaults instantiates a new NetworkGroupAssignmentsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkGroupAssignmentsPostRequestWithDefaults() *NetworkGroupAssignmentsPostRequest {
	this := NetworkGroupAssignmentsPostRequest{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *NetworkGroupAssignmentsPostRequest) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsPostRequest) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *NetworkGroupAssignmentsPostRequest) SetGroupId(v string) {
	o.GroupId = v
}

// GetNetworkId returns the NetworkId field value
func (o *NetworkGroupAssignmentsPostRequest) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *NetworkGroupAssignmentsPostRequest) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *NetworkGroupAssignmentsPostRequest) SetNetworkId(v string) {
	o.NetworkId = v
}

func (o NetworkGroupAssignmentsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkGroupAssignmentsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	toSerialize["networkId"] = o.NetworkId
	return toSerialize, nil
}

func (o *NetworkGroupAssignmentsPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varNetworkGroupAssignmentsPostRequest := _NetworkGroupAssignmentsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkGroupAssignmentsPostRequest)

	if err != nil {
		return err
	}

	*o = NetworkGroupAssignmentsPostRequest(varNetworkGroupAssignmentsPostRequest)

	return err
}

type NullableNetworkGroupAssignmentsPostRequest struct {
	value *NetworkGroupAssignmentsPostRequest
	isSet bool
}

func (v NullableNetworkGroupAssignmentsPostRequest) Get() *NetworkGroupAssignmentsPostRequest {
	return v.value
}

func (v *NullableNetworkGroupAssignmentsPostRequest) Set(val *NetworkGroupAssignmentsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkGroupAssignmentsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkGroupAssignmentsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkGroupAssignmentsPostRequest(val *NetworkGroupAssignmentsPostRequest) *NullableNetworkGroupAssignmentsPostRequest {
	return &NullableNetworkGroupAssignmentsPostRequest{value: val, isSet: true}
}

func (v NullableNetworkGroupAssignmentsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkGroupAssignmentsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
