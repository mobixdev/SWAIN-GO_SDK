# QueryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**QueryAggregationSpec**](QueryAggregationSpec.md) | @Description Aggregations can specify functions like COUNT, SUM, GROUP BY, etc. @Description If not provided, no aggregations will be performed. @Description Used for data analysis and reporting queries | [optional] 
**Expressions** | Pointer to **[]map[string]interface{}** | @Description Array of expressions to filter entities @Description Each expression can be a FieldExpression, LogicalExpression, or RelationshipExpression @Description If empty, no filtering will be applied | [optional] 
**Projections** | Pointer to **[]string** | @Description Array of field names to select from the main entity @Description If empty, all fields will be selected @Description Example: [\&quot;id\&quot;, \&quot;name\&quot;, \&quot;email\&quot;] to select only those fields | [optional] 
**Sort** | Pointer to [**[]QuerySortSpec**](QuerySortSpec.md) | @Description Array of sort specifications to order the results @Description Example: [{\&quot;field\&quot;: \&quot;name\&quot;, \&quot;direction\&quot;: \&quot;asc\&quot;}, {\&quot;field\&quot;: \&quot;createdAt\&quot;, \&quot;direction\&quot;: \&quot;desc\&quot;}] @Description Field names can be struct field names or database column names. Direction is case-insensitive (&#39;asc&#39; or &#39;desc&#39;). | [optional] 

## Methods

### NewQueryFilter

`func NewQueryFilter() *QueryFilter`

NewQueryFilter instantiates a new QueryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryFilterWithDefaults

`func NewQueryFilterWithDefaults() *QueryFilter`

NewQueryFilterWithDefaults instantiates a new QueryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *QueryFilter) GetAggregations() QueryAggregationSpec`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *QueryFilter) GetAggregationsOk() (*QueryAggregationSpec, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *QueryFilter) SetAggregations(v QueryAggregationSpec)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *QueryFilter) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetExpressions

`func (o *QueryFilter) GetExpressions() []map[string]interface{}`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *QueryFilter) GetExpressionsOk() (*[]map[string]interface{}, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *QueryFilter) SetExpressions(v []map[string]interface{})`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *QueryFilter) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.

### GetProjections

`func (o *QueryFilter) GetProjections() []string`

GetProjections returns the Projections field if non-nil, zero value otherwise.

### GetProjectionsOk

`func (o *QueryFilter) GetProjectionsOk() (*[]string, bool)`

GetProjectionsOk returns a tuple with the Projections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjections

`func (o *QueryFilter) SetProjections(v []string)`

SetProjections sets Projections field to given value.

### HasProjections

`func (o *QueryFilter) HasProjections() bool`

HasProjections returns a boolean if a field has been set.

### GetSort

`func (o *QueryFilter) GetSort() []QuerySortSpec`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QueryFilter) GetSortOk() (*[]QuerySortSpec, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QueryFilter) SetSort(v []QuerySortSpec)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QueryFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


