/*
Configuration Api

Nice description goes here

API version: 5.10.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"encoding/json"
)

// checks if the AgentsPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentsPatchRequest{}

// AgentsPatchRequest Request body for patching an agent.  Fields:     name: Optional string     notes: Optional string     pcap_mode: Optional string
type AgentsPatchRequest struct {
	Name     *string `json:"name,omitempty"`
	Notes    *string `json:"notes,omitempty"`
	PcapMode *string `json:"pcapMode,omitempty"`
}

// NewAgentsPatchRequest instantiates a new AgentsPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentsPatchRequest() *AgentsPatchRequest {
	this := AgentsPatchRequest{}
	return &this
}

// NewAgentsPatchRequestWithDefaults instantiates a new AgentsPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentsPatchRequestWithDefaults() *AgentsPatchRequest {
	this := AgentsPatchRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgentsPatchRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentsPatchRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgentsPatchRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgentsPatchRequest) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AgentsPatchRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentsPatchRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AgentsPatchRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AgentsPatchRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetPcapMode returns the PcapMode field value if set, zero value otherwise.
func (o *AgentsPatchRequest) GetPcapMode() string {
	if o == nil || IsNil(o.PcapMode) {
		var ret string
		return ret
	}
	return *o.PcapMode
}

// GetPcapModeOk returns a tuple with the PcapMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentsPatchRequest) GetPcapModeOk() (*string, bool) {
	if o == nil || IsNil(o.PcapMode) {
		return nil, false
	}
	return o.PcapMode, true
}

// HasPcapMode returns a boolean if a field has been set.
func (o *AgentsPatchRequest) HasPcapMode() bool {
	if o != nil && !IsNil(o.PcapMode) {
		return true
	}

	return false
}

// SetPcapMode gets a reference to the given string and assigns it to the PcapMode field.
func (o *AgentsPatchRequest) SetPcapMode(v string) {
	o.PcapMode = &v
}

func (o AgentsPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentsPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.PcapMode) {
		toSerialize["pcapMode"] = o.PcapMode
	}
	return toSerialize, nil
}

type NullableAgentsPatchRequest struct {
	value *AgentsPatchRequest
	isSet bool
}

func (v NullableAgentsPatchRequest) Get() *AgentsPatchRequest {
	return v.value
}

func (v *NullableAgentsPatchRequest) Set(val *AgentsPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentsPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentsPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentsPatchRequest(val *AgentsPatchRequest) *NullableAgentsPatchRequest {
	return &NullableAgentsPatchRequest{value: val, isSet: true}
}

func (v NullableAgentsPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentsPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
