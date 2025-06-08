# ApiUpdateFilterPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expressions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**UpdateData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewApiUpdateFilterPayload

`func NewApiUpdateFilterPayload() *ApiUpdateFilterPayload`

NewApiUpdateFilterPayload instantiates a new ApiUpdateFilterPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUpdateFilterPayloadWithDefaults

`func NewApiUpdateFilterPayloadWithDefaults() *ApiUpdateFilterPayload`

NewApiUpdateFilterPayloadWithDefaults instantiates a new ApiUpdateFilterPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpressions

`func (o *ApiUpdateFilterPayload) GetExpressions() []map[string]interface{}`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *ApiUpdateFilterPayload) GetExpressionsOk() (*[]map[string]interface{}, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *ApiUpdateFilterPayload) SetExpressions(v []map[string]interface{})`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *ApiUpdateFilterPayload) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.

### GetUpdateData

`func (o *ApiUpdateFilterPayload) GetUpdateData() map[string]interface{}`

GetUpdateData returns the UpdateData field if non-nil, zero value otherwise.

### GetUpdateDataOk

`func (o *ApiUpdateFilterPayload) GetUpdateDataOk() (*map[string]interface{}, bool)`

GetUpdateDataOk returns a tuple with the UpdateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateData

`func (o *ApiUpdateFilterPayload) SetUpdateData(v map[string]interface{})`

SetUpdateData sets UpdateData field to given value.

### HasUpdateData

`func (o *ApiUpdateFilterPayload) HasUpdateData() bool`

HasUpdateData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


