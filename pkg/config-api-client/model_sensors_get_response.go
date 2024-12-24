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

// checks if the SensorsGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorsGetResponse{}

// SensorsGetResponse struct for SensorsGetResponse
type SensorsGetResponse struct {
	// The list of resources.
	Items []SensorsGetItem `json:"items"`
	// The number of resources returned in the response.
	Count int32 `json:"count"`
	// The next cursor for pagination.
	Next NullableString `json:"next"`
}

type _SensorsGetResponse SensorsGetResponse

// NewSensorsGetResponse instantiates a new SensorsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorsGetResponse(
	items []SensorsGetItem,
	count int32,
	next NullableString,
) *SensorsGetResponse {
	this := SensorsGetResponse{}
	this.Items = items
	this.Count = count
	this.Next = next
	return &this
}

// NewSensorsGetResponseWithDefaults instantiates a new SensorsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorsGetResponseWithDefaults() *SensorsGetResponse {
	this := SensorsGetResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *SensorsGetResponse) GetItems() []SensorsGetItem {
	if o == nil {
		var ret []SensorsGetItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SensorsGetResponse) GetItemsOk() ([]SensorsGetItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SensorsGetResponse) SetItems(v []SensorsGetItem) {
	o.Items = v
}

// GetCount returns the Count field value
func (o *SensorsGetResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SensorsGetResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SensorsGetResponse) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SensorsGetResponse) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorsGetResponse) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *SensorsGetResponse) SetNext(v string) {
	o.Next.Set(&v)
}

func (o SensorsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorsGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	return toSerialize, nil
}

func (o *SensorsGetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"count",
		"next",
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

	varSensorsGetResponse := _SensorsGetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensorsGetResponse)

	if err != nil {
		return err
	}

	*o = SensorsGetResponse(varSensorsGetResponse)

	return err
}

type NullableSensorsGetResponse struct {
	value *SensorsGetResponse
	isSet bool
}

func (v NullableSensorsGetResponse) Get() *SensorsGetResponse {
	return v.value
}

func (v *NullableSensorsGetResponse) Set(val *SensorsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorsGetResponse(val *SensorsGetResponse) *NullableSensorsGetResponse {
	return &NullableSensorsGetResponse{value: val, isSet: true}
}

func (v NullableSensorsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
