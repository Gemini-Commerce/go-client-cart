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

// checks if the CartListShipmentMethodsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartListShipmentMethodsResponse{}

// CartListShipmentMethodsResponse struct for CartListShipmentMethodsResponse
type CartListShipmentMethodsResponse struct {
	Shipments []CartShipmentData `json:"shipments,omitempty"`
}

// NewCartListShipmentMethodsResponse instantiates a new CartListShipmentMethodsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartListShipmentMethodsResponse() *CartListShipmentMethodsResponse {
	this := CartListShipmentMethodsResponse{}
	return &this
}

// NewCartListShipmentMethodsResponseWithDefaults instantiates a new CartListShipmentMethodsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartListShipmentMethodsResponseWithDefaults() *CartListShipmentMethodsResponse {
	this := CartListShipmentMethodsResponse{}
	return &this
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *CartListShipmentMethodsResponse) GetShipments() []CartShipmentData {
	if o == nil || IsNil(o.Shipments) {
		var ret []CartShipmentData
		return ret
	}
	return o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartListShipmentMethodsResponse) GetShipmentsOk() ([]CartShipmentData, bool) {
	if o == nil || IsNil(o.Shipments) {
		return nil, false
	}
	return o.Shipments, true
}

// HasShipments returns a boolean if a field has been set.
func (o *CartListShipmentMethodsResponse) HasShipments() bool {
	if o != nil && !IsNil(o.Shipments) {
		return true
	}

	return false
}

// SetShipments gets a reference to the given []CartShipmentData and assigns it to the Shipments field.
func (o *CartListShipmentMethodsResponse) SetShipments(v []CartShipmentData) {
	o.Shipments = v
}

func (o CartListShipmentMethodsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartListShipmentMethodsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shipments) {
		toSerialize["shipments"] = o.Shipments
	}
	return toSerialize, nil
}

type NullableCartListShipmentMethodsResponse struct {
	value *CartListShipmentMethodsResponse
	isSet bool
}

func (v NullableCartListShipmentMethodsResponse) Get() *CartListShipmentMethodsResponse {
	return v.value
}

func (v *NullableCartListShipmentMethodsResponse) Set(val *CartListShipmentMethodsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCartListShipmentMethodsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCartListShipmentMethodsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartListShipmentMethodsResponse(val *CartListShipmentMethodsResponse) *NullableCartListShipmentMethodsResponse {
	return &NullableCartListShipmentMethodsResponse{value: val, isSet: true}
}

func (v NullableCartListShipmentMethodsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartListShipmentMethodsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


