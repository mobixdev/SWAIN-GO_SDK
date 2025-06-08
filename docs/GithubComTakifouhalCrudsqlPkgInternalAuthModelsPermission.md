# GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedActions** | Pointer to **string** | Comma-separated actions: \&quot;read,create,update,delete\&quot;, \&quot;*\&quot; | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EntityName** | Pointer to **string** | e.g., \&quot;users\&quot;, \&quot;orders\&quot;, \&quot;*\&quot; | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Roles** | Pointer to [**[]GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole**](GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole.md) |  | [optional] 
**ScopeExpression** | Pointer to **[]int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission

`func NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission() *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission`

NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission instantiates a new GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsPermissionWithDefaults

`func NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsPermissionWithDefaults() *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission`

NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsPermissionWithDefaults instantiates a new GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedActions

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetAllowedActions() string`

GetAllowedActions returns the AllowedActions field if non-nil, zero value otherwise.

### GetAllowedActionsOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetAllowedActionsOk() (*string, bool)`

GetAllowedActionsOk returns a tuple with the AllowedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedActions

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) SetAllowedActions(v string)`

SetAllowedActions sets AllowedActions field to given value.

### HasAllowedActions

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) HasAllowedActions() bool`

HasAllowedActions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntityName

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetId

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoles

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetRoles() []GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetRolesOk() (*[]GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) SetRoles(v []GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetScopeExpression

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetScopeExpression() []int32`

GetScopeExpression returns the ScopeExpression field if non-nil, zero value otherwise.

### GetScopeExpressionOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetScopeExpressionOk() (*[]int32, bool)`

GetScopeExpressionOk returns a tuple with the ScopeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeExpression

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) SetScopeExpression(v []int32)`

SetScopeExpression sets ScopeExpression field to given value.

### HasScopeExpression

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) HasScopeExpression() bool`

HasScopeExpression returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


