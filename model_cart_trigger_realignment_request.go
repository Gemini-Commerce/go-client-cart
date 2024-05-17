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
)

// checks if the CartTriggerRealignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartTriggerRealignmentRequest{}

// CartTriggerRealignmentRequest struct for CartTriggerRealignmentRequest
type CartTriggerRealignmentRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	CartId *string `json:"cartId,omitempty"`
}

// NewCartTriggerRealignmentRequest instantiates a new CartTriggerRealignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartTriggerRealignmentRequest() *CartTriggerRealignmentRequest {
	this := CartTriggerRealignmentRequest{}
	return &this
}

// NewCartTriggerRealignmentRequestWithDefaults instantiates a new CartTriggerRealignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartTriggerRealignmentRequestWithDefaults() *CartTriggerRealignmentRequest {
	this := CartTriggerRealignmentRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CartTriggerRealignmentRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartTriggerRealignmentRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CartTriggerRealignmentRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CartTriggerRealignmentRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCartId returns the CartId field value if set, zero value otherwise.
func (o *CartTriggerRealignmentRequest) GetCartId() string {
	if o == nil || IsNil(o.CartId) {
		var ret string
		return ret
	}
	return *o.CartId
}

// GetCartIdOk returns a tuple with the CartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartTriggerRealignmentRequest) GetCartIdOk() (*string, bool) {
	if o == nil || IsNil(o.CartId) {
		return nil, false
	}
	return o.CartId, true
}

// HasCartId returns a boolean if a field has been set.
func (o *CartTriggerRealignmentRequest) HasCartId() bool {
	if o != nil && !IsNil(o.CartId) {
		return true
	}

	return false
}

// SetCartId gets a reference to the given string and assigns it to the CartId field.
func (o *CartTriggerRealignmentRequest) SetCartId(v string) {
	o.CartId = &v
}

func (o CartTriggerRealignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartTriggerRealignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.CartId) {
		toSerialize["cartId"] = o.CartId
	}
	return toSerialize, nil
}

type NullableCartTriggerRealignmentRequest struct {
	value *CartTriggerRealignmentRequest
	isSet bool
}

func (v NullableCartTriggerRealignmentRequest) Get() *CartTriggerRealignmentRequest {
	return v.value
}

func (v *NullableCartTriggerRealignmentRequest) Set(val *CartTriggerRealignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCartTriggerRealignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCartTriggerRealignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartTriggerRealignmentRequest(val *CartTriggerRealignmentRequest) *NullableCartTriggerRealignmentRequest {
	return &NullableCartTriggerRealignmentRequest{value: val, isSet: true}
}

func (v NullableCartTriggerRealignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartTriggerRealignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

