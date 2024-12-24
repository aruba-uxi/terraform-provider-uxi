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

// checks if the AgentGroupAssignmentPostAgent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentGroupAssignmentPostAgent{}

// AgentGroupAssignmentPostAgent struct for AgentGroupAssignmentPostAgent
type AgentGroupAssignmentPostAgent struct {
	// The unique identifier of the agent
	Id string `json:"id"`
}

type _AgentGroupAssignmentPostAgent AgentGroupAssignmentPostAgent

// NewAgentGroupAssignmentPostAgent instantiates a new AgentGroupAssignmentPostAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentGroupAssignmentPostAgent(id string) *AgentGroupAssignmentPostAgent {
	this := AgentGroupAssignmentPostAgent{}
	this.Id = id
	return &this
}

// NewAgentGroupAssignmentPostAgentWithDefaults instantiates a new AgentGroupAssignmentPostAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentGroupAssignmentPostAgentWithDefaults() *AgentGroupAssignmentPostAgent {
	this := AgentGroupAssignmentPostAgent{}
	return &this
}

// GetId returns the Id field value
func (o *AgentGroupAssignmentPostAgent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentGroupAssignmentPostAgent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentGroupAssignmentPostAgent) SetId(v string) {
	o.Id = v
}

func (o AgentGroupAssignmentPostAgent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentGroupAssignmentPostAgent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AgentGroupAssignmentPostAgent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAgentGroupAssignmentPostAgent := _AgentGroupAssignmentPostAgent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentGroupAssignmentPostAgent)

	if err != nil {
		return err
	}

	*o = AgentGroupAssignmentPostAgent(varAgentGroupAssignmentPostAgent)

	return err
}

type NullableAgentGroupAssignmentPostAgent struct {
	value *AgentGroupAssignmentPostAgent
	isSet bool
}

func (v NullableAgentGroupAssignmentPostAgent) Get() *AgentGroupAssignmentPostAgent {
	return v.value
}

func (v *NullableAgentGroupAssignmentPostAgent) Set(val *AgentGroupAssignmentPostAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentGroupAssignmentPostAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentGroupAssignmentPostAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentGroupAssignmentPostAgent(
	val *AgentGroupAssignmentPostAgent,
) *NullableAgentGroupAssignmentPostAgent {
	return &NullableAgentGroupAssignmentPostAgent{value: val, isSet: true}
}

func (v NullableAgentGroupAssignmentPostAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentGroupAssignmentPostAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
