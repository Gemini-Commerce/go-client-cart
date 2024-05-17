# GeminiCommerce\Cart\PaymentsAPI

All URIs are relative to *https://cart.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPaymentMethods**](PaymentsAPI.md#ListPaymentMethods) | **Post** /cart.Cart/ListPaymentMethods | List Payment Methods
[**SetSetPayments**](PaymentsAPI.md#SetSetPayments) | **Post** /cart.Cart/SetPayments | Set SetPayments



## ListPaymentMethods

> CartListPaymentMethodsResponse ListPaymentMethods(ctx).Body(body).Execute()

List Payment Methods



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
	body := *openapiclient.NewCartListPaymentMethodsRequest() // CartListPaymentMethodsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ListPaymentMethods(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ListPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentMethods`: CartListPaymentMethodsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ListPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartListPaymentMethodsRequest**](CartListPaymentMethodsRequest.md) |  | 

### Return type

[**CartListPaymentMethodsResponse**](CartListPaymentMethodsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSetPayments

> CartSetPaymentsResponse SetSetPayments(ctx).Body(body).Execute()

Set SetPayments



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
	body := *openapiclient.NewCartSetPaymentsRequest() // CartSetPaymentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.SetSetPayments(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.SetSetPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSetPayments`: CartSetPaymentsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.SetSetPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSetPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartSetPaymentsRequest**](CartSetPaymentsRequest.md) |  | 

### Return type

[**CartSetPaymentsResponse**](CartSetPaymentsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

