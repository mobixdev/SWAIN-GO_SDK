# \SchemaAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModelSchemaGet**](SchemaAPI.md#ModelSchemaGet) | **Get** /{model}/schema | Get model schema
[**ModelsGet**](SchemaAPI.md#ModelsGet) | **Get** /models | List all available models



## ModelSchemaGet

> ApiModelSchema ModelSchemaGet(ctx, model).Execute()

Get model schema



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaAPI.ModelSchemaGet(context.Background(), model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.ModelSchemaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelSchemaGet`: ApiModelSchema
	fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.ModelSchemaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model name | 

### Other Parameters

Other parameters are passed through a pointer to a apiModelSchemaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiModelSchema**](ApiModelSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModelsGet

> ApiModelsResponse ModelsGet(ctx).Execute()

List all available models



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchemaAPI.ModelsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchemaAPI.ModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModelsGet`: ApiModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `SchemaAPI.ModelsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiModelsGetRequest struct via the builder pattern


### Return type

[**ApiModelsResponse**](ApiModelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

