/*
Configuration Api

Nice description goes here

API version: 2.9.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ServiceTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTest{}

// ServiceTest struct for ServiceTest
type ServiceTest struct {
	Id string `json:"id"`
}

type _ServiceTest ServiceTest

// NewServiceTest instantiates a new ServiceTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTest(id string) *ServiceTest {
	this := ServiceTest{}
	this.Id = id
	return &this
}

// NewServiceTestWithDefaults instantiates a new ServiceTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestWithDefaults() *ServiceTest {
	this := ServiceTest{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceTest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceTest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceTest) SetId(v string) {
	o.Id = v
}

func (o ServiceTest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ServiceTest) UnmarshalJSON(data []byte) (err error) {
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

	varServiceTest := _ServiceTest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTest)

	if err != nil {
		return err
	}

	*o = ServiceTest(varServiceTest)

	return err
}

type NullableServiceTest struct {
	value *ServiceTest
	isSet bool
}

func (v NullableServiceTest) Get() *ServiceTest {
	return v.value
}

func (v *NullableServiceTest) Set(val *ServiceTest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTest(val *ServiceTest) *NullableServiceTest {
	return &NullableServiceTest{value: val, isSet: true}
}

func (v NullableServiceTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
