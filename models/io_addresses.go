// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoAddresses io addresses
//
// swagger:model IoAddresses
type IoAddresses struct {

	// mac address
	MacAddress string `json:"macAddress,omitempty"`
}

// Validate validates this io addresses
func (m *IoAddresses) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io addresses based on context it is used
func (m *IoAddresses) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoAddresses) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoAddresses) UnmarshalBinary(b []byte) error {
	var res IoAddresses
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
