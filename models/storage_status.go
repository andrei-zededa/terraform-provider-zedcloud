// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageStatus storage status
//
// swagger:model StorageStatus
type StorageStatus struct {

	// mount path
	MountPath string `json:"mountPath,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size m b
	SizeMB string `json:"sizeMB,omitempty"`
}

// Validate validates this storage status
func (m *StorageStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage status based on context it is used
func (m *StorageStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageStatus) UnmarshalBinary(b []byte) error {
	var res StorageStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
