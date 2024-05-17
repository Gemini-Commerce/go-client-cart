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

// checks if the CartCreateCartRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartCreateCartRequest{}

// CartCreateCartRequest struct for CartCreateCartRequest
type CartCreateCartRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Channel *string `json:"channel,omitempty"`
	Market *string `json:"market,omitempty"`
	Locale *string `json:"locale,omitempty"`
	CustomerGrn *string `json:"customerGrn,omitempty"`
	Currency *CartCurrency `json:"currency,omitempty"`
}

// NewCartCreateCartRequest instantiates a new CartCreateCartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartCreateCartRequest() *CartCreateCartRequest {
	this := CartCreateCartRequest{}
	var currency CartCurrency = CARTCURRENCY_XXX
	this.Currency = &currency
	return &this
}

// NewCartCreateCartRequestWithDefaults instantiates a new CartCreateCartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartCreateCartRequestWithDefaults() *CartCreateCartRequest {
	this := CartCreateCartRequest{}
	var currency CartCurrency = CARTCURRENCY_XXX
	this.Currency = &currency
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CartCreateCartRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCreateCartRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CartCreateCartRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CartCreateCartRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *CartCreateCartRequest) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCreateCartRequest) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *CartCreateCartRequest) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *CartCreateCartRequest) SetChannel(v string) {
	o.Channel = &v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *CartCreateCartRequest) GetMarket() string {
	if o == nil || IsNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCreateCartRequest) GetMarketOk() (*string, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *CartCreateCartRequest) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *CartCreateCartRequest) SetMarket(v string) {
	o.Market = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *CartCreateCartRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCreateCartRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *CartCreateCartRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *CartCreateCartRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *CartCreateCartRequest) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCreateCartRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// HasCustomerGrn returns a boolean if a field has been set.
func (o *CartCreateCartRequest) HasCustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *CartCreateCartRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CartCreateCartRequest) GetCurrency() CartCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret CartCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartCreateCartRequest) GetCurrencyOk() (*CartCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CartCreateCartRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CartCurrency and assigns it to the Currency field.
func (o *CartCreateCartRequest) SetCurrency(v CartCurrency) {
	o.Currency = &v
}

func (o CartCreateCartRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartCreateCartRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
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
	if !IsNil(o.CustomerGrn) {
		toSerialize["customerGrn"] = o.CustomerGrn
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableCartCreateCartRequest struct {
	value *CartCreateCartRequest
	isSet bool
}

func (v NullableCartCreateCartRequest) Get() *CartCreateCartRequest {
	return v.value
}

func (v *NullableCartCreateCartRequest) Set(val *CartCreateCartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCartCreateCartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCartCreateCartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartCreateCartRequest(val *CartCreateCartRequest) *NullableCartCreateCartRequest {
	return &NullableCartCreateCartRequest{value: val, isSet: true}
}

func (v NullableCartCreateCartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartCreateCartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


