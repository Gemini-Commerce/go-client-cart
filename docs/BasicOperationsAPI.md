# GeminiCommerce\Cart\BasicOperationsAPI

All URIs are relative to *https://cart.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProducts**](BasicOperationsAPI.md#AddProducts) | **Post** /cart.Cart/AddProducts | Add Products
[**CreateCart**](BasicOperationsAPI.md#CreateCart) | **Post** /cart.Cart/CreateCart | Create Cart
[**GetActiveCartByCustomer**](BasicOperationsAPI.md#GetActiveCartByCustomer) | **Post** /cart.Cart/GetActiveCartByCustomer | Get Active Cart By Customer
[**GetCart**](BasicOperationsAPI.md#GetCart) | **Post** /cart.Cart/GetCart | Get Cart
[**ListCarts**](BasicOperationsAPI.md#ListCarts) | **Post** /cart.Cart/ListCarts | List Carts
[**MergeCarts**](BasicOperationsAPI.md#MergeCarts) | **Post** /cart.Cart/MergeCarts | Merge Carts
[**RemoveProducts**](BasicOperationsAPI.md#RemoveProducts) | **Post** /cart.Cart/RemoveProducts | Remove Products
[**SetAdditionalInfo**](BasicOperationsAPI.md#SetAdditionalInfo) | **Post** /cart.Cart/SetAdditionalInfo | Set Additional Info
[**SetNotes**](BasicOperationsAPI.md#SetNotes) | **Post** /cart.Cart/SetNotes | Set Notes
[**TriggerRealignment**](BasicOperationsAPI.md#TriggerRealignment) | **Post** /cart.Cart/TriggerRealignment | Trigger Realignment
[**UpdateCart**](BasicOperationsAPI.md#UpdateCart) | **Post** /cart.Cart/UpdateCart | Update Cart
[**UpdateProducts**](BasicOperationsAPI.md#UpdateProducts) | **Post** /cart.Cart/UpdateProducts | Update Products



## AddProducts

> CartAddProductsResponse AddProducts(ctx).Body(body).Execute()

Add Products



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
	body := *openapiclient.NewCartAddProductsRequest() // CartAddProductsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.AddProducts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.AddProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProducts`: CartAddProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.AddProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartAddProductsRequest**](CartAddProductsRequest.md) |  | 

### Return type

[**CartAddProductsResponse**](CartAddProductsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCart

> CartCreateCartResponse CreateCart(ctx).Body(body).Execute()

Create Cart



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
	body := *openapiclient.NewCartCreateCartRequest() // CartCreateCartRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.CreateCart(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.CreateCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCart`: CartCreateCartResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.CreateCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartCreateCartRequest**](CartCreateCartRequest.md) |  | 

### Return type

[**CartCreateCartResponse**](CartCreateCartResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveCartByCustomer

> CartGetActiveCartByCustomerResponse GetActiveCartByCustomer(ctx).Body(body).Execute()

Get Active Cart By Customer



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
	body := *openapiclient.NewCartGetActiveCartByCustomerRequest() // CartGetActiveCartByCustomerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.GetActiveCartByCustomer(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.GetActiveCartByCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveCartByCustomer`: CartGetActiveCartByCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.GetActiveCartByCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveCartByCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartGetActiveCartByCustomerRequest**](CartGetActiveCartByCustomerRequest.md) |  | 

### Return type

[**CartGetActiveCartByCustomerResponse**](CartGetActiveCartByCustomerResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCart

> CartGetCartResponse GetCart(ctx).Body(body).Execute()

Get Cart



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
	body := *openapiclient.NewCartGetCartRequest() // CartGetCartRequest | Get cart request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.GetCart(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.GetCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCart`: CartGetCartResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.GetCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartGetCartRequest**](CartGetCartRequest.md) | Get cart request | 

### Return type

[**CartGetCartResponse**](CartGetCartResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCarts

> CartListCartsResponse ListCarts(ctx).Body(body).Execute()

List Carts



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
	body := *openapiclient.NewCartListCartsRequest() // CartListCartsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.ListCarts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.ListCarts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCarts`: CartListCartsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.ListCarts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartListCartsRequest**](CartListCartsRequest.md) |  | 

### Return type

[**CartListCartsResponse**](CartListCartsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeCarts

> CartMergeCartsResponse MergeCarts(ctx).Body(body).Execute()

Merge Carts



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
	body := *openapiclient.NewCartMergeCartsRequest() // CartMergeCartsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.MergeCarts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.MergeCarts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeCarts`: CartMergeCartsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.MergeCarts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMergeCartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartMergeCartsRequest**](CartMergeCartsRequest.md) |  | 

### Return type

[**CartMergeCartsResponse**](CartMergeCartsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProducts

> map[string]interface{} RemoveProducts(ctx).Body(body).Execute()

Remove Products



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
	body := *openapiclient.NewCartRemoveProductsRequest() // CartRemoveProductsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.RemoveProducts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.RemoveProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveProducts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.RemoveProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartRemoveProductsRequest**](CartRemoveProductsRequest.md) |  | 

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


## SetAdditionalInfo

> map[string]interface{} SetAdditionalInfo(ctx).Body(body).Execute()

Set Additional Info



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
	body := *openapiclient.NewCartSetAdditionalInfoRequest() // CartSetAdditionalInfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.SetAdditionalInfo(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.SetAdditionalInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAdditionalInfo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.SetAdditionalInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAdditionalInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartSetAdditionalInfoRequest**](CartSetAdditionalInfoRequest.md) |  | 

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


## SetNotes

> map[string]interface{} SetNotes(ctx).Body(body).Execute()

Set Notes



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
	body := *openapiclient.NewCartSetNotesRequest() // CartSetNotesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.SetNotes(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.SetNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNotes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.SetNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartSetNotesRequest**](CartSetNotesRequest.md) |  | 

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


## TriggerRealignment

> CartTriggerRealignmentResponse TriggerRealignment(ctx).Body(body).Execute()

Trigger Realignment



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
	body := *openapiclient.NewCartTriggerRealignmentRequest() // CartTriggerRealignmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.TriggerRealignment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.TriggerRealignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerRealignment`: CartTriggerRealignmentResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.TriggerRealignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerRealignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartTriggerRealignmentRequest**](CartTriggerRealignmentRequest.md) |  | 

### Return type

[**CartTriggerRealignmentResponse**](CartTriggerRealignmentResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCart

> map[string]interface{} UpdateCart(ctx).Body(body).Execute()

Update Cart



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
	body := *openapiclient.NewCartUpdateCartRequest() // CartUpdateCartRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UpdateCart(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UpdateCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCart`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UpdateCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartUpdateCartRequest**](CartUpdateCartRequest.md) |  | 

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


## UpdateProducts

> map[string]interface{} UpdateProducts(ctx).Body(body).Execute()

Update Products



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
	body := *openapiclient.NewCartUpdateProductsRequest() // CartUpdateProductsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UpdateProducts(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UpdateProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProducts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UpdateProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CartUpdateProductsRequest**](CartUpdateProductsRequest.md) |  | 

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

