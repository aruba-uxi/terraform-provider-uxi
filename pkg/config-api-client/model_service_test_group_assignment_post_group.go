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

// checks if the ServiceTestGroupAssignmentPostGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTestGroupAssignmentPostGroup{}

// ServiceTestGroupAssignmentPostGroup struct for ServiceTestGroupAssignmentPostGroup
type ServiceTestGroupAssignmentPostGroup struct {
	// The unique identifier of the group
	Id string `json:"id"`
}

type _ServiceTestGroupAssignmentPostGroup ServiceTestGroupAssignmentPostGroup

// NewServiceTestGroupAssignmentPostGroup instantiates a new ServiceTestGroupAssignmentPostGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTestGroupAssignmentPostGroup(id string) *ServiceTestGroupAssignmentPostGroup {
	this := ServiceTestGroupAssignmentPostGroup{}
	this.Id = id
	return &this
}

// NewServiceTestGroupAssignmentPostGroupWithDefaults instantiates a new ServiceTestGroupAssignmentPostGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestGroupAssignmentPostGroupWithDefaults() *ServiceTestGroupAssignmentPostGroup {
	this := ServiceTestGroupAssignmentPostGroup{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceTestGroupAssignmentPostGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentPostGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceTestGroupAssignmentPostGroup) SetId(v string) {
	o.Id = v
}

func (o ServiceTestGroupAssignmentPostGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTestGroupAssignmentPostGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ServiceTestGroupAssignmentPostGroup) UnmarshalJSON(data []byte) (err error) {
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

	varServiceTestGroupAssignmentPostGroup := _ServiceTestGroupAssignmentPostGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTestGroupAssignmentPostGroup)

	if err != nil {
		return err
	}

	*o = ServiceTestGroupAssignmentPostGroup(varServiceTestGroupAssignmentPostGroup)

	return err
}

type NullableServiceTestGroupAssignmentPostGroup struct {
	value *ServiceTestGroupAssignmentPostGroup
	isSet bool
}

func (v NullableServiceTestGroupAssignmentPostGroup) Get() *ServiceTestGroupAssignmentPostGroup {
	return v.value
}

func (v *NullableServiceTestGroupAssignmentPostGroup) Set(
	val *ServiceTestGroupAssignmentPostGroup,
) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTestGroupAssignmentPostGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTestGroupAssignmentPostGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTestGroupAssignmentPostGroup(
	val *ServiceTestGroupAssignmentPostGroup,
) *NullableServiceTestGroupAssignmentPostGroup {
	return &NullableServiceTestGroupAssignmentPostGroup{value: val, isSet: true}
}

func (v NullableServiceTestGroupAssignmentPostGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTestGroupAssignmentPostGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
