// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateAppResponse openpitrix create app response
// swagger:model openpitrixCreateAppResponse
type OpenpitrixCreateAppResponse struct {

	// app
	App *OpenpitrixApp `json:"app,omitempty"`
}

// Validate validates this openpitrix create app response
func (m *OpenpitrixCreateAppResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixCreateAppResponse) validateApp(formats strfmt.Registry) error {

	if swag.IsZero(m.App) { // not required
		return nil
	}

	if m.App != nil {

		if err := m.App.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateAppResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateAppResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateAppResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
