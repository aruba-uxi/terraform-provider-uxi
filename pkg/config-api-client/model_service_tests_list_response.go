/*
Configuration Api

Nice description goes here

API version: 5.4.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ServiceTestsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTestsListResponse{}

// ServiceTestsListResponse struct for ServiceTestsListResponse
type ServiceTestsListResponse struct {
	Items []ServiceTestsListItem `json:"items"`
	Count int32                  `json:"count"`
	Next  NullableString         `json:"next"`
}

type _ServiceTestsListResponse ServiceTestsListResponse

// NewServiceTestsListResponse instantiates a new ServiceTestsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTestsListResponse(
	items []ServiceTestsListItem,
	count int32,
	next NullableString,
) *ServiceTestsListResponse {
	this := ServiceTestsListResponse{}
	this.Items = items
	this.Count = count
	this.Next = next
	return &this
}

// NewServiceTestsListResponseWithDefaults instantiates a new ServiceTestsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTestsListResponseWithDefaults() *ServiceTestsListResponse {
	this := ServiceTestsListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ServiceTestsListResponse) GetItems() []ServiceTestsListItem {
	if o == nil {
		var ret []ServiceTestsListItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ServiceTestsListResponse) GetItemsOk() ([]ServiceTestsListItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ServiceTestsListResponse) SetItems(v []ServiceTestsListItem) {
	o.Items = v
}

// GetCount returns the Count field value
func (o *ServiceTestsListResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ServiceTestsListResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ServiceTestsListResponse) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServiceTestsListResponse) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceTestsListResponse) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *ServiceTestsListResponse) SetNext(v string) {
	o.Next.Set(&v)
}

func (o ServiceTestsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTestsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["count"] = o.Count
	toSerialize["next"] = o.Next.Get()
	return toSerialize, nil
}

func (o *ServiceTestsListResponse) UnmarshalJSON(data []byte) (err error) {
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

	varServiceTestsListResponse := _ServiceTestsListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceTestsListResponse)

	if err != nil {
		return err
	}

	*o = ServiceTestsListResponse(varServiceTestsListResponse)

	return err
}

type NullableServiceTestsListResponse struct {
	value *ServiceTestsListResponse
	isSet bool
}

func (v NullableServiceTestsListResponse) Get() *ServiceTestsListResponse {
	return v.value
}

func (v *NullableServiceTestsListResponse) Set(val *ServiceTestsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTestsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTestsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTestsListResponse(
	val *ServiceTestsListResponse,
) *NullableServiceTestsListResponse {
	return &NullableServiceTestsListResponse{value: val, isSet: true}
}

func (v NullableServiceTestsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTestsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
