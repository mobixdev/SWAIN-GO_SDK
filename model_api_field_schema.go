/*
SWAIN API

A powerful dynamic CRUD API engine that automatically generates RESTful endpoints for your data models SWAIN provides automatic CRUD operations, filtering, pagination, and sorting capabilities for any data model. Features: - Automatic REST API generation - Dynamic model support - Complex filtering and querying - Pagination and sorting - Swagger/OpenAPI documentation - Multiple database support (SQL & NoSQL)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiFieldSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiFieldSchema{}

// ApiFieldSchema Schema information for a model field
type ApiFieldSchema struct {
	// AutoIncrement indicates if the field is auto-increment
	AutoIncrement *bool `json:"autoIncrement,omitempty"`
	// DBColumnName is the column name in the database
	DbColumnName *string `json:"dbColumnName,omitempty"`
	// IsPrimaryKey indicates if the field is a primary key
	IsPrimaryKey *bool `json:"isPrimaryKey,omitempty"`
	// Name of the field (from JSON tag if available)
	Name *string `json:"name,omitempty"`
	// List of supported operators for this field
	Operators []string `json:"operators,omitempty"`
	// Type of the field (e.g., string, int)
	Type *string `json:"type,omitempty"`
}

// NewApiFieldSchema instantiates a new ApiFieldSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFieldSchema() *ApiFieldSchema {
	this := ApiFieldSchema{}
	return &this
}

// NewApiFieldSchemaWithDefaults instantiates a new ApiFieldSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFieldSchemaWithDefaults() *ApiFieldSchema {
	this := ApiFieldSchema{}
	return &this
}

// GetAutoIncrement returns the AutoIncrement field value if set, zero value otherwise.
func (o *ApiFieldSchema) GetAutoIncrement() bool {
	if o == nil || IsNil(o.AutoIncrement) {
		var ret bool
		return ret
	}
	return *o.AutoIncrement
}

// GetAutoIncrementOk returns a tuple with the AutoIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFieldSchema) GetAutoIncrementOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoIncrement) {
		return nil, false
	}
	return o.AutoIncrement, true
}

// HasAutoIncrement returns a boolean if a field has been set.
func (o *ApiFieldSchema) HasAutoIncrement() bool {
	if o != nil && !IsNil(o.AutoIncrement) {
		return true
	}

	return false
}

// SetAutoIncrement gets a reference to the given bool and assigns it to the AutoIncrement field.
func (o *ApiFieldSchema) SetAutoIncrement(v bool) {
	o.AutoIncrement = &v
}

// GetDbColumnName returns the DbColumnName field value if set, zero value otherwise.
func (o *ApiFieldSchema) GetDbColumnName() string {
	if o == nil || IsNil(o.DbColumnName) {
		var ret string
		return ret
	}
	return *o.DbColumnName
}

// GetDbColumnNameOk returns a tuple with the DbColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFieldSchema) GetDbColumnNameOk() (*string, bool) {
	if o == nil || IsNil(o.DbColumnName) {
		return nil, false
	}
	return o.DbColumnName, true
}

// HasDbColumnName returns a boolean if a field has been set.
func (o *ApiFieldSchema) HasDbColumnName() bool {
	if o != nil && !IsNil(o.DbColumnName) {
		return true
	}

	return false
}

// SetDbColumnName gets a reference to the given string and assigns it to the DbColumnName field.
func (o *ApiFieldSchema) SetDbColumnName(v string) {
	o.DbColumnName = &v
}

// GetIsPrimaryKey returns the IsPrimaryKey field value if set, zero value otherwise.
func (o *ApiFieldSchema) GetIsPrimaryKey() bool {
	if o == nil || IsNil(o.IsPrimaryKey) {
		var ret bool
		return ret
	}
	return *o.IsPrimaryKey
}

// GetIsPrimaryKeyOk returns a tuple with the IsPrimaryKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFieldSchema) GetIsPrimaryKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimaryKey) {
		return nil, false
	}
	return o.IsPrimaryKey, true
}

// HasIsPrimaryKey returns a boolean if a field has been set.
func (o *ApiFieldSchema) HasIsPrimaryKey() bool {
	if o != nil && !IsNil(o.IsPrimaryKey) {
		return true
	}

	return false
}

// SetIsPrimaryKey gets a reference to the given bool and assigns it to the IsPrimaryKey field.
func (o *ApiFieldSchema) SetIsPrimaryKey(v bool) {
	o.IsPrimaryKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiFieldSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFieldSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiFieldSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiFieldSchema) SetName(v string) {
	o.Name = &v
}

// GetOperators returns the Operators field value if set, zero value otherwise.
func (o *ApiFieldSchema) GetOperators() []string {
	if o == nil || IsNil(o.Operators) {
		var ret []string
		return ret
	}
	return o.Operators
}

// GetOperatorsOk returns a tuple with the Operators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFieldSchema) GetOperatorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Operators) {
		return nil, false
	}
	return o.Operators, true
}

// HasOperators returns a boolean if a field has been set.
func (o *ApiFieldSchema) HasOperators() bool {
	if o != nil && !IsNil(o.Operators) {
		return true
	}

	return false
}

// SetOperators gets a reference to the given []string and assigns it to the Operators field.
func (o *ApiFieldSchema) SetOperators(v []string) {
	o.Operators = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiFieldSchema) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFieldSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiFieldSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiFieldSchema) SetType(v string) {
	o.Type = &v
}

func (o ApiFieldSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiFieldSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoIncrement) {
		toSerialize["autoIncrement"] = o.AutoIncrement
	}
	if !IsNil(o.DbColumnName) {
		toSerialize["dbColumnName"] = o.DbColumnName
	}
	if !IsNil(o.IsPrimaryKey) {
		toSerialize["isPrimaryKey"] = o.IsPrimaryKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Operators) {
		toSerialize["operators"] = o.Operators
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableApiFieldSchema struct {
	value *ApiFieldSchema
	isSet bool
}

func (v NullableApiFieldSchema) Get() *ApiFieldSchema {
	return v.value
}

func (v *NullableApiFieldSchema) Set(val *ApiFieldSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFieldSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFieldSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFieldSchema(val *ApiFieldSchema) *NullableApiFieldSchema {
	return &NullableApiFieldSchema{value: val, isSet: true}
}

func (v NullableApiFieldSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFieldSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


