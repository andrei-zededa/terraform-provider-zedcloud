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

// AssetGroupRead AssetGroupRead - refers to the response object for reading a asset group
//
// swagger:model AssetGroupRead
type AssetGroupRead struct {

	// asset count
	AssetCount int64 `json:"assetCount,omitempty"`

	// asset ids
	AssetIds *AssetIDs `json:"assetIds,omitempty"`

	// asset tags
	AssetTags *AssetTags `json:"assetTags,omitempty"`

	// Detailed description of the asset group.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// unique Id of the asset group.
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// User defined name of the asset group, unique across the enterprise. Once asset group is created, name can’t be changed.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name string `json:"name,omitempty"`

	// project id
	ProjectID string `json:"projectId,omitempty"`

	// User defined title of the asset group. Title can be changed at any time.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title string `json:"title,omitempty"`
}

// Validate validates this asset group read
func (m *AssetGroupRead) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssetTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetGroupRead) validateAssetIds(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetIds) { // not required
		return nil
	}

	if m.AssetIds != nil {
		if err := m.AssetIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assetIds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assetIds")
			}
			return err
		}
	}

	return nil
}

func (m *AssetGroupRead) validateAssetTags(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetTags) { // not required
		return nil
	}

	if m.AssetTags != nil {
		if err := m.AssetTags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assetTags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assetTags")
			}
			return err
		}
	}

	return nil
}

func (m *AssetGroupRead) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *AssetGroupRead) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *AssetGroupRead) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AssetGroupRead) validateTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MinLength("title", "body", m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", m.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this asset group read based on the context it is used
func (m *AssetGroupRead) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssetIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAssetTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetGroupRead) contextValidateAssetIds(ctx context.Context, formats strfmt.Registry) error {

	if m.AssetIds != nil {

		if swag.IsZero(m.AssetIds) { // not required
			return nil
		}

		if err := m.AssetIds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assetIds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assetIds")
			}
			return err
		}
	}

	return nil
}

func (m *AssetGroupRead) contextValidateAssetTags(ctx context.Context, formats strfmt.Registry) error {

	if m.AssetTags != nil {

		if swag.IsZero(m.AssetTags) { // not required
			return nil
		}

		if err := m.AssetTags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assetTags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assetTags")
			}
			return err
		}
	}

	return nil
}

func (m *AssetGroupRead) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetGroupRead) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetGroupRead) UnmarshalBinary(b []byte) error {
	var res AssetGroupRead
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
