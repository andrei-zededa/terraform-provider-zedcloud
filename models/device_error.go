// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceError DeviceError is used to store the error details of the hardware.
//
// # DeviceError store the error occurred at the device side
//
// swagger:model DeviceError
type DeviceError struct {

	// Description of the error
	// Required: true
	Description *string `json:"description"`

	// objects referenced by the description or retry_condition
	// Required: true
	Entities []*DeviceEntity `json:"entities"`

	// condition for retry
	RetryCondition string `json:"retryCondition,omitempty"`

	// Severity of the error
	// Required: true
	Severity *Severity `json:"severity"`

	// Timestamp at which error had occurred
	// Required: true
	Timestamp interface{} `json:"timestamp"`
}

// Validate validates this device error
func (m *DeviceError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceError) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *DeviceError) validateEntities(formats strfmt.Registry) error {

	if err := validate.Required("entities", "body", m.Entities); err != nil {
		return err
	}

	for i := 0; i < len(m.Entities); i++ {
		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {
			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceError) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	if m.Severity != nil {
		if err := m.Severity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("severity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("severity")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceError) validateTimestamp(formats strfmt.Registry) error {

	if m.Timestamp == nil {
		return errors.Required("timestamp", "body", nil)
	}

	return nil
}

// ContextValidate validate this device error based on the context it is used
func (m *DeviceError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceError) contextValidateEntities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entities); i++ {

		if m.Entities[i] != nil {
			if err := m.Entities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceError) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if m.Severity != nil {
		if err := m.Severity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("severity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("severity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceError) UnmarshalBinary(b []byte) error {
	var res DeviceError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
