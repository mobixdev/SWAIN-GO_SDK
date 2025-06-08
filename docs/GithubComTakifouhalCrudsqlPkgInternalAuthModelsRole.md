# GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | Role name must be unique | [optional] 
**Permissions** | Pointer to [**[]GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission**](GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser**](GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser.md) |  | [optional] 

## Methods

### NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsRole

`func NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsRole() *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole`

NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsRole instantiates a new GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsRoleWithDefaults

`func NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsRoleWithDefaults() *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole`

NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsRoleWithDefaults instantiates a new GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetPermissions() []GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetPermissionsOk() (*[]GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) SetPermissions(v []GithubComTakifouhalCrudsqlPkgInternalAuthModelsPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsers

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetUsers() []GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) GetUsersOk() (*[]GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) SetUsers(v []GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


