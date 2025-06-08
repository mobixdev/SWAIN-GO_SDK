# ApiRelationshipSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForeignKey** | Pointer to **string** | ForeignKey is the foreign key in the relationship | [optional] 
**IsPolymorphic** | Pointer to **bool** | IsPolymorphic indicates if the relationship is polymorphic | [optional] 
**IsSlice** | Pointer to **bool** | IsSlice indicates if the relationship is a slice | [optional] 
**JoinTable** | Pointer to **string** | JoinTable is the join table in the relationship | [optional] 
**JsonName** | Pointer to **string** | JsonName of the relationship (from JSON tag if available) | [optional] 
**Name** | Pointer to **string** | Name of the relationship (from JSON tag if available) | [optional] 
**PolymorphicType** | Pointer to **string** | PolymorphicType is the type of the polymorphic relationship | [optional] 
**References** | Pointer to **string** | References are the references in the relationship | [optional] 
**RelatedModel** | Pointer to **string** | Name of the related model | [optional] 
**RelatedTable** | Pointer to **string** | Name of the related table | [optional] 
**Type** | Pointer to **string** | Type of relationship (hasOne, hasMany, belongsTo, manyToMany) | [optional] 

## Methods

### NewApiRelationshipSchema

`func NewApiRelationshipSchema() *ApiRelationshipSchema`

NewApiRelationshipSchema instantiates a new ApiRelationshipSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRelationshipSchemaWithDefaults

`func NewApiRelationshipSchemaWithDefaults() *ApiRelationshipSchema`

NewApiRelationshipSchemaWithDefaults instantiates a new ApiRelationshipSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForeignKey

`func (o *ApiRelationshipSchema) GetForeignKey() string`

GetForeignKey returns the ForeignKey field if non-nil, zero value otherwise.

### GetForeignKeyOk

`func (o *ApiRelationshipSchema) GetForeignKeyOk() (*string, bool)`

GetForeignKeyOk returns a tuple with the ForeignKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignKey

`func (o *ApiRelationshipSchema) SetForeignKey(v string)`

SetForeignKey sets ForeignKey field to given value.

### HasForeignKey

`func (o *ApiRelationshipSchema) HasForeignKey() bool`

HasForeignKey returns a boolean if a field has been set.

### GetIsPolymorphic

`func (o *ApiRelationshipSchema) GetIsPolymorphic() bool`

GetIsPolymorphic returns the IsPolymorphic field if non-nil, zero value otherwise.

### GetIsPolymorphicOk

`func (o *ApiRelationshipSchema) GetIsPolymorphicOk() (*bool, bool)`

GetIsPolymorphicOk returns a tuple with the IsPolymorphic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPolymorphic

`func (o *ApiRelationshipSchema) SetIsPolymorphic(v bool)`

SetIsPolymorphic sets IsPolymorphic field to given value.

### HasIsPolymorphic

`func (o *ApiRelationshipSchema) HasIsPolymorphic() bool`

HasIsPolymorphic returns a boolean if a field has been set.

### GetIsSlice

`func (o *ApiRelationshipSchema) GetIsSlice() bool`

GetIsSlice returns the IsSlice field if non-nil, zero value otherwise.

### GetIsSliceOk

`func (o *ApiRelationshipSchema) GetIsSliceOk() (*bool, bool)`

GetIsSliceOk returns a tuple with the IsSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSlice

`func (o *ApiRelationshipSchema) SetIsSlice(v bool)`

SetIsSlice sets IsSlice field to given value.

### HasIsSlice

`func (o *ApiRelationshipSchema) HasIsSlice() bool`

HasIsSlice returns a boolean if a field has been set.

### GetJoinTable

`func (o *ApiRelationshipSchema) GetJoinTable() string`

GetJoinTable returns the JoinTable field if non-nil, zero value otherwise.

### GetJoinTableOk

`func (o *ApiRelationshipSchema) GetJoinTableOk() (*string, bool)`

GetJoinTableOk returns a tuple with the JoinTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTable

`func (o *ApiRelationshipSchema) SetJoinTable(v string)`

SetJoinTable sets JoinTable field to given value.

### HasJoinTable

`func (o *ApiRelationshipSchema) HasJoinTable() bool`

HasJoinTable returns a boolean if a field has been set.

### GetJsonName

`func (o *ApiRelationshipSchema) GetJsonName() string`

GetJsonName returns the JsonName field if non-nil, zero value otherwise.

### GetJsonNameOk

`func (o *ApiRelationshipSchema) GetJsonNameOk() (*string, bool)`

GetJsonNameOk returns a tuple with the JsonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonName

`func (o *ApiRelationshipSchema) SetJsonName(v string)`

SetJsonName sets JsonName field to given value.

### HasJsonName

`func (o *ApiRelationshipSchema) HasJsonName() bool`

HasJsonName returns a boolean if a field has been set.

### GetName

`func (o *ApiRelationshipSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiRelationshipSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiRelationshipSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiRelationshipSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolymorphicType

`func (o *ApiRelationshipSchema) GetPolymorphicType() string`

GetPolymorphicType returns the PolymorphicType field if non-nil, zero value otherwise.

### GetPolymorphicTypeOk

`func (o *ApiRelationshipSchema) GetPolymorphicTypeOk() (*string, bool)`

GetPolymorphicTypeOk returns a tuple with the PolymorphicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolymorphicType

`func (o *ApiRelationshipSchema) SetPolymorphicType(v string)`

SetPolymorphicType sets PolymorphicType field to given value.

### HasPolymorphicType

`func (o *ApiRelationshipSchema) HasPolymorphicType() bool`

HasPolymorphicType returns a boolean if a field has been set.

### GetReferences

`func (o *ApiRelationshipSchema) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiRelationshipSchema) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiRelationshipSchema) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiRelationshipSchema) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedModel

`func (o *ApiRelationshipSchema) GetRelatedModel() string`

GetRelatedModel returns the RelatedModel field if non-nil, zero value otherwise.

### GetRelatedModelOk

`func (o *ApiRelationshipSchema) GetRelatedModelOk() (*string, bool)`

GetRelatedModelOk returns a tuple with the RelatedModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedModel

`func (o *ApiRelationshipSchema) SetRelatedModel(v string)`

SetRelatedModel sets RelatedModel field to given value.

### HasRelatedModel

`func (o *ApiRelationshipSchema) HasRelatedModel() bool`

HasRelatedModel returns a boolean if a field has been set.

### GetRelatedTable

`func (o *ApiRelationshipSchema) GetRelatedTable() string`

GetRelatedTable returns the RelatedTable field if non-nil, zero value otherwise.

### GetRelatedTableOk

`func (o *ApiRelationshipSchema) GetRelatedTableOk() (*string, bool)`

GetRelatedTableOk returns a tuple with the RelatedTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTable

`func (o *ApiRelationshipSchema) SetRelatedTable(v string)`

SetRelatedTable sets RelatedTable field to given value.

### HasRelatedTable

`func (o *ApiRelationshipSchema) HasRelatedTable() bool`

HasRelatedTable returns a boolean if a field has been set.

### GetType

`func (o *ApiRelationshipSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiRelationshipSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiRelationshipSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiRelationshipSchema) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


