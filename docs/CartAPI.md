# GeminiCommerce\Cart\CartAPI

All URIs are relative to *https://cart.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CartSetCustomPriceOnItems**](CartAPI.md#CartSetCustomPriceOnItems) | **Post** /cart.Cart/SetCustomPriceOnItems | 



## CartSetCustomPriceOnItems

> map[string]interface{} CartSetCustomPriceOnItems(ctx).Body(body).Execute()



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
	body := *openapiclient.NewCartSetCustomPriceOnItemsRequest() // CartSetCustomPriceOnItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPI.CartSetCustomPriceOnItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPI.CartSetCustomPriceOnItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSetCustomPriceOnItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CartAPI.CartSetCustomPriceOnItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartSetCustomPriceOnItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartSetCustomPriceOnItemsRequest**](CartSetCustomPriceOnItemsRequest.md) |  | 

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

