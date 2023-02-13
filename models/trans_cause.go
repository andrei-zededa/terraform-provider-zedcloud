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

// TransCause trans cause
//
// swagger:model TransCause
type TransCause string

func NewTransCause(value TransCause) *TransCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TransCause.
func (m TransCause) Pointer() *TransCause {
	return &m
}

const (

	// TransCauseTRANSCAUSEUNSPECIFIED captures enum value "TRANS_CAUSE_UNSPECIFIED"
	TransCauseTRANSCAUSEUNSPECIFIED TransCause = "TRANS_CAUSE_UNSPECIFIED"

	// TransCauseTRANSCAUSEAPPDRIVECHANGED captures enum value "TRANS_CAUSE_APP_DRIVE_CHANGED"
	TransCauseTRANSCAUSEAPPDRIVECHANGED TransCause = "TRANS_CAUSE_APP_DRIVE_CHANGED"

	// TransCauseTRANSCAUSEAPPNETWORKCHANGED captures enum value "TRANS_CAUSE_APP_NETWORK_CHANGED"
	TransCauseTRANSCAUSEAPPNETWORKCHANGED TransCause = "TRANS_CAUSE_APP_NETWORK_CHANGED"

	// TransCauseTRANSCAUSEAPPCUSTOMCONFIGCHANGED captures enum value "TRANS_CAUSE_APP_CUSTOM_CONFIG_CHANGED"
	TransCauseTRANSCAUSEAPPCUSTOMCONFIGCHANGED TransCause = "TRANS_CAUSE_APP_CUSTOM_CONFIG_CHANGED"

	// TransCauseTRANSCAUSEAPPMODULEBUNDLEUPDATE captures enum value "TRANS_CAUSE_APP_MODULE_BUNDLE_UPDATE"
	TransCauseTRANSCAUSEAPPMODULEBUNDLEUPDATE TransCause = "TRANS_CAUSE_APP_MODULE_BUNDLE_UPDATE"

	// TransCauseTRANSCAUSEAPPDESCCHANGED captures enum value "TRANS_CAUSE_APP_DESC_CHANGED"
	TransCauseTRANSCAUSEAPPDESCCHANGED TransCause = "TRANS_CAUSE_APP_DESC_CHANGED"

	// TransCauseTRANSCAUSEAPPFIXEDRESOURCESCHANGED captures enum value "TRANS_CAUSE_APP_FIXED_RESOURCES_CHANGED"
	TransCauseTRANSCAUSEAPPFIXEDRESOURCESCHANGED TransCause = "TRANS_CAUSE_APP_FIXED_RESOURCES_CHANGED"

	// TransCauseTRANSCAUSEAPPVNCMODECHANGED captures enum value "TRANS_CAUSE_APP_VNC_MODE_CHANGED"
	TransCauseTRANSCAUSEAPPVNCMODECHANGED TransCause = "TRANS_CAUSE_APP_VNC_MODE_CHANGED"

	// TransCauseTRANSCAUSEAPPVMMODECHANGED captures enum value "TRANS_CAUSE_APP_VM_MODE_CHANGED"
	TransCauseTRANSCAUSEAPPVMMODECHANGED TransCause = "TRANS_CAUSE_APP_VM_MODE_CHANGED"

	// TransCauseTRANSCAUSEINTERCAEXPIRING captures enum value "TRANS_CAUSE_INTER_CA_EXPIRING"
	TransCauseTRANSCAUSEINTERCAEXPIRING TransCause = "TRANS_CAUSE_INTER_CA_EXPIRING"

	// TransCauseTRANSCAUSEINTERCAEXPIRED captures enum value "TRANS_CAUSE_INTER_CA_EXPIRED"
	TransCauseTRANSCAUSEINTERCAEXPIRED TransCause = "TRANS_CAUSE_INTER_CA_EXPIRED"

	// TransCauseTRANSCAUSEDEVICECAEXPIRING captures enum value "TRANS_CAUSE_DEVICE_CA_EXPIRING"
	TransCauseTRANSCAUSEDEVICECAEXPIRING TransCause = "TRANS_CAUSE_DEVICE_CA_EXPIRING"

	// TransCauseTRANSCAUSEDEVICECAEXPIRED captures enum value "TRANS_CAUSE_DEVICE_CA_EXPIRED"
	TransCauseTRANSCAUSEDEVICECAEXPIRED TransCause = "TRANS_CAUSE_DEVICE_CA_EXPIRED"
)

// for schema
var transCauseEnum []interface{}

func init() {
	var res []TransCause
	if err := json.Unmarshal([]byte(`["TRANS_CAUSE_UNSPECIFIED","TRANS_CAUSE_APP_DRIVE_CHANGED","TRANS_CAUSE_APP_NETWORK_CHANGED","TRANS_CAUSE_APP_CUSTOM_CONFIG_CHANGED","TRANS_CAUSE_APP_MODULE_BUNDLE_UPDATE","TRANS_CAUSE_APP_DESC_CHANGED","TRANS_CAUSE_APP_FIXED_RESOURCES_CHANGED","TRANS_CAUSE_APP_VNC_MODE_CHANGED","TRANS_CAUSE_APP_VM_MODE_CHANGED","TRANS_CAUSE_INTER_CA_EXPIRING","TRANS_CAUSE_INTER_CA_EXPIRED","TRANS_CAUSE_DEVICE_CA_EXPIRING","TRANS_CAUSE_DEVICE_CA_EXPIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transCauseEnum = append(transCauseEnum, v)
	}
}

func (m TransCause) validateTransCauseEnum(path, location string, value TransCause) error {
	if err := validate.EnumCase(path, location, value, transCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this trans cause
func (m TransCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTransCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this trans cause based on context it is used
func (m TransCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
