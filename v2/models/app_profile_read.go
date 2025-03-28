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
)

// AppProfileRead app profile read
//
// swagger:model AppProfileRead
type AppProfileRead struct {

	// list of app instance policies
	AppPolicies []*ProfileAppPolicy `json:"appPolicies"`

	// unique id for an App profile
	ID string `json:"id,omitempty"`

	// latest version of app profile
	LatestVersion int64 `json:"latestVersion,omitempty"`

	// user defined name for the App profile
	Name string `json:"name,omitempty"`

	// list of network instance policies
	NetworkPolicies []*ProfileNetworkPolicy `json:"networkPolicies"`

	// status of the app profile
	Status *AppProfileStatusType `json:"status,omitempty"`

	// user defined title for the app profile
	Title string `json:"title,omitempty"`

	// list of volume instance policies
	VolumePolicies []*ProfileVolumePolicy `json:"volumePolicies"`
}

// Validate validates this app profile read
func (m *AppProfileRead) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppProfileRead) validateAppPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.AppPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.AppPolicies); i++ {
		if swag.IsZero(m.AppPolicies[i]) { // not required
			continue
		}

		if m.AppPolicies[i] != nil {
			if err := m.AppPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppProfileRead) validateNetworkPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkPolicies); i++ {
		if swag.IsZero(m.NetworkPolicies[i]) { // not required
			continue
		}

		if m.NetworkPolicies[i] != nil {
			if err := m.NetworkPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networkPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networkPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppProfileRead) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *AppProfileRead) validateVolumePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumePolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumePolicies); i++ {
		if swag.IsZero(m.VolumePolicies[i]) { // not required
			continue
		}

		if m.VolumePolicies[i] != nil {
			if err := m.VolumePolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumePolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumePolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this app profile read based on the context it is used
func (m *AppProfileRead) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppProfileRead) contextValidateAppPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppPolicies); i++ {

		if m.AppPolicies[i] != nil {

			if swag.IsZero(m.AppPolicies[i]) { // not required
				return nil
			}

			if err := m.AppPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppProfileRead) contextValidateNetworkPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetworkPolicies); i++ {

		if m.NetworkPolicies[i] != nil {

			if swag.IsZero(m.NetworkPolicies[i]) { // not required
				return nil
			}

			if err := m.NetworkPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networkPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networkPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppProfileRead) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *AppProfileRead) contextValidateVolumePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumePolicies); i++ {

		if m.VolumePolicies[i] != nil {

			if swag.IsZero(m.VolumePolicies[i]) { // not required
				return nil
			}

			if err := m.VolumePolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumePolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumePolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppProfileRead) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppProfileRead) UnmarshalBinary(b []byte) error {
	var res AppProfileRead
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
