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

// checks if the CartSetAdditionalInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartSetAdditionalInfoRequest{}

// CartSetAdditionalInfoRequest struct for CartSetAdditionalInfoRequest
type CartSetAdditionalInfoRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	CartId *string `json:"cartId,omitempty"`
	ProductId *string `json:"productId,omitempty"`
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
}

// NewCartSetAdditionalInfoRequest instantiates a new CartSetAdditionalInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartSetAdditionalInfoRequest() *CartSetAdditionalInfoRequest {
	this := CartSetAdditionalInfoRequest{}
	return &this
}

// NewCartSetAdditionalInfoRequestWithDefaults instantiates a new CartSetAdditionalInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartSetAdditionalInfoRequestWithDefaults() *CartSetAdditionalInfoRequest {
	this := CartSetAdditionalInfoRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CartSetAdditionalInfoRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetAdditionalInfoRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// IsSetTenantId returns a boolean if a field has been set.
func (o *CartSetAdditionalInfoRequest) IsSetTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CartSetAdditionalInfoRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCartId returns the CartId field value if set, zero value otherwise.
func (o *CartSetAdditionalInfoRequest) GetCartId() string {
	if o == nil || IsNil(o.CartId) {
		var ret string
		return ret
	}
	return *o.CartId
}

// GetCartIdOk returns a tuple with the CartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetAdditionalInfoRequest) GetCartIdOk() (*string, bool) {
	if o == nil || IsNil(o.CartId) {
		return nil, false
	}
	return o.CartId, true
}

// IsSetCartId returns a boolean if a field has been set.
func (o *CartSetAdditionalInfoRequest) IsSetCartId() bool {
	if o != nil && !IsNil(o.CartId) {
		return true
	}

	return false
}

// SetCartId gets a reference to the given string and assigns it to the CartId field.
func (o *CartSetAdditionalInfoRequest) SetCartId(v string) {
	o.CartId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *CartSetAdditionalInfoRequest) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetAdditionalInfoRequest) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// IsSetProductId returns a boolean if a field has been set.
func (o *CartSetAdditionalInfoRequest) IsSetProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *CartSetAdditionalInfoRequest) SetProductId(v string) {
	o.ProductId = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *CartSetAdditionalInfoRequest) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetAdditionalInfoRequest) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// IsSetAdditionalInfo returns a boolean if a field has been set.
func (o *CartSetAdditionalInfoRequest) IsSetAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *CartSetAdditionalInfoRequest) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

func (o CartSetAdditionalInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartSetAdditionalInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.CartId) {
		toSerialize["cartId"] = o.CartId
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	return toSerialize, nil
}

type NullableCartSetAdditionalInfoRequest struct {
	value *CartSetAdditionalInfoRequest
	isSet bool
}

func (v NullableCartSetAdditionalInfoRequest) Get() *CartSetAdditionalInfoRequest {
	return v.value
}

func (v *NullableCartSetAdditionalInfoRequest) Set(val *CartSetAdditionalInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCartSetAdditionalInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCartSetAdditionalInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartSetAdditionalInfoRequest(val *CartSetAdditionalInfoRequest) *NullableCartSetAdditionalInfoRequest {
	return &NullableCartSetAdditionalInfoRequest{value: val, isSet: true}
}

func (v NullableCartSetAdditionalInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartSetAdditionalInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


