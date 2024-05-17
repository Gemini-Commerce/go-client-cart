/*
Cart Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cart

import (
	"encoding/json"
	"time"
)

// checks if the ListCartsRequestFilterDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCartsRequestFilterDate{}

// ListCartsRequestFilterDate struct for ListCartsRequestFilterDate
type ListCartsRequestFilterDate struct {
	From *time.Time `json:"from,omitempty"`
	To *time.Time `json:"to,omitempty"`
}

// NewListCartsRequestFilterDate instantiates a new ListCartsRequestFilterDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCartsRequestFilterDate() *ListCartsRequestFilterDate {
	this := ListCartsRequestFilterDate{}
	return &this
}

// NewListCartsRequestFilterDateWithDefaults instantiates a new ListCartsRequestFilterDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCartsRequestFilterDateWithDefaults() *ListCartsRequestFilterDate {
	this := ListCartsRequestFilterDate{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ListCartsRequestFilterDate) GetFrom() time.Time {
	if o == nil || IsNil(o.From) {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCartsRequestFilterDate) GetFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ListCartsRequestFilterDate) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *ListCartsRequestFilterDate) SetFrom(v time.Time) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ListCartsRequestFilterDate) GetTo() time.Time {
	if o == nil || IsNil(o.To) {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCartsRequestFilterDate) GetToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ListCartsRequestFilterDate) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *ListCartsRequestFilterDate) SetTo(v time.Time) {
	o.To = &v
}

func (o ListCartsRequestFilterDate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCartsRequestFilterDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableListCartsRequestFilterDate struct {
	value *ListCartsRequestFilterDate
	isSet bool
}

func (v NullableListCartsRequestFilterDate) Get() *ListCartsRequestFilterDate {
	return v.value
}

func (v *NullableListCartsRequestFilterDate) Set(val *ListCartsRequestFilterDate) {
	v.value = val
	v.isSet = true
}

func (v NullableListCartsRequestFilterDate) IsSet() bool {
	return v.isSet
}

func (v *NullableListCartsRequestFilterDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCartsRequestFilterDate(val *ListCartsRequestFilterDate) *NullableListCartsRequestFilterDate {
	return &NullableListCartsRequestFilterDate{value: val, isSet: true}
}

func (v NullableListCartsRequestFilterDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCartsRequestFilterDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


