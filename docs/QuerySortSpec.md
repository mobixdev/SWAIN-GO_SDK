# QuerySortSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **string** | @Description The sort direction: \&quot;asc\&quot; for ascending, \&quot;desc\&quot; for descending @Description Case-insensitive, so \&quot;ASC\&quot;, \&quot;asc\&quot;, and \&quot;Asc\&quot; are all valid | [optional] 
**Field** | Pointer to **string** | @Description The name of the field to sort by @Description Can be a struct field name (e.g., \&quot;FirstName\&quot;), JSON field name (from json tag), @Description or database column name (e.g., \&quot;first_name\&quot;) | [optional] 

## Methods

### NewQuerySortSpec

`func NewQuerySortSpec() *QuerySortSpec`

NewQuerySortSpec instantiates a new QuerySortSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySortSpecWithDefaults

`func NewQuerySortSpecWithDefaults() *QuerySortSpec`

NewQuerySortSpecWithDefaults instantiates a new QuerySortSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *QuerySortSpec) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QuerySortSpec) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QuerySortSpec) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *QuerySortSpec) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetField

`func (o *QuerySortSpec) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *QuerySortSpec) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *QuerySortSpec) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *QuerySortSpec) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


