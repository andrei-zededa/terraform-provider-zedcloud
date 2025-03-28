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

// AppProfileConfigSummary app profile config summary
//
// swagger:model AppProfileConfigSummary
type AppProfileConfigSummary struct {

	// list of app instance policies
	AppPolicies []*ProfileAppPolicy `json:"appPolicies"`

	// Detailed description of the app profile.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// unique id for app profile
	ID string `json:"id,omitempty"`

	// user defined app profile name
	Name string `json:"name,omitempty"`

	// list of network instance policies
	NetworkPolicies []*ProfileNetworkPolicy `json:"networkPolicies"`

	// user defined app profile title
	Title string `json:"title,omitempty"`

	// list of volume instance policies
	VolumePolicies []*ProfileVolumePolicy `json:"volumePolicies"`
}

// Validate validates this app profile config summary
func (m *AppProfileConfigSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkPolicies(formats); err != nil {
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

func (m *AppProfileConfigSummary) validateAppPolicies(formats strfmt.Registry) error {
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

func (m *AppProfileConfigSummary) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *AppProfileConfigSummary) validateNetworkPolicies(formats strfmt.Registry) error {
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

func (m *AppProfileConfigSummary) validateVolumePolicies(formats strfmt.Registry) error {
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

// ContextValidate validate this app profile config summary based on the context it is used
func (m *AppProfileConfigSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkPolicies(ctx, formats); err != nil {
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

func (m *AppProfileConfigSummary) contextValidateAppPolicies(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AppProfileConfigSummary) contextValidateNetworkPolicies(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AppProfileConfigSummary) contextValidateVolumePolicies(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AppProfileConfigSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppProfileConfigSummary) UnmarshalBinary(b []byte) error {
	var res AppProfileConfigSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
