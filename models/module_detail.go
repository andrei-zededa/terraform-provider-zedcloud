// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModuleDetail Module detail
//
// Azure module specific details
// Example: {"twinDetail":"IHsKICAgICAgIm5hbWUiOiAiU2VuZEludGVydmFsIiwKICAgICAgInZhbHVlIjogIjUiCiAgICB9"}
//
// swagger:model ModuleDetail
type ModuleDetail struct {

	// Extra information to module to make configuration easier
	Environment map[string]string `json:"environment,omitempty"`

	// Type of modules
	ModuleType *ModuleType `json:"moduleType,omitempty"`

	// Send messages between modules or send messages from modules to iot hub
	Routes map[string]string `json:"routes,omitempty"`

	// Base64 encoded module twin details, desired properties of the module will be updated to reflect these values
	TwinDetail string `json:"twinDetail,omitempty"`
}

// Validate validates this module detail
func (m *ModuleDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModuleType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModuleDetail) validateModuleType(formats strfmt.Registry) error {
	if swag.IsZero(m.ModuleType) { // not required
		return nil
	}

	if m.ModuleType != nil {
		if err := m.ModuleType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moduleType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moduleType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this module detail based on the context it is used
func (m *ModuleDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModuleType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModuleDetail) contextValidateModuleType(ctx context.Context, formats strfmt.Registry) error {

	if m.ModuleType != nil {
		if err := m.ModuleType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moduleType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("moduleType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModuleDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModuleDetail) UnmarshalBinary(b []byte) error {
	var res ModuleDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
