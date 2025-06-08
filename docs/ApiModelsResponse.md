# ApiModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | Pointer to **map[string]string** | Models is a map of model names to their corresponding table names example: {\&quot;User\&quot;: \&quot;users\&quot;, \&quot;Post\&quot;: \&quot;posts\&quot;} | [optional] 

## Methods

### NewApiModelsResponse

`func NewApiModelsResponse() *ApiModelsResponse`

NewApiModelsResponse instantiates a new ApiModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiModelsResponseWithDefaults

`func NewApiModelsResponseWithDefaults() *ApiModelsResponse`

NewApiModelsResponseWithDefaults instantiates a new ApiModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *ApiModelsResponse) GetModels() map[string]string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ApiModelsResponse) GetModelsOk() (*map[string]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ApiModelsResponse) SetModels(v map[string]string)`

SetModels sets Models field to given value.

### HasModels

`func (o *ApiModelsResponse) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


