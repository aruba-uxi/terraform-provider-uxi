/*
Configuration Api

Nice description goes here

API version: 5.10.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SensorGroupAssignmentsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorGroupAssignmentsPostRequest{}

// SensorGroupAssignmentsPostRequest struct for SensorGroupAssignmentsPostRequest
type SensorGroupAssignmentsPostRequest struct {
	GroupId  string `json:"groupId"`
	SensorId string `json:"sensorId"`
}

type _SensorGroupAssignmentsPostRequest SensorGroupAssignmentsPostRequest

// NewSensorGroupAssignmentsPostRequest instantiates a new SensorGroupAssignmentsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorGroupAssignmentsPostRequest(
	groupId string,
	sensorId string,
) *SensorGroupAssignmentsPostRequest {
	this := SensorGroupAssignmentsPostRequest{}
	this.GroupId = groupId
	this.SensorId = sensorId
	return &this
}

// NewSensorGroupAssignmentsPostRequestWithDefaults instantiates a new SensorGroupAssignmentsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorGroupAssignmentsPostRequestWithDefaults() *SensorGroupAssignmentsPostRequest {
	this := SensorGroupAssignmentsPostRequest{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *SensorGroupAssignmentsPostRequest) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentsPostRequest) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *SensorGroupAssignmentsPostRequest) SetGroupId(v string) {
	o.GroupId = v
}

// GetSensorId returns the SensorId field value
func (o *SensorGroupAssignmentsPostRequest) GetSensorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SensorId
}

// GetSensorIdOk returns a tuple with the SensorId field value
// and a boolean to check if the value has been set.
func (o *SensorGroupAssignmentsPostRequest) GetSensorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SensorId, true
}

// SetSensorId sets field value
func (o *SensorGroupAssignmentsPostRequest) SetSensorId(v string) {
	o.SensorId = v
}

func (o SensorGroupAssignmentsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorGroupAssignmentsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	toSerialize["sensorId"] = o.SensorId
	return toSerialize, nil
}

func (o *SensorGroupAssignmentsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupId",
		"sensorId",
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

	varSensorGroupAssignmentsPostRequest := _SensorGroupAssignmentsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSensorGroupAssignmentsPostRequest)

	if err != nil {
		return err
	}

	*o = SensorGroupAssignmentsPostRequest(varSensorGroupAssignmentsPostRequest)

	return err
}

type NullableSensorGroupAssignmentsPostRequest struct {
	value *SensorGroupAssignmentsPostRequest
	isSet bool
}

func (v NullableSensorGroupAssignmentsPostRequest) Get() *SensorGroupAssignmentsPostRequest {
	return v.value
}

func (v *NullableSensorGroupAssignmentsPostRequest) Set(val *SensorGroupAssignmentsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorGroupAssignmentsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorGroupAssignmentsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorGroupAssignmentsPostRequest(
	val *SensorGroupAssignmentsPostRequest,
) *NullableSensorGroupAssignmentsPostRequest {
	return &NullableSensorGroupAssignmentsPostRequest{value: val, isSet: true}
}

func (v NullableSensorGroupAssignmentsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorGroupAssignmentsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
