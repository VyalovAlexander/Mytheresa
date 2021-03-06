// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ErrorResponseItem error response item
//
// swagger:model ErrorResponseItem
type ErrorResponseItem struct {

	// Error code
	// Example: 1001
	// Minimum: 0
	Code *int64 `json:"code,omitempty"`

	// Fields with errors
	// Example: name
	Field string `json:"field,omitempty"`

	// Error message
	// Example: Some error message
	Message string `json:"message,omitempty"`
}

// Validate validates this error response item
func (m *ErrorResponseItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponseItem) validateCode(formats strfmt.Registry) error {
	if swag.IsZero(m.Code) { // not required
		return nil
	}

	if err := validate.MinimumInt("code", "body", *m.Code, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error response item based on context it is used
func (m *ErrorResponseItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponseItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponseItem) UnmarshalBinary(b []byte) error {
	var res ErrorResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
