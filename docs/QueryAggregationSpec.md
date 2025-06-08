# QueryAggregationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expressions** | Pointer to **[]map[string]interface{}** | @Description Optional aggregator-based conditions to be applied in a HAVING clause (if using SQL) @Description These expressions may reference aggregator functions like SUM(...) or COUNT(...) | [optional] 
**Functions** | Pointer to [**[]QueryAggregateFunctionSpec**](QueryAggregateFunctionSpec.md) | @Description A list of aggregation functions (e.g., COUNT, SUM, MIN, MAX) to apply @Description Each function specifies the type, field, and optional alias | [optional] 
**GroupBy** | Pointer to **[]string** | @Description Fields for grouping results (e.g., by \&quot;status\&quot; or [\&quot;status\&quot;,\&quot;category\&quot;]) @Description Similar to SQL GROUP BY clause | [optional] 

## Methods

### NewQueryAggregationSpec

`func NewQueryAggregationSpec() *QueryAggregationSpec`

NewQueryAggregationSpec instantiates a new QueryAggregationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAggregationSpecWithDefaults

`func NewQueryAggregationSpecWithDefaults() *QueryAggregationSpec`

NewQueryAggregationSpecWithDefaults instantiates a new QueryAggregationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpressions

`func (o *QueryAggregationSpec) GetExpressions() []map[string]interface{}`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *QueryAggregationSpec) GetExpressionsOk() (*[]map[string]interface{}, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *QueryAggregationSpec) SetExpressions(v []map[string]interface{})`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *QueryAggregationSpec) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.

### GetFunctions

`func (o *QueryAggregationSpec) GetFunctions() []QueryAggregateFunctionSpec`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *QueryAggregationSpec) GetFunctionsOk() (*[]QueryAggregateFunctionSpec, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *QueryAggregationSpec) SetFunctions(v []QueryAggregateFunctionSpec)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *QueryAggregationSpec) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetGroupBy

`func (o *QueryAggregationSpec) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *QueryAggregationSpec) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *QueryAggregationSpec) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *QueryAggregationSpec) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


