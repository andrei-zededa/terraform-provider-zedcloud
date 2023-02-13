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

// NetInstList net inst list
//
// swagger:model NetInstList
type NetInstList struct {

	// cfg list
	CfgList []*NetworkInstance `json:"cfgList"`

	// list
	List []*NetInstShortConfig `json:"list"`

	// next
	Next *Cursor `json:"next,omitempty"`

	// Summary information about netinstance list records by addressing type.
	SummaryByAddressType *Summary `json:"summaryByAddressType,omitempty"`

	// Summary information about netinstance list records by network instance kind.
	SummaryByKind *Summary `json:"summaryByKind,omitempty"`
}

// Validate validates this net inst list
func (m *NetInstList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCfgList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByAddressType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaryByKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetInstList) validateCfgList(formats strfmt.Registry) error {
	if swag.IsZero(m.CfgList) { // not required
		return nil
	}

	for i := 0; i < len(m.CfgList); i++ {
		if swag.IsZero(m.CfgList[i]) { // not required
			continue
		}

		if m.CfgList[i] != nil {
			if err := m.CfgList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cfgList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cfgList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetInstList) validateList(formats strfmt.Registry) error {
	if swag.IsZero(m.List) { // not required
		return nil
	}

	for i := 0; i < len(m.List); i++ {
		if swag.IsZero(m.List[i]) { // not required
			continue
		}

		if m.List[i] != nil {
			if err := m.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetInstList) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstList) validateSummaryByAddressType(formats strfmt.Registry) error {
	if swag.IsZero(m.SummaryByAddressType) { // not required
		return nil
	}

	if m.SummaryByAddressType != nil {
		if err := m.SummaryByAddressType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByAddressType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByAddressType")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstList) validateSummaryByKind(formats strfmt.Registry) error {
	if swag.IsZero(m.SummaryByKind) { // not required
		return nil
	}

	if m.SummaryByKind != nil {
		if err := m.SummaryByKind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByKind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByKind")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this net inst list based on the context it is used
func (m *NetInstList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCfgList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByAddressType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaryByKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetInstList) contextValidateCfgList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CfgList); i++ {

		if m.CfgList[i] != nil {
			if err := m.CfgList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cfgList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cfgList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetInstList) contextValidateList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.List); i++ {

		if m.List[i] != nil {
			if err := m.List[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetInstList) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstList) contextValidateSummaryByAddressType(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByAddressType != nil {
		if err := m.SummaryByAddressType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByAddressType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByAddressType")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstList) contextValidateSummaryByKind(ctx context.Context, formats strfmt.Registry) error {

	if m.SummaryByKind != nil {
		if err := m.SummaryByKind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summaryByKind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("summaryByKind")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetInstList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetInstList) UnmarshalBinary(b []byte) error {
	var res NetInstList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
