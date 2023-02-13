// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ImageType Image types
//
// - IMAGE_TYPE_EVE: Base OS images for edge gateway
//   - IMAGE_TYPE_APPLICATION: Edge application images
//   - IMAGE_TYPE_EVEPRIVATE: Private Base OS images for edge gateway
//
// swagger:model ImageType
type ImageType string

func NewImageType(value ImageType) *ImageType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ImageType.
func (m ImageType) Pointer() *ImageType {
	return &m
}

const (

	// ImageTypeIMAGETYPEUNSPECIFIED captures enum value "IMAGE_TYPE_UNSPECIFIED"
	ImageTypeIMAGETYPEUNSPECIFIED ImageType = "IMAGE_TYPE_UNSPECIFIED"

	// ImageTypeIMAGETYPEEVE captures enum value "IMAGE_TYPE_EVE"
	ImageTypeIMAGETYPEEVE ImageType = "IMAGE_TYPE_EVE"

	// ImageTypeIMAGETYPEAPPLICATION captures enum value "IMAGE_TYPE_APPLICATION"
	ImageTypeIMAGETYPEAPPLICATION ImageType = "IMAGE_TYPE_APPLICATION"

	// ImageTypeIMAGETYPEEVEPRIVATE captures enum value "IMAGE_TYPE_EVEPRIVATE"
	ImageTypeIMAGETYPEEVEPRIVATE ImageType = "IMAGE_TYPE_EVEPRIVATE"
)

// for schema
var imageTypeEnum []interface{}

func init() {
	var res []ImageType
	if err := json.Unmarshal([]byte(`["IMAGE_TYPE_UNSPECIFIED","IMAGE_TYPE_EVE","IMAGE_TYPE_APPLICATION","IMAGE_TYPE_EVEPRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		imageTypeEnum = append(imageTypeEnum, v)
	}
}

func (m ImageType) validateImageTypeEnum(path, location string, value ImageType) error {
	if err := validate.EnumCase(path, location, value, imageTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this image type
func (m ImageType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateImageTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this image type based on context it is used
func (m ImageType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
