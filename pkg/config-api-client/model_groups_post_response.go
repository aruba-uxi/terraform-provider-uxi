/*
Configuration Api

Nice description goes here

API version: 1.5.0
Contact: support@capenetworks.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package config_api_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GroupsPostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsPostResponse{}

// GroupsPostResponse struct for GroupsPostResponse
type GroupsPostResponse struct {
	ParentUid string `json:"parent_uid"`
	Name      string `json:"name"`
	Uid       string `json:"uid"`
	Path      string `json:"path"`
}

type _GroupsPostResponse GroupsPostResponse

// NewGroupsPostResponse instantiates a new GroupsPostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsPostResponse(parentUid string, name string, uid string, path string) *GroupsPostResponse {
	this := GroupsPostResponse{}
	this.ParentUid = parentUid
	this.Name = name
	this.Uid = uid
	this.Path = path
	return &this
}

// NewGroupsPostResponseWithDefaults instantiates a new GroupsPostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsPostResponseWithDefaults() *GroupsPostResponse {
	this := GroupsPostResponse{}
	return &this
}

// GetParentUid returns the ParentUid field value
func (o *GroupsPostResponse) GetParentUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentUid
}

// GetParentUidOk returns a tuple with the ParentUid field value
// and a boolean to check if the value has been set.
func (o *GroupsPostResponse) GetParentUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentUid, true
}

// SetParentUid sets field value
func (o *GroupsPostResponse) SetParentUid(v string) {
	o.ParentUid = v
}

// GetName returns the Name field value
func (o *GroupsPostResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupsPostResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupsPostResponse) SetName(v string) {
	o.Name = v
}

// GetUid returns the Uid field value
func (o *GroupsPostResponse) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *GroupsPostResponse) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *GroupsPostResponse) SetUid(v string) {
	o.Uid = v
}

// GetPath returns the Path field value
func (o *GroupsPostResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *GroupsPostResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *GroupsPostResponse) SetPath(v string) {
	o.Path = v
}

func (o GroupsPostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsPostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parent_uid"] = o.ParentUid
	toSerialize["name"] = o.Name
	toSerialize["uid"] = o.Uid
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

func (o *GroupsPostResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parent_uid",
		"name",
		"uid",
		"path",
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

	varGroupsPostResponse := _GroupsPostResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupsPostResponse)

	if err != nil {
		return err
	}

	*o = GroupsPostResponse(varGroupsPostResponse)

	return err
}

type NullableGroupsPostResponse struct {
	value *GroupsPostResponse
	isSet bool
}

func (v NullableGroupsPostResponse) Get() *GroupsPostResponse {
	return v.value
}

func (v *NullableGroupsPostResponse) Set(val *GroupsPostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsPostResponse(val *GroupsPostResponse) *NullableGroupsPostResponse {
	return &NullableGroupsPostResponse{value: val, isSet: true}
}

func (v NullableGroupsPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
