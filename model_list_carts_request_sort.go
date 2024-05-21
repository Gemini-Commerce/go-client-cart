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

// checks if the ListCartsRequestSort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCartsRequestSort{}

// ListCartsRequestSort struct for ListCartsRequestSort
type ListCartsRequestSort struct {
	// evaluation_order is the order in which the sort will be applied. The lower the number, the earlier the sort will be applied.
	EvaluationOrder *int64 `json:"evaluationOrder,omitempty"`
	Field *SortSortField `json:"field,omitempty"`
	Order *CartSortOrder `json:"order,omitempty"`
}

// NewListCartsRequestSort instantiates a new ListCartsRequestSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCartsRequestSort() *ListCartsRequestSort {
	this := ListCartsRequestSort{}
	var field SortSortField = SORTSORTFIELD_UNKNOWN
	this.Field = &field
	var order CartSortOrder = CARTSORTORDER_DESC
	this.Order = &order
	return &this
}

// NewListCartsRequestSortWithDefaults instantiates a new ListCartsRequestSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCartsRequestSortWithDefaults() *ListCartsRequestSort {
	this := ListCartsRequestSort{}
	var field SortSortField = SORTSORTFIELD_UNKNOWN
	this.Field = &field
	var order CartSortOrder = CARTSORTORDER_DESC
	this.Order = &order
	return &this
}

// GetEvaluationOrder returns the EvaluationOrder field value if set, zero value otherwise.
func (o *ListCartsRequestSort) GetEvaluationOrder() int64 {
	if o == nil || IsNil(o.EvaluationOrder) {
		var ret int64
		return ret
	}
	return *o.EvaluationOrder
}

// GetEvaluationOrderOk returns a tuple with the EvaluationOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCartsRequestSort) GetEvaluationOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.EvaluationOrder) {
		return nil, false
	}
	return o.EvaluationOrder, true
}

// IsSetEvaluationOrder returns a boolean if a field has been set.
func (o *ListCartsRequestSort) IsSetEvaluationOrder() bool {
	if o != nil && !IsNil(o.EvaluationOrder) {
		return true
	}

	return false
}

// SetEvaluationOrder gets a reference to the given int64 and assigns it to the EvaluationOrder field.
func (o *ListCartsRequestSort) SetEvaluationOrder(v int64) {
	o.EvaluationOrder = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ListCartsRequestSort) GetField() SortSortField {
	if o == nil || IsNil(o.Field) {
		var ret SortSortField
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCartsRequestSort) GetFieldOk() (*SortSortField, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// IsSetField returns a boolean if a field has been set.
func (o *ListCartsRequestSort) IsSetField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given SortSortField and assigns it to the Field field.
func (o *ListCartsRequestSort) SetField(v SortSortField) {
	o.Field = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ListCartsRequestSort) GetOrder() CartSortOrder {
	if o == nil || IsNil(o.Order) {
		var ret CartSortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCartsRequestSort) GetOrderOk() (*CartSortOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// IsSetOrder returns a boolean if a field has been set.
func (o *ListCartsRequestSort) IsSetOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given CartSortOrder and assigns it to the Order field.
func (o *ListCartsRequestSort) SetOrder(v CartSortOrder) {
	o.Order = &v
}

func (o ListCartsRequestSort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCartsRequestSort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EvaluationOrder) {
		toSerialize["evaluationOrder"] = o.EvaluationOrder
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableListCartsRequestSort struct {
	value *ListCartsRequestSort
	isSet bool
}

func (v NullableListCartsRequestSort) Get() *ListCartsRequestSort {
	return v.value
}

func (v *NullableListCartsRequestSort) Set(val *ListCartsRequestSort) {
	v.value = val
	v.isSet = true
}

func (v NullableListCartsRequestSort) IsSet() bool {
	return v.isSet
}

func (v *NullableListCartsRequestSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCartsRequestSort(val *ListCartsRequestSort) *NullableListCartsRequestSort {
	return &NullableListCartsRequestSort{value: val, isSet: true}
}

func (v NullableListCartsRequestSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCartsRequestSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


