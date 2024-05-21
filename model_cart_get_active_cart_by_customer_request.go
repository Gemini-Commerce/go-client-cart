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

// checks if the CartGetActiveCartByCustomerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartGetActiveCartByCustomerRequest{}

// CartGetActiveCartByCustomerRequest struct for CartGetActiveCartByCustomerRequest
type CartGetActiveCartByCustomerRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	CustomerGrn *string `json:"customerGrn,omitempty"`
	Channel *string `json:"channel,omitempty"`
	Market *string `json:"market,omitempty"`
	Currency *CartCurrency `json:"currency,omitempty"`
}

// NewCartGetActiveCartByCustomerRequest instantiates a new CartGetActiveCartByCustomerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartGetActiveCartByCustomerRequest() *CartGetActiveCartByCustomerRequest {
	this := CartGetActiveCartByCustomerRequest{}
	var currency CartCurrency = CARTCURRENCY_XXX
	this.Currency = &currency
	return &this
}

// NewCartGetActiveCartByCustomerRequestWithDefaults instantiates a new CartGetActiveCartByCustomerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartGetActiveCartByCustomerRequestWithDefaults() *CartGetActiveCartByCustomerRequest {
	this := CartGetActiveCartByCustomerRequest{}
	var currency CartCurrency = CARTCURRENCY_XXX
	this.Currency = &currency
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CartGetActiveCartByCustomerRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartGetActiveCartByCustomerRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// IsSetTenantId returns a boolean if a field has been set.
func (o *CartGetActiveCartByCustomerRequest) IsSetTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CartGetActiveCartByCustomerRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCustomerGrn returns the CustomerGrn field value if set, zero value otherwise.
func (o *CartGetActiveCartByCustomerRequest) GetCustomerGrn() string {
	if o == nil || IsNil(o.CustomerGrn) {
		var ret string
		return ret
	}
	return *o.CustomerGrn
}

// GetCustomerGrnOk returns a tuple with the CustomerGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartGetActiveCartByCustomerRequest) GetCustomerGrnOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerGrn) {
		return nil, false
	}
	return o.CustomerGrn, true
}

// IsSetCustomerGrn returns a boolean if a field has been set.
func (o *CartGetActiveCartByCustomerRequest) IsSetCustomerGrn() bool {
	if o != nil && !IsNil(o.CustomerGrn) {
		return true
	}

	return false
}

// SetCustomerGrn gets a reference to the given string and assigns it to the CustomerGrn field.
func (o *CartGetActiveCartByCustomerRequest) SetCustomerGrn(v string) {
	o.CustomerGrn = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *CartGetActiveCartByCustomerRequest) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartGetActiveCartByCustomerRequest) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// IsSetChannel returns a boolean if a field has been set.
func (o *CartGetActiveCartByCustomerRequest) IsSetChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *CartGetActiveCartByCustomerRequest) SetChannel(v string) {
	o.Channel = &v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *CartGetActiveCartByCustomerRequest) GetMarket() string {
	if o == nil || IsNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartGetActiveCartByCustomerRequest) GetMarketOk() (*string, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// IsSetMarket returns a boolean if a field has been set.
func (o *CartGetActiveCartByCustomerRequest) IsSetMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *CartGetActiveCartByCustomerRequest) SetMarket(v string) {
	o.Market = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CartGetActiveCartByCustomerRequest) GetCurrency() CartCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret CartCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartGetActiveCartByCustomerRequest) GetCurrencyOk() (*CartCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// IsSetCurrency returns a boolean if a field has been set.
func (o *CartGetActiveCartByCustomerRequest) IsSetCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CartCurrency and assigns it to the Currency field.
func (o *CartGetActiveCartByCustomerRequest) SetCurrency(v CartCurrency) {
	o.Currency = &v
}

func (o CartGetActiveCartByCustomerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartGetActiveCartByCustomerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.CustomerGrn) {
		toSerialize["customerGrn"] = o.CustomerGrn
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableCartGetActiveCartByCustomerRequest struct {
	value *CartGetActiveCartByCustomerRequest
	isSet bool
}

func (v NullableCartGetActiveCartByCustomerRequest) Get() *CartGetActiveCartByCustomerRequest {
	return v.value
}

func (v *NullableCartGetActiveCartByCustomerRequest) Set(val *CartGetActiveCartByCustomerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCartGetActiveCartByCustomerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCartGetActiveCartByCustomerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartGetActiveCartByCustomerRequest(val *CartGetActiveCartByCustomerRequest) *NullableCartGetActiveCartByCustomerRequest {
	return &NullableCartGetActiveCartByCustomerRequest{value: val, isSet: true}
}

func (v NullableCartGetActiveCartByCustomerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartGetActiveCartByCustomerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


