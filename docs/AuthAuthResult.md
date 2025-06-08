# AuthAuthResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshToken** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** | Added: Include assigned role names | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | \&quot;local\&quot;, \&quot;github\&quot;, \&quot;google\&quot;, etc. | [optional] 
**User** | Pointer to [**GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser**](GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser.md) |  | [optional] 

## Methods

### NewAuthAuthResult

`func NewAuthAuthResult() *AuthAuthResult`

NewAuthAuthResult instantiates a new AuthAuthResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAuthResultWithDefaults

`func NewAuthAuthResultWithDefaults() *AuthAuthResult`

NewAuthAuthResultWithDefaults instantiates a new AuthAuthResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshToken

`func (o *AuthAuthResult) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AuthAuthResult) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AuthAuthResult) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AuthAuthResult) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRoles

`func (o *AuthAuthResult) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AuthAuthResult) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AuthAuthResult) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AuthAuthResult) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetToken

`func (o *AuthAuthResult) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthAuthResult) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthAuthResult) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthAuthResult) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *AuthAuthResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthAuthResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthAuthResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthAuthResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *AuthAuthResult) GetUser() GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthAuthResult) GetUserOk() (*GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthAuthResult) SetUser(v GithubComTakifouhalCrudsqlPkgInternalAuthModelsUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthAuthResult) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


