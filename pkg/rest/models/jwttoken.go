// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Jwttoken a JWT token to carry identity information
// swagger:model jwttoken
type Jwttoken struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// expires in
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// scope
	Scope []string `json:"scope"`

	// token type
	TokenType *string `json:"token_type,omitempty"`

	// jwttoken additional properties
	JwttokenAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *Jwttoken) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// access token
		AccessToken string `json:"access_token,omitempty"`

		// expires in
		ExpiresIn int64 `json:"expires_in,omitempty"`

		// refresh token
		RefreshToken string `json:"refresh_token,omitempty"`

		// scope
		Scope []string `json:"scope"`

		// token type
		TokenType *string `json:"token_type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv Jwttoken

	rcv.AccessToken = stage1.AccessToken

	rcv.ExpiresIn = stage1.ExpiresIn

	rcv.RefreshToken = stage1.RefreshToken

	rcv.Scope = stage1.Scope

	rcv.TokenType = stage1.TokenType

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "access_token")

	delete(stage2, "expires_in")

	delete(stage2, "refresh_token")

	delete(stage2, "scope")

	delete(stage2, "token_type")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.JwttokenAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m Jwttoken) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// access token
		AccessToken string `json:"access_token,omitempty"`

		// expires in
		ExpiresIn int64 `json:"expires_in,omitempty"`

		// refresh token
		RefreshToken string `json:"refresh_token,omitempty"`

		// scope
		Scope []string `json:"scope"`

		// token type
		TokenType *string `json:"token_type,omitempty"`
	}

	stage1.AccessToken = m.AccessToken

	stage1.ExpiresIn = m.ExpiresIn

	stage1.RefreshToken = m.RefreshToken

	stage1.Scope = m.Scope

	stage1.TokenType = m.TokenType

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.JwttokenAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.JwttokenAdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this jwttoken
func (m *Jwttoken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScope(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Jwttoken) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Jwttoken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Jwttoken) UnmarshalBinary(b []byte) error {
	var res Jwttoken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
