/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.0.1
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// CreateSegmentConflictResponse struct for CreateSegmentConflictResponse
type CreateSegmentConflictResponse struct {
	Success *bool `json:"success,omitempty"`
	Errors []string `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSegmentConflictResponse CreateSegmentConflictResponse

// NewCreateSegmentConflictResponse instantiates a new CreateSegmentConflictResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSegmentConflictResponse() *CreateSegmentConflictResponse {
	this := CreateSegmentConflictResponse{}
	return &this
}

// NewCreateSegmentConflictResponseWithDefaults instantiates a new CreateSegmentConflictResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSegmentConflictResponseWithDefaults() *CreateSegmentConflictResponse {
	this := CreateSegmentConflictResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateSegmentConflictResponse) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSegmentConflictResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CreateSegmentConflictResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateSegmentConflictResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateSegmentConflictResponse) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSegmentConflictResponse) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateSegmentConflictResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *CreateSegmentConflictResponse) SetErrors(v []string) {
	o.Errors = v
}

func (o CreateSegmentConflictResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateSegmentConflictResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreateSegmentConflictResponse := _CreateSegmentConflictResponse{}

	if err = json.Unmarshal(bytes, &varCreateSegmentConflictResponse); err == nil {
		*o = CreateSegmentConflictResponse(varCreateSegmentConflictResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "success")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSegmentConflictResponse struct {
	value *CreateSegmentConflictResponse
	isSet bool
}

func (v NullableCreateSegmentConflictResponse) Get() *CreateSegmentConflictResponse {
	return v.value
}

func (v *NullableCreateSegmentConflictResponse) Set(val *CreateSegmentConflictResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSegmentConflictResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSegmentConflictResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSegmentConflictResponse(val *CreateSegmentConflictResponse) *NullableCreateSegmentConflictResponse {
	return &NullableCreateSegmentConflictResponse{value: val, isSet: true}
}

func (v NullableCreateSegmentConflictResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSegmentConflictResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

