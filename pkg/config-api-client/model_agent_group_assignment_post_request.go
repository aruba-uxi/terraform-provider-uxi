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

// checks if the AgentGroupAssignmentPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentGroupAssignmentPostRequest{}

// AgentGroupAssignmentPostRequest struct for AgentGroupAssignmentPostRequest
type AgentGroupAssignmentPostRequest struct {
	GroupId string `json:"groupId"`
	AgentId string `json:"agentId"`
}

type _AgentGroupAssignmentPostRequest AgentGroupAssignmentPostRequest

// NewAgentGroupAssignmentPostRequest instantiates a new AgentGroupAssignmentPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentGroupAssignmentPostRequest(
	groupId string,
	agentId string,
) *AgentGroupAssignmentPostRequest {
	this := AgentGroupAssignmentPostRequest{}
	this.GroupId = groupId
	this.AgentId = agentId
	return &this
}

// NewAgentGroupAssignmentPostRequestWithDefaults instantiates a new AgentGroupAssignmentPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentGroupAssignmentPostRequestWithDefaults() *AgentGroupAssignmentPostRequest {
	this := AgentGroupAssignmentPostRequest{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *AgentGroupAssignmentPostRequest) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentPostRequest) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *AgentGroupAssignmentPostRequest) SetGroupId(v string) {
	o.GroupId = v
}

// GetAgentId returns the AgentId field value
func (o *AgentGroupAssignmentPostRequest) GetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentPostRequest) GetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *AgentGroupAssignmentPostRequest) SetAgentId(v string) {
	o.AgentId = v
}

func (o AgentGroupAssignmentPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentGroupAssignmentPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	toSerialize["agentId"] = o.AgentId
	return toSerialize, nil
}

func (o *AgentGroupAssignmentPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupId",
		"agentId",
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

	varAgentGroupAssignmentPostRequest := _AgentGroupAssignmentPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentGroupAssignmentPostRequest)

	if err != nil {
		return err
	}

	*o = AgentGroupAssignmentPostRequest(varAgentGroupAssignmentPostRequest)

	return err
}

type NullableAgentGroupAssignmentPostRequest struct {
	value *AgentGroupAssignmentPostRequest
	isSet bool
}

func (v NullableAgentGroupAssignmentPostRequest) Get() *AgentGroupAssignmentPostRequest {
	return v.value
}

func (v *NullableAgentGroupAssignmentPostRequest) Set(val *AgentGroupAssignmentPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentGroupAssignmentPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentGroupAssignmentPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentGroupAssignmentPostRequest(
	val *AgentGroupAssignmentPostRequest,
) *NullableAgentGroupAssignmentPostRequest {
	return &NullableAgentGroupAssignmentPostRequest{value: val, isSet: true}
}

func (v NullableAgentGroupAssignmentPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentGroupAssignmentPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
