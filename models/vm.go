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

// VM vm detail
//
// # Virtual machine memory configuration
//
// swagger:model vm
type VM struct {

	// Enable CpuPinning
	CPUPinningEnabled bool `json:"cpuPinningEnabled,omitempty"`

	// CPUs
	// Required: true
	Cpus *int64 `json:"cpus"`

	// Memory
	// Required: true
	Memory *int64 `json:"memory"`

	// Hardware Virtualization
	// Required: true
	Mode *HvMode `json:"mode"`

	// VNC
	// Required: true
	Vnc *bool `json:"vnc"`

	// VNC display
	// Read Only: true
	VncDisplay int64 `json:"vncDisplay,omitempty"`
}

// Validate validates this vm
func (m *VM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVnc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VM) validateCpus(formats strfmt.Registry) error {

	if err := validate.Required("cpus", "body", m.Cpus); err != nil {
		return err
	}

	return nil
}

func (m *VM) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *VM) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *VM) validateVnc(formats strfmt.Registry) error {

	if err := validate.Required("vnc", "body", m.Vnc); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vm based on the context it is used
func (m *VM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVncDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VM) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if m.Mode != nil {
		if err := m.Mode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *VM) contextValidateVncDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "vncDisplay", "body", int64(m.VncDisplay)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VM) UnmarshalBinary(b []byte) error {
	var res VM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
