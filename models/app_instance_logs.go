// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppInstanceLogs App Instance Logs configuration
//
// # App Instance Logs configuration
//
// swagger:model AppInstanceLogs
type AppInstanceLogs struct {

	// Flags to enable / disable sending of logs generated by app instance to zedcloud
	// Required: true
	Access *string `json:"access"`
}

// Validate validates this app instance logs
func (m *AppInstanceLogs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstanceLogs) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("access", "body", m.Access); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this app instance logs based on context it is used
func (m *AppInstanceLogs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppInstanceLogs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstanceLogs) UnmarshalBinary(b []byte) error {
	var res AppInstanceLogs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
