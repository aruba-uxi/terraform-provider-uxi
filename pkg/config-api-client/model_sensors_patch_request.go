/*
Configuration Api

Nice description goes here

API version: 5.3.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"encoding/json"
)

// checks if the SensorsPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SensorsPatchRequest{}

// SensorsPatchRequest struct for SensorsPatchRequest
type SensorsPatchRequest struct {
	Name        *string        `json:"name,omitempty"`
	AddressNote NullableString `json:"address_note,omitempty"`
	Notes       NullableString `json:"notes,omitempty"`
	PcapMode    *string        `json:"pcap_mode,omitempty"`
}

// NewSensorsPatchRequest instantiates a new SensorsPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensorsPatchRequest() *SensorsPatchRequest {
	this := SensorsPatchRequest{}
	return &this
}

// NewSensorsPatchRequestWithDefaults instantiates a new SensorsPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensorsPatchRequestWithDefaults() *SensorsPatchRequest {
	this := SensorsPatchRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SensorsPatchRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensorsPatchRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SensorsPatchRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SensorsPatchRequest) SetName(v string) {
	o.Name = &v
}

// GetAddressNote returns the AddressNote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SensorsPatchRequest) GetAddressNote() string {
	if o == nil || IsNil(o.AddressNote.Get()) {
		var ret string
		return ret
	}
	return *o.AddressNote.Get()
}

// GetAddressNoteOk returns a tuple with the AddressNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorsPatchRequest) GetAddressNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressNote.Get(), o.AddressNote.IsSet()
}

// HasAddressNote returns a boolean if a field has been set.
func (o *SensorsPatchRequest) HasAddressNote() bool {
	if o != nil && o.AddressNote.IsSet() {
		return true
	}

	return false
}

// SetAddressNote gets a reference to the given NullableString and assigns it to the AddressNote field.
func (o *SensorsPatchRequest) SetAddressNote(v string) {
	o.AddressNote.Set(&v)
}

// SetAddressNoteNil sets the value for AddressNote to be an explicit nil
func (o *SensorsPatchRequest) SetAddressNoteNil() {
	o.AddressNote.Set(nil)
}

// UnsetAddressNote ensures that no value is present for AddressNote, not even an explicit nil
func (o *SensorsPatchRequest) UnsetAddressNote() {
	o.AddressNote.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SensorsPatchRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SensorsPatchRequest) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *SensorsPatchRequest) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *SensorsPatchRequest) SetNotes(v string) {
	o.Notes.Set(&v)
}

// SetNotesNil sets the value for Notes to be an explicit nil
func (o *SensorsPatchRequest) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *SensorsPatchRequest) UnsetNotes() {
	o.Notes.Unset()
}

// GetPcapMode returns the PcapMode field value if set, zero value otherwise.
func (o *SensorsPatchRequest) GetPcapMode() string {
	if o == nil || IsNil(o.PcapMode) {
		var ret string
		return ret
	}
	return *o.PcapMode
}

// GetPcapModeOk returns a tuple with the PcapMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensorsPatchRequest) GetPcapModeOk() (*string, bool) {
	if o == nil || IsNil(o.PcapMode) {
		return nil, false
	}
	return o.PcapMode, true
}

// HasPcapMode returns a boolean if a field has been set.
func (o *SensorsPatchRequest) HasPcapMode() bool {
	if o != nil && !IsNil(o.PcapMode) {
		return true
	}

	return false
}

// SetPcapMode gets a reference to the given string and assigns it to the PcapMode field.
func (o *SensorsPatchRequest) SetPcapMode(v string) {
	o.PcapMode = &v
}

func (o SensorsPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SensorsPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.AddressNote.IsSet() {
		toSerialize["address_note"] = o.AddressNote.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if !IsNil(o.PcapMode) {
		toSerialize["pcap_mode"] = o.PcapMode
	}
	return toSerialize, nil
}

type NullableSensorsPatchRequest struct {
	value *SensorsPatchRequest
	isSet bool
}

func (v NullableSensorsPatchRequest) Get() *SensorsPatchRequest {
	return v.value
}

func (v *NullableSensorsPatchRequest) Set(val *SensorsPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorsPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorsPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorsPatchRequest(val *SensorsPatchRequest) *NullableSensorsPatchRequest {
	return &NullableSensorsPatchRequest{value: val, isSet: true}
}

func (v NullableSensorsPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorsPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
