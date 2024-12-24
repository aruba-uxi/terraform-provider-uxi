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

// checks if the SensorGroupAssignmentPostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorGroupAssignmentPostResponse{}

// SensorGroupAssignmentPostResponse struct for SensorGroupAssignmentPostResponse
type SensorGroupAssignmentPostResponse struct {
	// The unique identifier of the sensor group assignment
	Id     string                          `json:"id"`
	Group  SensorGroupAssignmentPostGroup  `json:"group"`
	Sensor SensorGroupAssignmentPostSensor `json:"sensor"`
	// The type of the resource.
	Type string `json:"type"`
}

type _SensorGroupAssignmentPostResponse SensorGroupAssignmentPostResponse

// NewSensorGroupAssignmentPostResponse instantiates a new SensorGroupAssignmentPostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorGroupAssignmentPostResponse(
	id string,
	group SensorGroupAssignmentPostGroup,
	sensor SensorGroupAssignmentPostSensor,
	type_ string,
) *SensorGroupAssignmentPostResponse {
	this := SensorGroupAssignmentPostResponse{}
	this.Id = id
	this.Group = group
	this.Sensor = sensor
	this.Type = type_
	return &this
}

// NewSensorGroupAssignmentPostResponseWithDefaults instantiates a new SensorGroupAssignmentPostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorGroupAssignmentPostResponseWithDefaults() *SensorGroupAssignmentPostResponse {
	this := SensorGroupAssignmentPostResponse{}
	return &this
}

// GetId returns the Id field value
func (o *SensorGroupAssignmentPostResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentPostResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SensorGroupAssignmentPostResponse) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *SensorGroupAssignmentPostResponse) GetGroup() SensorGroupAssignmentPostGroup {
	if o == nil {
		var ret SensorGroupAssignmentPostGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentPostResponse) GetGroupOk() (*SensorGroupAssignmentPostGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SensorGroupAssignmentPostResponse) SetGroup(v SensorGroupAssignmentPostGroup) {
	o.Group = v
}

// GetSensor returns the Sensor field value
func (o *SensorGroupAssignmentPostResponse) GetSensor() SensorGroupAssignmentPostSensor {
	if o == nil {
		var ret SensorGroupAssignmentPostSensor
		return ret
	}

	return o.Sensor
}

// GetSensorOk returns a tuple with the Sensor field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentPostResponse) GetSensorOk() (*SensorGroupAssignmentPostSensor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sensor, true
}

// SetSensor sets field value
func (o *SensorGroupAssignmentPostResponse) SetSensor(v SensorGroupAssignmentPostSensor) {
	o.Sensor = v
}

// GetType returns the Type field value
func (o *SensorGroupAssignmentPostResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentPostResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SensorGroupAssignmentPostResponse) SetType(v string) {
	o.Type = v
}

func (o SensorGroupAssignmentPostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorGroupAssignmentPostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["sensor"] = o.Sensor
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *SensorGroupAssignmentPostResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"sensor",
		"type",
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

	varSensorGroupAssignmentPostResponse := _SensorGroupAssignmentPostResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensorGroupAssignmentPostResponse)

	if err != nil {
		return err
	}

	*o = SensorGroupAssignmentPostResponse(varSensorGroupAssignmentPostResponse)

	return err
}

type NullableSensorGroupAssignmentPostResponse struct {
	value *SensorGroupAssignmentPostResponse
	isSet bool
}

func (v NullableSensorGroupAssignmentPostResponse) Get() *SensorGroupAssignmentPostResponse {
	return v.value
}

func (v *NullableSensorGroupAssignmentPostResponse) Set(val *SensorGroupAssignmentPostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorGroupAssignmentPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorGroupAssignmentPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorGroupAssignmentPostResponse(
	val *SensorGroupAssignmentPostResponse,
) *NullableSensorGroupAssignmentPostResponse {
	return &NullableSensorGroupAssignmentPostResponse{value: val, isSet: true}
}

func (v NullableSensorGroupAssignmentPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorGroupAssignmentPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
