/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"os"
	"time"
)

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	// None
	Integer *int32 `json:"integer,omitempty"`
	// None
	Int32 *int32 `json:"int32,omitempty"`
	// None
	Int64 *int64 `json:"int64,omitempty"`
	// None
	Number float32 `json:"number"`
	// None
	Float *float32 `json:"float,omitempty"`
	// None
	Double float64 `json:"double"`
	// None
	String *string `json:"string,omitempty"`
	// None
	PatternWithoutDelimiter string `json:"pattern_without_delimiter"`
	// None
	Byte string `json:"byte"`
	// None
	Binary **os.File `json:"binary,omitempty"`
	// None
	Date *string `json:"date,omitempty"`
	// None
	DateTime *time.Time `json:"dateTime,omitempty"`
	// None
	Password *string `json:"password,omitempty"`
	// None
	Callback             *string `json:"callback,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineObject3 InlineObject3

// NewInlineObject3 instantiates a new InlineObject3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject3(number float32, double float64, patternWithoutDelimiter string, byte_ string) *InlineObject3 {
	this := InlineObject3{}
	this.Number = number
	this.Double = double
	this.PatternWithoutDelimiter = patternWithoutDelimiter
	this.Byte = byte_
	return &this
}

// NewInlineObject3WithDefaults instantiates a new InlineObject3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject3WithDefaults() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// GetInteger returns the Integer field value if set, zero value otherwise.
func (o *InlineObject3) GetInteger() int32 {
	if o == nil || o.Integer == nil {
		var ret int32
		return ret
	}
	return *o.Integer
}

// GetIntegerOk returns a tuple with the Integer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetIntegerOk() (*int32, bool) {
	if o == nil || o.Integer == nil {
		return nil, false
	}
	return o.Integer, true
}

// HasInteger returns a boolean if a field has been set.
func (o *InlineObject3) HasInteger() bool {
	if o != nil && o.Integer != nil {
		return true
	}

	return false
}

// SetInteger gets a reference to the given int32 and assigns it to the Integer field.
func (o *InlineObject3) SetInteger(v int32) {
	o.Integer = &v
}

// GetInt32 returns the Int32 field value if set, zero value otherwise.
func (o *InlineObject3) GetInt32() int32 {
	if o == nil || o.Int32 == nil {
		var ret int32
		return ret
	}
	return *o.Int32
}

// GetInt32Ok returns a tuple with the Int32 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetInt32Ok() (*int32, bool) {
	if o == nil || o.Int32 == nil {
		return nil, false
	}
	return o.Int32, true
}

// HasInt32 returns a boolean if a field has been set.
func (o *InlineObject3) HasInt32() bool {
	if o != nil && o.Int32 != nil {
		return true
	}

	return false
}

// SetInt32 gets a reference to the given int32 and assigns it to the Int32 field.
func (o *InlineObject3) SetInt32(v int32) {
	o.Int32 = &v
}

// GetInt64 returns the Int64 field value if set, zero value otherwise.
func (o *InlineObject3) GetInt64() int64 {
	if o == nil || o.Int64 == nil {
		var ret int64
		return ret
	}
	return *o.Int64
}

// GetInt64Ok returns a tuple with the Int64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetInt64Ok() (*int64, bool) {
	if o == nil || o.Int64 == nil {
		return nil, false
	}
	return o.Int64, true
}

// HasInt64 returns a boolean if a field has been set.
func (o *InlineObject3) HasInt64() bool {
	if o != nil && o.Int64 != nil {
		return true
	}

	return false
}

// SetInt64 gets a reference to the given int64 and assigns it to the Int64 field.
func (o *InlineObject3) SetInt64(v int64) {
	o.Int64 = &v
}

// GetNumber returns the Number field value
func (o *InlineObject3) GetNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *InlineObject3) SetNumber(v float32) {
	o.Number = v
}

// GetFloat returns the Float field value if set, zero value otherwise.
func (o *InlineObject3) GetFloat() float32 {
	if o == nil || o.Float == nil {
		var ret float32
		return ret
	}
	return *o.Float
}

// GetFloatOk returns a tuple with the Float field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetFloatOk() (*float32, bool) {
	if o == nil || o.Float == nil {
		return nil, false
	}
	return o.Float, true
}

// HasFloat returns a boolean if a field has been set.
func (o *InlineObject3) HasFloat() bool {
	if o != nil && o.Float != nil {
		return true
	}

	return false
}

// SetFloat gets a reference to the given float32 and assigns it to the Float field.
func (o *InlineObject3) SetFloat(v float32) {
	o.Float = &v
}

// GetDouble returns the Double field value
func (o *InlineObject3) GetDouble() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Double
}

// GetDoubleOk returns a tuple with the Double field value
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetDoubleOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Double, true
}

// SetDouble sets field value
func (o *InlineObject3) SetDouble(v float64) {
	o.Double = v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *InlineObject3) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetStringOk() (*string, bool) {
	if o == nil || o.String == nil {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *InlineObject3) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *InlineObject3) SetString(v string) {
	o.String = &v
}

// GetPatternWithoutDelimiter returns the PatternWithoutDelimiter field value
func (o *InlineObject3) GetPatternWithoutDelimiter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatternWithoutDelimiter
}

// GetPatternWithoutDelimiterOk returns a tuple with the PatternWithoutDelimiter field value
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetPatternWithoutDelimiterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatternWithoutDelimiter, true
}

// SetPatternWithoutDelimiter sets field value
func (o *InlineObject3) SetPatternWithoutDelimiter(v string) {
	o.PatternWithoutDelimiter = v
}

// GetByte returns the Byte field value
func (o *InlineObject3) GetByte() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Byte
}

// GetByteOk returns a tuple with the Byte field value
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetByteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Byte, true
}

// SetByte sets field value
func (o *InlineObject3) SetByte(v string) {
	o.Byte = v
}

// GetBinary returns the Binary field value if set, zero value otherwise.
func (o *InlineObject3) GetBinary() *os.File {
	if o == nil || o.Binary == nil {
		var ret *os.File
		return ret
	}
	return *o.Binary
}

// GetBinaryOk returns a tuple with the Binary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetBinaryOk() (**os.File, bool) {
	if o == nil || o.Binary == nil {
		return nil, false
	}
	return o.Binary, true
}

// HasBinary returns a boolean if a field has been set.
func (o *InlineObject3) HasBinary() bool {
	if o != nil && o.Binary != nil {
		return true
	}

	return false
}

// SetBinary gets a reference to the given *os.File and assigns it to the Binary field.
func (o *InlineObject3) SetBinary(v *os.File) {
	o.Binary = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *InlineObject3) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *InlineObject3) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *InlineObject3) SetDate(v string) {
	o.Date = &v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *InlineObject3) GetDateTime() time.Time {
	if o == nil || o.DateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetDateTimeOk() (*time.Time, bool) {
	if o == nil || o.DateTime == nil {
		return nil, false
	}
	return o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *InlineObject3) HasDateTime() bool {
	if o != nil && o.DateTime != nil {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given time.Time and assigns it to the DateTime field.
func (o *InlineObject3) SetDateTime(v time.Time) {
	o.DateTime = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *InlineObject3) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *InlineObject3) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *InlineObject3) SetPassword(v string) {
	o.Password = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *InlineObject3) GetCallback() string {
	if o == nil || o.Callback == nil {
		var ret string
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetCallbackOk() (*string, bool) {
	if o == nil || o.Callback == nil {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *InlineObject3) HasCallback() bool {
	if o != nil && o.Callback != nil {
		return true
	}

	return false
}

// SetCallback gets a reference to the given string and assigns it to the Callback field.
func (o *InlineObject3) SetCallback(v string) {
	o.Callback = &v
}

func (o InlineObject3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Integer != nil {
		toSerialize["integer"] = o.Integer
	}
	if o.Int32 != nil {
		toSerialize["int32"] = o.Int32
	}
	if o.Int64 != nil {
		toSerialize["int64"] = o.Int64
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if o.Float != nil {
		toSerialize["float"] = o.Float
	}
	if true {
		toSerialize["double"] = o.Double
	}
	if o.String != nil {
		toSerialize["string"] = o.String
	}
	if true {
		toSerialize["pattern_without_delimiter"] = o.PatternWithoutDelimiter
	}
	if true {
		toSerialize["byte"] = o.Byte
	}
	if o.Binary != nil {
		toSerialize["binary"] = o.Binary
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.DateTime != nil {
		toSerialize["dateTime"] = o.DateTime
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Callback != nil {
		toSerialize["callback"] = o.Callback
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineObject3) UnmarshalJSON(bytes []byte) (err error) {
	varInlineObject3 := _InlineObject3{}

	if err = json.Unmarshal(bytes, &varInlineObject3); err == nil {
		*o = InlineObject3(varInlineObject3)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "integer")
		delete(additionalProperties, "int32")
		delete(additionalProperties, "int64")
		delete(additionalProperties, "number")
		delete(additionalProperties, "float")
		delete(additionalProperties, "double")
		delete(additionalProperties, "string")
		delete(additionalProperties, "pattern_without_delimiter")
		delete(additionalProperties, "byte")
		delete(additionalProperties, "binary")
		delete(additionalProperties, "date")
		delete(additionalProperties, "dateTime")
		delete(additionalProperties, "password")
		delete(additionalProperties, "callback")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineObject3 struct {
	value *InlineObject3
	isSet bool
}

func (v NullableInlineObject3) Get() *InlineObject3 {
	return v.value
}

func (v *NullableInlineObject3) Set(val *InlineObject3) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject3) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject3(val *InlineObject3) *NullableInlineObject3 {
	return &NullableInlineObject3{value: val, isSet: true}
}

func (v NullableInlineObject3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
