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

// checks if the CartCartData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartCartData{}

// CartCartData struct for CartCartData
type CartCartData struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Channel *string `json:"channel,omitempty"`
	Market *string `json:"market,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Items []CartCartItem `json:"items,omitempty"`
	Payments []CartPaymentData `json:"payments,omitempty"`
	Shipments []CartShipmentData `json:"shipments,omitempty"`
	Promotions []CartPromotionData `json:"promotions,omitempty"`
	Currency *CartCurrency `json:"currency,omitempty"`
	Subtotals []CartCartSubtotal `json:"subtotals,omitempty"`
	Total *CartMoney `json:"total,omitempty"`
	TotalDue *CartMoney `json:"totalDue,omitempty"`
	VatIncluded *bool `json:"vatIncluded,omitempty"`
	BillingAddress *CartPostalAddress `json:"billingAddress,omitempty"`
	ShippingAddress *CartPostalAddress `json:"shippingAddress,omitempty"`
	Status *CartCartStatus `json:"status,omitempty"`
	Customer *CartCustomerData `json:"customer,omitempty"`
	Notes *string `json:"notes,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewCartCartData instantiates a new CartCartData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartCartData() *CartCartData {
	this := CartCartData{}
	var currency CartCurrency = CARTCURRENCY_XXX
	this.Currency = &currency
	var status CartCartStatus = CARTCARTSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewCartCartDataWithDefaults instantiates a new CartCartData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartCartDataWithDefaults() *CartCartData {
	this := CartCartData{}
	var currency CartCurrency = CARTCURRENCY_XXX
	this.Currency = &currency
	var status CartCartStatus = CARTCARTSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CartCartData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *CartCartData) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CartCartData) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *CartCartData) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// IsSetGrn returns a boolean if a field has been set.
func (o *CartCartData) IsSetGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *CartCartData) SetGrn(v string) {
	o.Grn = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *CartCartData) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// IsSetChannel returns a boolean if a field has been set.
func (o *CartCartData) IsSetChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *CartCartData) SetChannel(v string) {
	o.Channel = &v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *CartCartData) GetMarket() string {
	if o == nil || IsNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetMarketOk() (*string, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// IsSetMarket returns a boolean if a field has been set.
func (o *CartCartData) IsSetMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *CartCartData) SetMarket(v string) {
	o.Market = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *CartCartData) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// IsSetLocale returns a boolean if a field has been set.
func (o *CartCartData) IsSetLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *CartCartData) SetLocale(v string) {
	o.Locale = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *CartCartData) GetItems() []CartCartItem {
	if o == nil || IsNil(o.Items) {
		var ret []CartCartItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetItemsOk() ([]CartCartItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// IsSetItems returns a boolean if a field has been set.
func (o *CartCartData) IsSetItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CartCartItem and assigns it to the Items field.
func (o *CartCartData) SetItems(v []CartCartItem) {
	o.Items = v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *CartCartData) GetPayments() []CartPaymentData {
	if o == nil || IsNil(o.Payments) {
		var ret []CartPaymentData
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetPaymentsOk() ([]CartPaymentData, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// IsSetPayments returns a boolean if a field has been set.
func (o *CartCartData) IsSetPayments() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []CartPaymentData and assigns it to the Payments field.
func (o *CartCartData) SetPayments(v []CartPaymentData) {
	o.Payments = v
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *CartCartData) GetShipments() []CartShipmentData {
	if o == nil || IsNil(o.Shipments) {
		var ret []CartShipmentData
		return ret
	}
	return o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetShipmentsOk() ([]CartShipmentData, bool) {
	if o == nil || IsNil(o.Shipments) {
		return nil, false
	}
	return o.Shipments, true
}

// IsSetShipments returns a boolean if a field has been set.
func (o *CartCartData) IsSetShipments() bool {
	if o != nil && !IsNil(o.Shipments) {
		return true
	}

	return false
}

// SetShipments gets a reference to the given []CartShipmentData and assigns it to the Shipments field.
func (o *CartCartData) SetShipments(v []CartShipmentData) {
	o.Shipments = v
}

// GetPromotions returns the Promotions field value if set, zero value otherwise.
func (o *CartCartData) GetPromotions() []CartPromotionData {
	if o == nil || IsNil(o.Promotions) {
		var ret []CartPromotionData
		return ret
	}
	return o.Promotions
}

// GetPromotionsOk returns a tuple with the Promotions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetPromotionsOk() ([]CartPromotionData, bool) {
	if o == nil || IsNil(o.Promotions) {
		return nil, false
	}
	return o.Promotions, true
}

// IsSetPromotions returns a boolean if a field has been set.
func (o *CartCartData) IsSetPromotions() bool {
	if o != nil && !IsNil(o.Promotions) {
		return true
	}

	return false
}

// SetPromotions gets a reference to the given []CartPromotionData and assigns it to the Promotions field.
func (o *CartCartData) SetPromotions(v []CartPromotionData) {
	o.Promotions = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CartCartData) GetCurrency() CartCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret CartCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetCurrencyOk() (*CartCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// IsSetCurrency returns a boolean if a field has been set.
func (o *CartCartData) IsSetCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CartCurrency and assigns it to the Currency field.
func (o *CartCartData) SetCurrency(v CartCurrency) {
	o.Currency = &v
}

// GetSubtotals returns the Subtotals field value if set, zero value otherwise.
func (o *CartCartData) GetSubtotals() []CartCartSubtotal {
	if o == nil || IsNil(o.Subtotals) {
		var ret []CartCartSubtotal
		return ret
	}
	return o.Subtotals
}

// GetSubtotalsOk returns a tuple with the Subtotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetSubtotalsOk() ([]CartCartSubtotal, bool) {
	if o == nil || IsNil(o.Subtotals) {
		return nil, false
	}
	return o.Subtotals, true
}

// IsSetSubtotals returns a boolean if a field has been set.
func (o *CartCartData) IsSetSubtotals() bool {
	if o != nil && !IsNil(o.Subtotals) {
		return true
	}

	return false
}

// SetSubtotals gets a reference to the given []CartCartSubtotal and assigns it to the Subtotals field.
func (o *CartCartData) SetSubtotals(v []CartCartSubtotal) {
	o.Subtotals = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CartCartData) GetTotal() CartMoney {
	if o == nil || IsNil(o.Total) {
		var ret CartMoney
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetTotalOk() (*CartMoney, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// IsSetTotal returns a boolean if a field has been set.
func (o *CartCartData) IsSetTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given CartMoney and assigns it to the Total field.
func (o *CartCartData) SetTotal(v CartMoney) {
	o.Total = &v
}

// GetTotalDue returns the TotalDue field value if set, zero value otherwise.
func (o *CartCartData) GetTotalDue() CartMoney {
	if o == nil || IsNil(o.TotalDue) {
		var ret CartMoney
		return ret
	}
	return *o.TotalDue
}

// GetTotalDueOk returns a tuple with the TotalDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetTotalDueOk() (*CartMoney, bool) {
	if o == nil || IsNil(o.TotalDue) {
		return nil, false
	}
	return o.TotalDue, true
}

// IsSetTotalDue returns a boolean if a field has been set.
func (o *CartCartData) IsSetTotalDue() bool {
	if o != nil && !IsNil(o.TotalDue) {
		return true
	}

	return false
}

// SetTotalDue gets a reference to the given CartMoney and assigns it to the TotalDue field.
func (o *CartCartData) SetTotalDue(v CartMoney) {
	o.TotalDue = &v
}

// GetVatIncluded returns the VatIncluded field value if set, zero value otherwise.
func (o *CartCartData) GetVatIncluded() bool {
	if o == nil || IsNil(o.VatIncluded) {
		var ret bool
		return ret
	}
	return *o.VatIncluded
}

// GetVatIncludedOk returns a tuple with the VatIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetVatIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.VatIncluded) {
		return nil, false
	}
	return o.VatIncluded, true
}

// IsSetVatIncluded returns a boolean if a field has been set.
func (o *CartCartData) IsSetVatIncluded() bool {
	if o != nil && !IsNil(o.VatIncluded) {
		return true
	}

	return false
}

// SetVatIncluded gets a reference to the given bool and assigns it to the VatIncluded field.
func (o *CartCartData) SetVatIncluded(v bool) {
	o.VatIncluded = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *CartCartData) GetBillingAddress() CartPostalAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret CartPostalAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetBillingAddressOk() (*CartPostalAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// IsSetBillingAddress returns a boolean if a field has been set.
func (o *CartCartData) IsSetBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given CartPostalAddress and assigns it to the BillingAddress field.
func (o *CartCartData) SetBillingAddress(v CartPostalAddress) {
	o.BillingAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *CartCartData) GetShippingAddress() CartPostalAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret CartPostalAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetShippingAddressOk() (*CartPostalAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// IsSetShippingAddress returns a boolean if a field has been set.
func (o *CartCartData) IsSetShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given CartPostalAddress and assigns it to the ShippingAddress field.
func (o *CartCartData) SetShippingAddress(v CartPostalAddress) {
	o.ShippingAddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CartCartData) GetStatus() CartCartStatus {
	if o == nil || IsNil(o.Status) {
		var ret CartCartStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetStatusOk() (*CartCartStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// IsSetStatus returns a boolean if a field has been set.
func (o *CartCartData) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CartCartStatus and assigns it to the Status field.
func (o *CartCartData) SetStatus(v CartCartStatus) {
	o.Status = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CartCartData) GetCustomer() CartCustomerData {
	if o == nil || IsNil(o.Customer) {
		var ret CartCustomerData
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetCustomerOk() (*CartCustomerData, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// IsSetCustomer returns a boolean if a field has been set.
func (o *CartCartData) IsSetCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CartCustomerData and assigns it to the Customer field.
func (o *CartCartData) SetCustomer(v CartCustomerData) {
	o.Customer = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CartCartData) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// IsSetNotes returns a boolean if a field has been set.
func (o *CartCartData) IsSetNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CartCartData) SetNotes(v string) {
	o.Notes = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CartCartData) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// IsSetCreatedAt returns a boolean if a field has been set.
func (o *CartCartData) IsSetCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CartCartData) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CartCartData) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCartData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// IsSetUpdatedAt returns a boolean if a field has been set.
func (o *CartCartData) IsSetUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CartCartData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CartCartData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartCartData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Payments) {
		toSerialize["payments"] = o.Payments
	}
	if !IsNil(o.Shipments) {
		toSerialize["shipments"] = o.Shipments
	}
	if !IsNil(o.Promotions) {
		toSerialize["promotions"] = o.Promotions
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Subtotals) {
		toSerialize["subtotals"] = o.Subtotals
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.TotalDue) {
		toSerialize["totalDue"] = o.TotalDue
	}
	if !IsNil(o.VatIncluded) {
		toSerialize["vatIncluded"] = o.VatIncluded
	}
	if !IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shippingAddress"] = o.ShippingAddress
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableCartCartData struct {
	value *CartCartData
	isSet bool
}

func (v NullableCartCartData) Get() *CartCartData {
	return v.value
}

func (v *NullableCartCartData) Set(val *CartCartData) {
	v.value = val
	v.isSet = true
}

func (v NullableCartCartData) IsSet() bool {
	return v.isSet
}

func (v *NullableCartCartData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartCartData(val *CartCartData) *NullableCartCartData {
	return &NullableCartCartData{value: val, isSet: true}
}

func (v NullableCartCartData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartCartData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


