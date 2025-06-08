# \DynamicAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModelCountPost**](DynamicAPI.md#ModelCountPost) | **Post** /{model}/count | Count entities by filter
[**ModelFilterPost**](DynamicAPI.md#ModelFilterPost) | **Post** /{model}/filter | Filter entities
[**ModelGet**](DynamicAPI.md#ModelGet) | **Get** /{model} | List entities
[**ModelIdDelete**](DynamicAPI.md#ModelIdDelete) | **Delete** /{model}/{id} | Delete an entity
[**ModelIdGet**](DynamicAPI.md#ModelIdGet) | **Get** /{model}/{id} | Get an entity by ID
[**ModelIdPut**](DynamicAPI.md#ModelIdPut) | **Put** /{model}/{id} | Update an entity
[**ModelPost**](DynamicAPI.md#ModelPost) | **Post** /{model} | Create a new entity



## ModelCountPost

> map[string]interface{} ModelCountPost(ctx, model).Filter(filter).Execute()

Count entities by filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model name
	filter := *openapiclient.NewApiFilterPayload() // ApiFilterPayload | Filter conditions (optional) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelCountPost(context.Background(), model).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelCountPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelCountPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelCountPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelCountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | [**ApiFilterPayload**](ApiFilterPayload.md) | Filter conditions (optional) | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelFilterPost

> QueryFilterResponse ModelFilterPost(ctx, model).Filter(filter).Page(page).PageSize(pageSize).Execute()

Filter entities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | model name
	filter := *openapiclient.NewQueryFilter() // QueryFilter | Filter conditions, including optional 'sort' array (e.g., {\\
	page := int32(56) // int32 | Page number (optional) (default to 1)
	pageSize := int32(56) // int32 | Items per page (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelFilterPost(context.Background(), model).Filter(filter).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelFilterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelFilterPost`: QueryFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelFilterPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | model name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelFilterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | [**QueryFilter**](QueryFilter.md) | Filter conditions, including optional &#39;sort&#39; array (e.g., {\\ | 
 **page** | **int32** | Page number | [default to 1]
 **pageSize** | **int32** | Items per page | [default to 10]

### Return type

[**QueryFilterResponse**](QueryFilterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelGet

> QueryFilterResponse ModelGet(ctx, model).Page(page).PageSize(pageSize).Execute()

List entities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	page := int32(56) // int32 | Page number (optional)
	pageSize := int32(56) // int32 | Items per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelGet(context.Background(), model).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelGet`: QueryFilterResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number | 
 **pageSize** | **int32** | Items per page | 

### Return type

[**QueryFilterResponse**](QueryFilterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelIdDelete

> map[string]interface{} ModelIdDelete(ctx, model, id).Execute()

Delete an entity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	id := "id_example" // string | Entity ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelIdDelete(context.Background(), model, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 
**id** | **string** | Entity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelIdGet

> map[string]interface{} ModelIdGet(ctx, model, id).Execute()

Get an entity by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	id := "id_example" // string | Entity ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelIdGet(context.Background(), model, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelIdGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 
**id** | **string** | Entity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelIdPut

> map[string]interface{} ModelIdPut(ctx, model, id).Entity(entity).Execute()

Update an entity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	id := "id_example" // string | Entity ID
	entity := map[string]interface{}{ ... } // map[string]interface{} | Entity Data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelIdPut(context.Background(), model, id).Entity(entity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelIdPut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 
**id** | **string** | Entity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entity** | **map[string]interface{}** | Entity Data | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelPost

> map[string]interface{} ModelPost(ctx, model).Entity(entity).Execute()

Create a new entity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	model := "model_example" // string | Model Name
	entity := map[string]interface{}{ ... } // map[string]interface{} | Entity Data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAPI.ModelPost(context.Background(), model).Entity(entity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAPI.ModelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DynamicAPI.ModelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entity** | **map[string]interface{}** | Entity Data | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

