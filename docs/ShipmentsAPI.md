# GeminiCommerce\Cart\ShipmentsAPI

All URIs are relative to *https://cart.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListShipmentMethods**](ShipmentsAPI.md#ListShipmentMethods) | **Post** /cart.Cart/ListShipmentMethods | List Shipment Methods
[**SetShipments**](ShipmentsAPI.md#SetShipments) | **Post** /cart.Cart/SetShipments | Set Shipments



## ListShipmentMethods

> CartListShipmentMethodsResponse ListShipmentMethods(ctx).Body(body).Execute()

List Shipment Methods



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
	body := *openapiclient.NewCartListShipmentMethodsRequest() // CartListShipmentMethodsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentsAPI.ListShipmentMethods(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsAPI.ListShipmentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListShipmentMethods`: CartListShipmentMethodsResponse
	fmt.Fprintf(os.Stdout, "Response from `ShipmentsAPI.ListShipmentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListShipmentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartListShipmentMethodsRequest**](CartListShipmentMethodsRequest.md) |  | 

### Return type

[**CartListShipmentMethodsResponse**](CartListShipmentMethodsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetShipments

> map[string]interface{} SetShipments(ctx).Body(body).Execute()

Set Shipments



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
	body := *openapiclient.NewCartSetShipmentsRequest() // CartSetShipmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentsAPI.SetShipments(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsAPI.SetShipments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetShipments`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ShipmentsAPI.SetShipments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetShipmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartSetShipmentsRequest**](CartSetShipmentsRequest.md) |  | 

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

