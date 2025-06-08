# GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **map[string]interface{}** |  | [optional] 
**Anonymous** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**[]GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole**](GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole.md) |  | [optional] 

## Methods

### NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsUser

`func NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsUser() *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser`

NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsUser instantiates a new GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsUserWithDefaults

`func NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsUserWithDefaults() *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser`

NewGithubComTakifouhalCrudsqlPkgInternalAuthModelsUserWithDefaults instantiates a new GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAnonymous

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetEmail

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastLogin

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetName

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetRoles() []GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) GetRolesOk() (*[]GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) SetRoles(v []GithubComTakifouhalCrudsqlPkgInternalAuthModelsRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


