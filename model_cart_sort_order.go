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
	"fmt"
)

// CartSortOrder the model 'CartSortOrder'
type CartSortOrder string

// List of cartSortOrder
const (
	CARTSORTORDER_DESC CartSortOrder = "DESC"
	CARTSORTORDER_ASC CartSortOrder = "ASC"
)

// All allowed values of CartSortOrder enum
var AllowedCartSortOrderEnumValues = []CartSortOrder{
	"DESC",
	"ASC",
}

func (v *CartSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CartSortOrder(value)
	for _, existing := range AllowedCartSortOrderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CartSortOrder", value)
}

// NewCartSortOrderFromValue returns a pointer to a valid CartSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCartSortOrderFromValue(v string) (*CartSortOrder, error) {
	ev := CartSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CartSortOrder: valid values are %v", v, AllowedCartSortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CartSortOrder) IsValid() bool {
	for _, existing := range AllowedCartSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cartSortOrder value
func (v CartSortOrder) Ptr() *CartSortOrder {
	return &v
}

type NullableCartSortOrder struct {
	value *CartSortOrder
	isSet bool
}

func (v NullableCartSortOrder) Get() *CartSortOrder {
	return v.value
}

func (v *NullableCartSortOrder) Set(val *CartSortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCartSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCartSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartSortOrder(val *CartSortOrder) *NullableCartSortOrder {
	return &NullableCartSortOrder{value: val, isSet: true}
}

func (v NullableCartSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

