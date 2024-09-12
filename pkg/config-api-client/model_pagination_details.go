/*
Configuration Api

Nice description goes here

API version: 1.11.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PaginationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationDetails{}

// PaginationDetails struct for PaginationDetails
type PaginationDetails struct {
	Limit    int32          `json:"limit"`
	Next     NullableString `json:"next"`
	Previous NullableString `json:"previous"`
	First    NullableString `json:"first"`
	Last     NullableString `json:"last"`
}

type _PaginationDetails PaginationDetails

// NewPaginationDetails instantiates a new PaginationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationDetails(limit int32, next NullableString, previous NullableString, first NullableString, last NullableString) *PaginationDetails {
	this := PaginationDetails{}
	this.Limit = limit
	this.Next = next
	this.Previous = previous
	this.First = first
	this.Last = last
	return &this
}

// NewPaginationDetailsWithDefaults instantiates a new PaginationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationDetailsWithDefaults() *PaginationDetails {
	this := PaginationDetails{}
	return &this
}

// GetLimit returns the Limit field value
func (o *PaginationDetails) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PaginationDetails) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PaginationDetails) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationDetails) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationDetails) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *PaginationDetails) SetNext(v string) {
	o.Next.Set(&v)
}

// GetPrevious returns the Previous field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationDetails) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}

	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationDetails) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// SetPrevious sets field value
func (o *PaginationDetails) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// GetFirst returns the First field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationDetails) GetFirst() string {
	if o == nil || o.First.Get() == nil {
		var ret string
		return ret
	}

	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationDetails) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// SetFirst sets field value
func (o *PaginationDetails) SetFirst(v string) {
	o.First.Set(&v)
}

// GetLast returns the Last field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginationDetails) GetLast() string {
	if o == nil || o.Last.Get() == nil {
		var ret string
		return ret
	}

	return *o.Last.Get()
}

// GetLastOk returns a tuple with the Last field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationDetails) GetLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Last.Get(), o.Last.IsSet()
}

// SetLast sets field value
func (o *PaginationDetails) SetLast(v string) {
	o.Last.Set(&v)
}

func (o PaginationDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["next"] = o.Next.Get()
	toSerialize["previous"] = o.Previous.Get()
	toSerialize["first"] = o.First.Get()
	toSerialize["last"] = o.Last.Get()
	return toSerialize, nil
}

func (o *PaginationDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limit",
		"next",
		"previous",
		"first",
		"last",
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

	varPaginationDetails := _PaginationDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginationDetails)

	if err != nil {
		return err
	}

	*o = PaginationDetails(varPaginationDetails)

	return err
}

type NullablePaginationDetails struct {
	value *PaginationDetails
	isSet bool
}

func (v NullablePaginationDetails) Get() *PaginationDetails {
	return v.value
}

func (v *NullablePaginationDetails) Set(val *PaginationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationDetails(val *PaginationDetails) *NullablePaginationDetails {
	return &NullablePaginationDetails{value: val, isSet: true}
}

func (v NullablePaginationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
