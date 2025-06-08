# \BulkOperationsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModelFilterDelete**](BulkOperationsAPI.md#ModelFilterDelete) | **Delete** /{model}/filter | Mass delete entities by filter
[**ModelFilterPut**](BulkOperationsAPI.md#ModelFilterPut) | **Put** /{model}/filter | Mass update entities by filter



## ModelFilterDelete

> ApiSuccessResponse ModelFilterDelete(ctx, model).Filter(filter).Execute()

Mass delete entities by filter



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
	filter := *openapiclient.NewApiFilterPayload() // ApiFilterPayload | Filter conditions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkOperationsAPI.ModelFilterDelete(context.Background(), model).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkOperationsAPI.ModelFilterDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelFilterDelete`: ApiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkOperationsAPI.ModelFilterDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelFilterDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | [**ApiFilterPayload**](ApiFilterPayload.md) | Filter conditions | 

### Return type

[**ApiSuccessResponse**](ApiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelFilterPut

> ApiSuccessResponse ModelFilterPut(ctx, model).Filter(filter).Execute()

Mass update entities by filter



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
	filter := *openapiclient.NewApiUpdateFilterPayload() // ApiUpdateFilterPayload | Filter conditions and update data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkOperationsAPI.ModelFilterPut(context.Background(), model).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkOperationsAPI.ModelFilterPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelFilterPut`: ApiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkOperationsAPI.ModelFilterPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelFilterPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | [**ApiUpdateFilterPayload**](ApiUpdateFilterPayload.md) | Filter conditions and update data | 

### Return type

[**ApiSuccessResponse**](ApiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

