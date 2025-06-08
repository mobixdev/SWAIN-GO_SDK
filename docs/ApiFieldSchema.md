# ApiFieldSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoIncrement** | Pointer to **bool** | AutoIncrement indicates if the field is auto-increment | [optional] 
**DbColumnName** | Pointer to **string** | DBColumnName is the column name in the database | [optional] 
**IsPrimaryKey** | Pointer to **bool** | IsPrimaryKey indicates if the field is a primary key | [optional] 
**Name** | Pointer to **string** | Name of the field (from JSON tag if available) | [optional] 
**Operators** | Pointer to **[]string** | List of supported operators for this field | [optional] 
**Type** | Pointer to **string** | Type of the field (e.g., string, int) | [optional] 

## Methods

### NewApiFieldSchema

`func NewApiFieldSchema() *ApiFieldSchema`

NewApiFieldSchema instantiates a new ApiFieldSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFieldSchemaWithDefaults

`func NewApiFieldSchemaWithDefaults() *ApiFieldSchema`

NewApiFieldSchemaWithDefaults instantiates a new ApiFieldSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoIncrement

`func (o *ApiFieldSchema) GetAutoIncrement() bool`

GetAutoIncrement returns the AutoIncrement field if non-nil, zero value otherwise.

### GetAutoIncrementOk

`func (o *ApiFieldSchema) GetAutoIncrementOk() (*bool, bool)`

GetAutoIncrementOk returns a tuple with the AutoIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIncrement

`func (o *ApiFieldSchema) SetAutoIncrement(v bool)`

SetAutoIncrement sets AutoIncrement field to given value.

### HasAutoIncrement

`func (o *ApiFieldSchema) HasAutoIncrement() bool`

HasAutoIncrement returns a boolean if a field has been set.

### GetDbColumnName

`func (o *ApiFieldSchema) GetDbColumnName() string`

GetDbColumnName returns the DbColumnName field if non-nil, zero value otherwise.

### GetDbColumnNameOk

`func (o *ApiFieldSchema) GetDbColumnNameOk() (*string, bool)`

GetDbColumnNameOk returns a tuple with the DbColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbColumnName

`func (o *ApiFieldSchema) SetDbColumnName(v string)`

SetDbColumnName sets DbColumnName field to given value.

### HasDbColumnName

`func (o *ApiFieldSchema) HasDbColumnName() bool`

HasDbColumnName returns a boolean if a field has been set.

### GetIsPrimaryKey

`func (o *ApiFieldSchema) GetIsPrimaryKey() bool`

GetIsPrimaryKey returns the IsPrimaryKey field if non-nil, zero value otherwise.

### GetIsPrimaryKeyOk

`func (o *ApiFieldSchema) GetIsPrimaryKeyOk() (*bool, bool)`

GetIsPrimaryKeyOk returns a tuple with the IsPrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryKey

`func (o *ApiFieldSchema) SetIsPrimaryKey(v bool)`

SetIsPrimaryKey sets IsPrimaryKey field to given value.

### HasIsPrimaryKey

`func (o *ApiFieldSchema) HasIsPrimaryKey() bool`

HasIsPrimaryKey returns a boolean if a field has been set.

### GetName

`func (o *ApiFieldSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiFieldSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiFieldSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiFieldSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperators

`func (o *ApiFieldSchema) GetOperators() []string`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *ApiFieldSchema) GetOperatorsOk() (*[]string, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *ApiFieldSchema) SetOperators(v []string)`

SetOperators sets Operators field to given value.

### HasOperators

`func (o *ApiFieldSchema) HasOperators() bool`

HasOperators returns a boolean if a field has been set.

### GetType

`func (o *ApiFieldSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiFieldSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiFieldSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiFieldSchema) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


