# \AuthenticationAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthLoginPost**](AuthenticationAPI.md#AuthLoginPost) | **Post** /auth/login | User login



## AuthLoginPost

> AuthAuthResult AuthLoginPost(ctx).Credentials(credentials).Execute()

User login



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
	credentials := *openapiclient.NewAuthLoginRequest() // AuthLoginRequest | Login credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthLoginPost(context.Background()).Credentials(credentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthLoginPost`: AuthAuthResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**AuthLoginRequest**](AuthLoginRequest.md) | Login credentials | 

### Return type

[**AuthAuthResult**](AuthAuthResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

