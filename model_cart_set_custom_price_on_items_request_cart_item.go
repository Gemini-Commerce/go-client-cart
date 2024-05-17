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

// checks if the CartSetCustomPriceOnItemsRequestCartItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartSetCustomPriceOnItemsRequestCartItem{}

// CartSetCustomPriceOnItemsRequestCartItem struct for CartSetCustomPriceOnItemsRequestCartItem
type CartSetCustomPriceOnItemsRequestCartItem struct {
	Id *string `json:"id,omitempty"`
	FreeOfCharge *bool `json:"freeOfCharge,omitempty"`
	Unset *bool `json:"unset,omitempty"`
	CustomPriceRow *CartMoney `json:"customPriceRow,omitempty"`
	CustomPriceUnit *CartMoney `json:"customPriceUnit,omitempty"`
	DiscountPercentage *float32 `json:"discountPercentage,omitempty"`
}

// NewCartSetCustomPriceOnItemsRequestCartItem instantiates a new CartSetCustomPriceOnItemsRequestCartItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartSetCustomPriceOnItemsRequestCartItem() *CartSetCustomPriceOnItemsRequestCartItem {
	this := CartSetCustomPriceOnItemsRequestCartItem{}
	return &this
}

// NewCartSetCustomPriceOnItemsRequestCartItemWithDefaults instantiates a new CartSetCustomPriceOnItemsRequestCartItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartSetCustomPriceOnItemsRequestCartItemWithDefaults() *CartSetCustomPriceOnItemsRequestCartItem {
	this := CartSetCustomPriceOnItemsRequestCartItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CartSetCustomPriceOnItemsRequestCartItem) SetId(v string) {
	o.Id = &v
}

// GetFreeOfCharge returns the FreeOfCharge field value if set, zero value otherwise.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetFreeOfCharge() bool {
	if o == nil || IsNil(o.FreeOfCharge) {
		var ret bool
		return ret
	}
	return *o.FreeOfCharge
}

// GetFreeOfChargeOk returns a tuple with the FreeOfCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetFreeOfChargeOk() (*bool, bool) {
	if o == nil || IsNil(o.FreeOfCharge) {
		return nil, false
	}
	return o.FreeOfCharge, true
}

// HasFreeOfCharge returns a boolean if a field has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) HasFreeOfCharge() bool {
	if o != nil && !IsNil(o.FreeOfCharge) {
		return true
	}

	return false
}

// SetFreeOfCharge gets a reference to the given bool and assigns it to the FreeOfCharge field.
func (o *CartSetCustomPriceOnItemsRequestCartItem) SetFreeOfCharge(v bool) {
	o.FreeOfCharge = &v
}

// GetUnset returns the Unset field value if set, zero value otherwise.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetUnset() bool {
	if o == nil || IsNil(o.Unset) {
		var ret bool
		return ret
	}
	return *o.Unset
}

// GetUnsetOk returns a tuple with the Unset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetUnsetOk() (*bool, bool) {
	if o == nil || IsNil(o.Unset) {
		return nil, false
	}
	return o.Unset, true
}

// HasUnset returns a boolean if a field has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) HasUnset() bool {
	if o != nil && !IsNil(o.Unset) {
		return true
	}

	return false
}

// SetUnset gets a reference to the given bool and assigns it to the Unset field.
func (o *CartSetCustomPriceOnItemsRequestCartItem) SetUnset(v bool) {
	o.Unset = &v
}

// GetCustomPriceRow returns the CustomPriceRow field value if set, zero value otherwise.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetCustomPriceRow() CartMoney {
	if o == nil || IsNil(o.CustomPriceRow) {
		var ret CartMoney
		return ret
	}
	return *o.CustomPriceRow
}

// GetCustomPriceRowOk returns a tuple with the CustomPriceRow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetCustomPriceRowOk() (*CartMoney, bool) {
	if o == nil || IsNil(o.CustomPriceRow) {
		return nil, false
	}
	return o.CustomPriceRow, true
}

// HasCustomPriceRow returns a boolean if a field has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) HasCustomPriceRow() bool {
	if o != nil && !IsNil(o.CustomPriceRow) {
		return true
	}

	return false
}

// SetCustomPriceRow gets a reference to the given CartMoney and assigns it to the CustomPriceRow field.
func (o *CartSetCustomPriceOnItemsRequestCartItem) SetCustomPriceRow(v CartMoney) {
	o.CustomPriceRow = &v
}

// GetCustomPriceUnit returns the CustomPriceUnit field value if set, zero value otherwise.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetCustomPriceUnit() CartMoney {
	if o == nil || IsNil(o.CustomPriceUnit) {
		var ret CartMoney
		return ret
	}
	return *o.CustomPriceUnit
}

// GetCustomPriceUnitOk returns a tuple with the CustomPriceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetCustomPriceUnitOk() (*CartMoney, bool) {
	if o == nil || IsNil(o.CustomPriceUnit) {
		return nil, false
	}
	return o.CustomPriceUnit, true
}

// HasCustomPriceUnit returns a boolean if a field has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) HasCustomPriceUnit() bool {
	if o != nil && !IsNil(o.CustomPriceUnit) {
		return true
	}

	return false
}

// SetCustomPriceUnit gets a reference to the given CartMoney and assigns it to the CustomPriceUnit field.
func (o *CartSetCustomPriceOnItemsRequestCartItem) SetCustomPriceUnit(v CartMoney) {
	o.CustomPriceUnit = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetDiscountPercentage() float32 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret float32
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) GetDiscountPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *CartSetCustomPriceOnItemsRequestCartItem) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given float32 and assigns it to the DiscountPercentage field.
func (o *CartSetCustomPriceOnItemsRequestCartItem) SetDiscountPercentage(v float32) {
	o.DiscountPercentage = &v
}

func (o CartSetCustomPriceOnItemsRequestCartItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartSetCustomPriceOnItemsRequestCartItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FreeOfCharge) {
		toSerialize["freeOfCharge"] = o.FreeOfCharge
	}
	if !IsNil(o.Unset) {
		toSerialize["unset"] = o.Unset
	}
	if !IsNil(o.CustomPriceRow) {
		toSerialize["customPriceRow"] = o.CustomPriceRow
	}
	if !IsNil(o.CustomPriceUnit) {
		toSerialize["customPriceUnit"] = o.CustomPriceUnit
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}
	return toSerialize, nil
}

type NullableCartSetCustomPriceOnItemsRequestCartItem struct {
	value *CartSetCustomPriceOnItemsRequestCartItem
	isSet bool
}

func (v NullableCartSetCustomPriceOnItemsRequestCartItem) Get() *CartSetCustomPriceOnItemsRequestCartItem {
	return v.value
}

func (v *NullableCartSetCustomPriceOnItemsRequestCartItem) Set(val *CartSetCustomPriceOnItemsRequestCartItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCartSetCustomPriceOnItemsRequestCartItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCartSetCustomPriceOnItemsRequestCartItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartSetCustomPriceOnItemsRequestCartItem(val *CartSetCustomPriceOnItemsRequestCartItem) *NullableCartSetCustomPriceOnItemsRequestCartItem {
	return &NullableCartSetCustomPriceOnItemsRequestCartItem{value: val, isSet: true}
}

func (v NullableCartSetCustomPriceOnItemsRequestCartItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartSetCustomPriceOnItemsRequestCartItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


