# GeminiCommerce\Cart\AddressesAPI

All URIs are relative to *https://cart.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetBillingAddress**](AddressesAPI.md#SetBillingAddress) | **Post** /cart.Cart/SetBillingAddress | Set Billing Address
[**SetShipmentAddress**](AddressesAPI.md#SetShipmentAddress) | **Post** /cart.Cart/SetShipmentAddress | Set Shipment Address



## SetBillingAddress

> map[string]interface{} SetBillingAddress(ctx).Body(body).Execute()

Set Billing Address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-cart"
)

func main() {
	body := *openapiclient.NewCartSetBillingAddressRequest() // CartSetBillingAddressRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.SetBillingAddress(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.SetBillingAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBillingAddress`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.SetBillingAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetBillingAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartSetBillingAddressRequest**](CartSetBillingAddressRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetShipmentAddress

> map[string]interface{} SetShipmentAddress(ctx).Body(body).Execute()

Set Shipment Address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-cart"
)

func main() {
	body := *openapiclient.NewCartSetShipmentAddressRequest() // CartSetShipmentAddressRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.SetShipmentAddress(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.SetShipmentAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetShipmentAddress`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.SetShipmentAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetShipmentAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartSetShipmentAddressRequest**](CartSetShipmentAddressRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

