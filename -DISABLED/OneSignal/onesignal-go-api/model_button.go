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

// Button struct for Button
type Button struct {
	Id string `json:"id"`
	Text *string `json:"text,omitempty"`
	Icon *string `json:"icon,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Button Button

// NewButton instantiates a new Button object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewButton(id string) *Button {
	this := Button{}
	this.Id = id
	return &this
}

// NewButtonWithDefaults instantiates a new Button object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewButtonWithDefaults() *Button {
	this := Button{}
	return &this
}

// GetId returns the Id field value
func (o *Button) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Button) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Button) SetId(v string) {
	o.Id = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Button) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Button) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Button) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Button) SetText(v string) {
	o.Text = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Button) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Button) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Button) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *Button) SetIcon(v string) {
	o.Icon = &v
}

func (o Button) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Button) UnmarshalJSON(bytes []byte) (err error) {
	varButton := _Button{}

	if err = json.Unmarshal(bytes, &varButton); err == nil {
		*o = Button(varButton)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "text")
		delete(additionalProperties, "icon")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableButton struct {
	value *Button
	isSet bool
}

func (v NullableButton) Get() *Button {
	return v.value
}

func (v *NullableButton) Set(val *Button) {
	v.value = val
	v.isSet = true
}

func (v NullableButton) IsSet() bool {
	return v.isSet
}

func (v *NullableButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableButton(val *Button) *NullableButton {
	return &NullableButton{value: val, isSet: true}
}

func (v NullableButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

