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
	"time"
)

// checks if the WiredNetworksItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WiredNetworksItem{}

// WiredNetworksItem struct for WiredNetworksItem
type WiredNetworksItem struct {
	Uid                  string         `json:"uid"`
	Alias                string         `json:"alias"`
	IpVersion            string         `json:"ip_version"`
	DatetimeUpdated      time.Time      `json:"datetime_updated"`
	DatetimeCreated      time.Time      `json:"datetime_created"`
	Security             NullableString `json:"security"`
	DnsLookupDomain      NullableString `json:"dns_lookup_domain"`
	DisableEdns          bool           `json:"disable_edns"`
	UseDns64             bool           `json:"use_dns64"`
	ExternalConnectivity bool           `json:"external_connectivity"`
	VlanId               NullableInt32  `json:"vlan_id"`
}

type _WiredNetworksItem WiredNetworksItem

// NewWiredNetworksItem instantiates a new WiredNetworksItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWiredNetworksItem(uid string, alias string, ipVersion string, datetimeUpdated time.Time, datetimeCreated time.Time, security NullableString, dnsLookupDomain NullableString, disableEdns bool, useDns64 bool, externalConnectivity bool, vlanId NullableInt32) *WiredNetworksItem {
	this := WiredNetworksItem{}
	this.Uid = uid
	this.Alias = alias
	this.IpVersion = ipVersion
	this.DatetimeUpdated = datetimeUpdated
	this.DatetimeCreated = datetimeCreated
	this.Security = security
	this.DnsLookupDomain = dnsLookupDomain
	this.DisableEdns = disableEdns
	this.UseDns64 = useDns64
	this.ExternalConnectivity = externalConnectivity
	this.VlanId = vlanId
	return &this
}

// NewWiredNetworksItemWithDefaults instantiates a new WiredNetworksItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWiredNetworksItemWithDefaults() *WiredNetworksItem {
	this := WiredNetworksItem{}
	return &this
}

// GetUid returns the Uid field value
func (o *WiredNetworksItem) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *WiredNetworksItem) SetUid(v string) {
	o.Uid = v
}

// GetAlias returns the Alias field value
func (o *WiredNetworksItem) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *WiredNetworksItem) SetAlias(v string) {
	o.Alias = v
}

// GetIpVersion returns the IpVersion field value
func (o *WiredNetworksItem) GetIpVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetIpVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpVersion, true
}

// SetIpVersion sets field value
func (o *WiredNetworksItem) SetIpVersion(v string) {
	o.IpVersion = v
}

// GetDatetimeUpdated returns the DatetimeUpdated field value
func (o *WiredNetworksItem) GetDatetimeUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DatetimeUpdated
}

// GetDatetimeUpdatedOk returns a tuple with the DatetimeUpdated field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetDatetimeUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatetimeUpdated, true
}

// SetDatetimeUpdated sets field value
func (o *WiredNetworksItem) SetDatetimeUpdated(v time.Time) {
	o.DatetimeUpdated = v
}

// GetDatetimeCreated returns the DatetimeCreated field value
func (o *WiredNetworksItem) GetDatetimeCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DatetimeCreated
}

// GetDatetimeCreatedOk returns a tuple with the DatetimeCreated field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetDatetimeCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatetimeCreated, true
}

// SetDatetimeCreated sets field value
func (o *WiredNetworksItem) SetDatetimeCreated(v time.Time) {
	o.DatetimeCreated = v
}

// GetSecurity returns the Security field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WiredNetworksItem) GetSecurity() string {
	if o == nil || o.Security.Get() == nil {
		var ret string
		return ret
	}

	return *o.Security.Get()
}

// GetSecurityOk returns a tuple with the Security field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WiredNetworksItem) GetSecurityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Security.Get(), o.Security.IsSet()
}

// SetSecurity sets field value
func (o *WiredNetworksItem) SetSecurity(v string) {
	o.Security.Set(&v)
}

// GetDnsLookupDomain returns the DnsLookupDomain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WiredNetworksItem) GetDnsLookupDomain() string {
	if o == nil || o.DnsLookupDomain.Get() == nil {
		var ret string
		return ret
	}

	return *o.DnsLookupDomain.Get()
}

// GetDnsLookupDomainOk returns a tuple with the DnsLookupDomain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WiredNetworksItem) GetDnsLookupDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsLookupDomain.Get(), o.DnsLookupDomain.IsSet()
}

// SetDnsLookupDomain sets field value
func (o *WiredNetworksItem) SetDnsLookupDomain(v string) {
	o.DnsLookupDomain.Set(&v)
}

// GetDisableEdns returns the DisableEdns field value
func (o *WiredNetworksItem) GetDisableEdns() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEdns
}

// GetDisableEdnsOk returns a tuple with the DisableEdns field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetDisableEdnsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEdns, true
}

// SetDisableEdns sets field value
func (o *WiredNetworksItem) SetDisableEdns(v bool) {
	o.DisableEdns = v
}

// GetUseDns64 returns the UseDns64 field value
func (o *WiredNetworksItem) GetUseDns64() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseDns64
}

// GetUseDns64Ok returns a tuple with the UseDns64 field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetUseDns64Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseDns64, true
}

// SetUseDns64 sets field value
func (o *WiredNetworksItem) SetUseDns64(v bool) {
	o.UseDns64 = v
}

// GetExternalConnectivity returns the ExternalConnectivity field value
func (o *WiredNetworksItem) GetExternalConnectivity() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExternalConnectivity
}

// GetExternalConnectivityOk returns a tuple with the ExternalConnectivity field value
// and a boolean to check if the value has been set.
func (o *WiredNetworksItem) GetExternalConnectivityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalConnectivity, true
}

// SetExternalConnectivity sets field value
func (o *WiredNetworksItem) SetExternalConnectivity(v bool) {
	o.ExternalConnectivity = v
}

// GetVlanId returns the VlanId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *WiredNetworksItem) GetVlanId() int32 {
	if o == nil || o.VlanId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.VlanId.Get()
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WiredNetworksItem) GetVlanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanId.Get(), o.VlanId.IsSet()
}

// SetVlanId sets field value
func (o *WiredNetworksItem) SetVlanId(v int32) {
	o.VlanId.Set(&v)
}

func (o WiredNetworksItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WiredNetworksItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uid"] = o.Uid
	toSerialize["alias"] = o.Alias
	toSerialize["ip_version"] = o.IpVersion
	toSerialize["datetime_updated"] = o.DatetimeUpdated
	toSerialize["datetime_created"] = o.DatetimeCreated
	toSerialize["security"] = o.Security.Get()
	toSerialize["dns_lookup_domain"] = o.DnsLookupDomain.Get()
	toSerialize["disable_edns"] = o.DisableEdns
	toSerialize["use_dns64"] = o.UseDns64
	toSerialize["external_connectivity"] = o.ExternalConnectivity
	toSerialize["vlan_id"] = o.VlanId.Get()
	return toSerialize, nil
}

func (o *WiredNetworksItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uid",
		"alias",
		"ip_version",
		"datetime_updated",
		"datetime_created",
		"security",
		"dns_lookup_domain",
		"disable_edns",
		"use_dns64",
		"external_connectivity",
		"vlan_id",
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

	varWiredNetworksItem := _WiredNetworksItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWiredNetworksItem)

	if err != nil {
		return err
	}

	*o = WiredNetworksItem(varWiredNetworksItem)

	return err
}

type NullableWiredNetworksItem struct {
	value *WiredNetworksItem
	isSet bool
}

func (v NullableWiredNetworksItem) Get() *WiredNetworksItem {
	return v.value
}

func (v *NullableWiredNetworksItem) Set(val *WiredNetworksItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWiredNetworksItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWiredNetworksItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWiredNetworksItem(val *WiredNetworksItem) *NullableWiredNetworksItem {
	return &NullableWiredNetworksItem{value: val, isSet: true}
}

func (v NullableWiredNetworksItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWiredNetworksItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
