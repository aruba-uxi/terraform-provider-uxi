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

// checks if the ServiceTestGroupAssignmentsGetServiceTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTestGroupAssignmentsGetServiceTest{}

// ServiceTestGroupAssignmentsGetServiceTest struct for ServiceTestGroupAssignmentsGetServiceTest
type ServiceTestGroupAssignmentsGetServiceTest struct {
	// The unique identifier of the service test
	Id string `json:"id"`
}

type _ServiceTestGroupAssignmentsGetServiceTest ServiceTestGroupAssignmentsGetServiceTest

// NewServiceTestGroupAssignmentsGetServiceTest instantiates a new ServiceTestGroupAssignmentsGetServiceTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTestGroupAssignmentsGetServiceTest(
	id string,
) *ServiceTestGroupAssignmentsGetServiceTest {
	this := ServiceTestGroupAssignmentsGetServiceTest{}
	this.Id = id
	return &this
}

// NewServiceTestGroupAssignmentsGetServiceTestWithDefaults instantiates a new ServiceTestGroupAssignmentsGetServiceTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestGroupAssignmentsGetServiceTestWithDefaults() *ServiceTestGroupAssignmentsGetServiceTest {
	this := ServiceTestGroupAssignmentsGetServiceTest{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceTestGroupAssignmentsGetServiceTest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceTestGroupAssignmentsGetServiceTest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceTestGroupAssignmentsGetServiceTest) SetId(v string) {
	o.Id = v
}

func (o ServiceTestGroupAssignmentsGetServiceTest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTestGroupAssignmentsGetServiceTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ServiceTestGroupAssignmentsGetServiceTest) UnmarshalJSON(data []byte) (err error) {
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

	varServiceTestGroupAssignmentsGetServiceTest := _ServiceTestGroupAssignmentsGetServiceTest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTestGroupAssignmentsGetServiceTest)

	if err != nil {
		return err
	}

	*o = ServiceTestGroupAssignmentsGetServiceTest(varServiceTestGroupAssignmentsGetServiceTest)

	return err
}

type NullableServiceTestGroupAssignmentsGetServiceTest struct {
	value *ServiceTestGroupAssignmentsGetServiceTest
	isSet bool
}

func (v NullableServiceTestGroupAssignmentsGetServiceTest) Get() *ServiceTestGroupAssignmentsGetServiceTest {
	return v.value
}

func (v *NullableServiceTestGroupAssignmentsGetServiceTest) Set(
	val *ServiceTestGroupAssignmentsGetServiceTest,
) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTestGroupAssignmentsGetServiceTest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTestGroupAssignmentsGetServiceTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTestGroupAssignmentsGetServiceTest(
	val *ServiceTestGroupAssignmentsGetServiceTest,
) *NullableServiceTestGroupAssignmentsGetServiceTest {
	return &NullableServiceTestGroupAssignmentsGetServiceTest{value: val, isSet: true}
}

func (v NullableServiceTestGroupAssignmentsGetServiceTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTestGroupAssignmentsGetServiceTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
