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

// checks if the CartUpdateCartRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartUpdateCartRequest{}

// CartUpdateCartRequest struct for CartUpdateCartRequest
type CartUpdateCartRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	CartId *string `json:"cartId,omitempty"`
	Payload *UpdateCartRequestPayload `json:"payload,omitempty"`
	PayloadMask *string `json:"payloadMask,omitempty"`
}

// NewCartUpdateCartRequest instantiates a new CartUpdateCartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartUpdateCartRequest() *CartUpdateCartRequest {
	this := CartUpdateCartRequest{}
	return &this
}

// NewCartUpdateCartRequestWithDefaults instantiates a new CartUpdateCartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartUpdateCartRequestWithDefaults() *CartUpdateCartRequest {
	this := CartUpdateCartRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CartUpdateCartRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartUpdateCartRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// IsSetTenantId returns a boolean if a field has been set.
func (o *CartUpdateCartRequest) IsSetTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CartUpdateCartRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCartId returns the CartId field value if set, zero value otherwise.
func (o *CartUpdateCartRequest) GetCartId() string {
	if o == nil || IsNil(o.CartId) {
		var ret string
		return ret
	}
	return *o.CartId
}

// GetCartIdOk returns a tuple with the CartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartUpdateCartRequest) GetCartIdOk() (*string, bool) {
	if o == nil || IsNil(o.CartId) {
		return nil, false
	}
	return o.CartId, true
}

// IsSetCartId returns a boolean if a field has been set.
func (o *CartUpdateCartRequest) IsSetCartId() bool {
	if o != nil && !IsNil(o.CartId) {
		return true
	}

	return false
}

// SetCartId gets a reference to the given string and assigns it to the CartId field.
func (o *CartUpdateCartRequest) SetCartId(v string) {
	o.CartId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CartUpdateCartRequest) GetPayload() UpdateCartRequestPayload {
	if o == nil || IsNil(o.Payload) {
		var ret UpdateCartRequestPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartUpdateCartRequest) GetPayloadOk() (*UpdateCartRequestPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// IsSetPayload returns a boolean if a field has been set.
func (o *CartUpdateCartRequest) IsSetPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given UpdateCartRequestPayload and assigns it to the Payload field.
func (o *CartUpdateCartRequest) SetPayload(v UpdateCartRequestPayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *CartUpdateCartRequest) GetPayloadMask() string {
	if o == nil || IsNil(o.PayloadMask) {
		var ret string
		return ret
	}
	return *o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartUpdateCartRequest) GetPayloadMaskOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadMask) {
		return nil, false
	}
	return o.PayloadMask, true
}

// IsSetPayloadMask returns a boolean if a field has been set.
func (o *CartUpdateCartRequest) IsSetPayloadMask() bool {
	if o != nil && !IsNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given string and assigns it to the PayloadMask field.
func (o *CartUpdateCartRequest) SetPayloadMask(v string) {
	o.PayloadMask = &v
}

func (o CartUpdateCartRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartUpdateCartRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.CartId) {
		toSerialize["cartId"] = o.CartId
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.PayloadMask) {
		toSerialize["payloadMask"] = o.PayloadMask
	}
	return toSerialize, nil
}

type NullableCartUpdateCartRequest struct {
	value *CartUpdateCartRequest
	isSet bool
}

func (v NullableCartUpdateCartRequest) Get() *CartUpdateCartRequest {
	return v.value
}

func (v *NullableCartUpdateCartRequest) Set(val *CartUpdateCartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCartUpdateCartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCartUpdateCartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartUpdateCartRequest(val *CartUpdateCartRequest) *NullableCartUpdateCartRequest {
	return &NullableCartUpdateCartRequest{value: val, isSet: true}
}

func (v NullableCartUpdateCartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartUpdateCartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


