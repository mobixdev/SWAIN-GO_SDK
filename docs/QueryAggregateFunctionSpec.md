# QueryAggregateFunctionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | @Description Optional alias for the result of this aggregation @Description If omitted, the function name may be used as an alias @Description Example: \&quot;total_amount\&quot; for SUM(amount) | [optional] 
**Field** | Pointer to **string** | @Description The field on which the aggregation function is applied @Description For COUNT(*), use \&quot;*\&quot; or leave empty | [optional] 
**Type** | Pointer to **string** | @Description The aggregation function type (e.g., COUNT, SUM, MIN, MAX) @Description Common types: COUNT, SUM, AVG, MIN, MAX | [optional] 

## Methods

### NewQueryAggregateFunctionSpec

`func NewQueryAggregateFunctionSpec() *QueryAggregateFunctionSpec`

NewQueryAggregateFunctionSpec instantiates a new QueryAggregateFunctionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAggregateFunctionSpecWithDefaults

`func NewQueryAggregateFunctionSpecWithDefaults() *QueryAggregateFunctionSpec`

NewQueryAggregateFunctionSpecWithDefaults instantiates a new QueryAggregateFunctionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *QueryAggregateFunctionSpec) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *QueryAggregateFunctionSpec) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *QueryAggregateFunctionSpec) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *QueryAggregateFunctionSpec) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetField

`func (o *QueryAggregateFunctionSpec) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *QueryAggregateFunctionSpec) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *QueryAggregateFunctionSpec) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *QueryAggregateFunctionSpec) HasField() bool`

HasField returns a boolean if a field has been set.

### GetType

`func (o *QueryAggregateFunctionSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryAggregateFunctionSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryAggregateFunctionSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueryAggregateFunctionSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


