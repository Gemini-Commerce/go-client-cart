# GeminiCommerce\Cart\PromotionsAPI

All URIs are relative to *https://cart.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyCoupons**](PromotionsAPI.md#ApplyCoupons) | **Post** /cart.Cart/ApplyCoupons | Apply Coupons
[**RemoveCoupons**](PromotionsAPI.md#RemoveCoupons) | **Post** /cart.Cart/RemoveCoupons | Remove Coupons



## ApplyCoupons

> map[string]interface{} ApplyCoupons(ctx).Body(body).Execute()

Apply Coupons



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
	body := *openapiclient.NewCartApplyCouponsRequest() // CartApplyCouponsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromotionsAPI.ApplyCoupons(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromotionsAPI.ApplyCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyCoupons`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PromotionsAPI.ApplyCoupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartApplyCouponsRequest**](CartApplyCouponsRequest.md) |  | 

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


## RemoveCoupons

> map[string]interface{} RemoveCoupons(ctx).Body(body).Execute()

Remove Coupons



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
	body := *openapiclient.NewCartRemoveCouponsRequest() // CartRemoveCouponsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromotionsAPI.RemoveCoupons(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromotionsAPI.RemoveCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveCoupons`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PromotionsAPI.RemoveCoupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartRemoveCouponsRequest**](CartRemoveCouponsRequest.md) |  | 

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

