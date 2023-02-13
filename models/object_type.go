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

// ObjectType object type
//
// swagger:model ObjectType
type ObjectType string

func NewObjectType(value ObjectType) *ObjectType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ObjectType.
func (m ObjectType) Pointer() *ObjectType {
	return &m
}

const (

	// ObjectTypeOBJECTTYPEUNSPECIFIED captures enum value "OBJECT_TYPE_UNSPECIFIED"
	ObjectTypeOBJECTTYPEUNSPECIFIED ObjectType = "OBJECT_TYPE_UNSPECIFIED"

	// ObjectTypeOBJECTTYPEUSER captures enum value "OBJECT_TYPE_USER"
	ObjectTypeOBJECTTYPEUSER ObjectType = "OBJECT_TYPE_USER"

	// ObjectTypeOBJECTTYPEEDGENODE captures enum value "OBJECT_TYPE_EDGE_NODE"
	ObjectTypeOBJECTTYPEEDGENODE ObjectType = "OBJECT_TYPE_EDGE_NODE"

	// ObjectTypeOBJECTTYPEEDGEAPP captures enum value "OBJECT_TYPE_EDGE_APP"
	ObjectTypeOBJECTTYPEEDGEAPP ObjectType = "OBJECT_TYPE_EDGE_APP"

	// ObjectTypeOBJECTTYPEIMAGE captures enum value "OBJECT_TYPE_IMAGE"
	ObjectTypeOBJECTTYPEIMAGE ObjectType = "OBJECT_TYPE_IMAGE"

	// ObjectTypeOBJECTTYPEEDGEAPPINSTANCE captures enum value "OBJECT_TYPE_EDGE_APP_INSTANCE"
	ObjectTypeOBJECTTYPEEDGEAPPINSTANCE ObjectType = "OBJECT_TYPE_EDGE_APP_INSTANCE"

	// ObjectTypeOBJECTTYPEPROJECT captures enum value "OBJECT_TYPE_PROJECT"
	ObjectTypeOBJECTTYPEPROJECT ObjectType = "OBJECT_TYPE_PROJECT"

	// ObjectTypeOBJECTTYPENETWORK captures enum value "OBJECT_TYPE_NETWORK"
	ObjectTypeOBJECTTYPENETWORK ObjectType = "OBJECT_TYPE_NETWORK"

	// ObjectTypeOBJECTTYPEDATASTORE captures enum value "OBJECT_TYPE_DATASTORE"
	ObjectTypeOBJECTTYPEDATASTORE ObjectType = "OBJECT_TYPE_DATASTORE"

	// ObjectTypeOBJECTTYPESERVICE captures enum value "OBJECT_TYPE_SERVICE"
	ObjectTypeOBJECTTYPESERVICE ObjectType = "OBJECT_TYPE_SERVICE"

	// ObjectTypeOBJECTTYPESERVICEINSTANCE captures enum value "OBJECT_TYPE_SERVICE_INSTANCE"
	ObjectTypeOBJECTTYPESERVICEINSTANCE ObjectType = "OBJECT_TYPE_SERVICE_INSTANCE"

	// ObjectTypeOBJECTTYPEENTERPRISE captures enum value "OBJECT_TYPE_ENTERPRISE"
	ObjectTypeOBJECTTYPEENTERPRISE ObjectType = "OBJECT_TYPE_ENTERPRISE"

	// ObjectTypeOBJECTTYPEROLE captures enum value "OBJECT_TYPE_ROLE"
	ObjectTypeOBJECTTYPEROLE ObjectType = "OBJECT_TYPE_ROLE"

	// ObjectTypeOBJECTTYPECREDENTIAL captures enum value "OBJECT_TYPE_CREDENTIAL"
	ObjectTypeOBJECTTYPECREDENTIAL ObjectType = "OBJECT_TYPE_CREDENTIAL"

	// ObjectTypeOBJECTTYPENETWORKINSTANCE captures enum value "OBJECT_TYPE_NETWORK_INSTANCE"
	ObjectTypeOBJECTTYPENETWORKINSTANCE ObjectType = "OBJECT_TYPE_NETWORK_INSTANCE"

	// ObjectTypeOBJECTTYPEVOLUMEINSTANCE captures enum value "OBJECT_TYPE_VOLUME_INSTANCE"
	ObjectTypeOBJECTTYPEVOLUMEINSTANCE ObjectType = "OBJECT_TYPE_VOLUME_INSTANCE"

	// ObjectTypeOBJECTTYPEREALM captures enum value "OBJECT_TYPE_REALM"
	ObjectTypeOBJECTTYPEREALM ObjectType = "OBJECT_TYPE_REALM"

	// ObjectTypeOBJECTTYPEAUTHPROFILE captures enum value "OBJECT_TYPE_AUTHPROFILE"
	ObjectTypeOBJECTTYPEAUTHPROFILE ObjectType = "OBJECT_TYPE_AUTHPROFILE"

	// ObjectTypeOBJECTTYPEPOLICY captures enum value "OBJECT_TYPE_POLICY"
	ObjectTypeOBJECTTYPEPOLICY ObjectType = "OBJECT_TYPE_POLICY"

	// ObjectTypeOBJECTTYPEAPPPOLICY captures enum value "OBJECT_TYPE_APP_POLICY"
	ObjectTypeOBJECTTYPEAPPPOLICY ObjectType = "OBJECT_TYPE_APP_POLICY"

	// ObjectTypeOBJECTTYPECLUSTERINSTANCE captures enum value "OBJECT_TYPE_CLUSTER_INSTANCE"
	ObjectTypeOBJECTTYPECLUSTERINSTANCE ObjectType = "OBJECT_TYPE_CLUSTER_INSTANCE"

	// ObjectTypeOBJECTTYPEPLUGIN captures enum value "OBJECT_TYPE_PLUGIN"
	ObjectTypeOBJECTTYPEPLUGIN ObjectType = "OBJECT_TYPE_PLUGIN"

	// ObjectTypeOBJECTTYPEDOCPOLICY captures enum value "OBJECT_TYPE_DOC_POLICY"
	ObjectTypeOBJECTTYPEDOCPOLICY ObjectType = "OBJECT_TYPE_DOC_POLICY"

	// ObjectTypeOBJECTTYPEORCHESTRATORCLUSTER captures enum value "OBJECT_TYPE_ORCHESTRATOR_CLUSTER"
	ObjectTypeOBJECTTYPEORCHESTRATORCLUSTER ObjectType = "OBJECT_TYPE_ORCHESTRATOR_CLUSTER"

	// ObjectTypeOBJECTTYPETAGS captures enum value "OBJECT_TYPE_TAGS"
	ObjectTypeOBJECTTYPETAGS ObjectType = "OBJECT_TYPE_TAGS"

	// ObjectTypeOBJECTTYPEEDGENODEINTERFACE captures enum value "OBJECT_TYPE_EDGE_NODE_INTERFACE"
	ObjectTypeOBJECTTYPEEDGENODEINTERFACE ObjectType = "OBJECT_TYPE_EDGE_NODE_INTERFACE"

	// ObjectTypeOBJECTTYPEDEPLOYMENT captures enum value "OBJECT_TYPE_DEPLOYMENT"
	ObjectTypeOBJECTTYPEDEPLOYMENT ObjectType = "OBJECT_TYPE_DEPLOYMENT"

	// ObjectTypeOBJECTTYPEENTITLEMENTS captures enum value "OBJECT_TYPE_ENTITLEMENTS"
	ObjectTypeOBJECTTYPEENTITLEMENTS ObjectType = "OBJECT_TYPE_ENTITLEMENTS"

	// ObjectTypeOBJECTTYPEDATASTREAM captures enum value "OBJECT_TYPE_DATA_STREAM"
	ObjectTypeOBJECTTYPEDATASTREAM ObjectType = "OBJECT_TYPE_DATA_STREAM"

	// ObjectTypeOBJECTTYPEAPIUSAGE captures enum value "OBJECT_TYPE_API_USAGE"
	ObjectTypeOBJECTTYPEAPIUSAGE ObjectType = "OBJECT_TYPE_API_USAGE"
)

// for schema
var objectTypeEnum []interface{}

func init() {
	var res []ObjectType
	if err := json.Unmarshal([]byte(`["OBJECT_TYPE_UNSPECIFIED","OBJECT_TYPE_USER","OBJECT_TYPE_EDGE_NODE","OBJECT_TYPE_EDGE_APP","OBJECT_TYPE_IMAGE","OBJECT_TYPE_EDGE_APP_INSTANCE","OBJECT_TYPE_PROJECT","OBJECT_TYPE_NETWORK","OBJECT_TYPE_DATASTORE","OBJECT_TYPE_SERVICE","OBJECT_TYPE_SERVICE_INSTANCE","OBJECT_TYPE_ENTERPRISE","OBJECT_TYPE_ROLE","OBJECT_TYPE_CREDENTIAL","OBJECT_TYPE_NETWORK_INSTANCE","OBJECT_TYPE_VOLUME_INSTANCE","OBJECT_TYPE_REALM","OBJECT_TYPE_AUTHPROFILE","OBJECT_TYPE_POLICY","OBJECT_TYPE_APP_POLICY","OBJECT_TYPE_CLUSTER_INSTANCE","OBJECT_TYPE_PLUGIN","OBJECT_TYPE_DOC_POLICY","OBJECT_TYPE_ORCHESTRATOR_CLUSTER","OBJECT_TYPE_TAGS","OBJECT_TYPE_EDGE_NODE_INTERFACE","OBJECT_TYPE_DEPLOYMENT","OBJECT_TYPE_ENTITLEMENTS","OBJECT_TYPE_DATA_STREAM","OBJECT_TYPE_API_USAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectTypeEnum = append(objectTypeEnum, v)
	}
}

func (m ObjectType) validateObjectTypeEnum(path, location string, value ObjectType) error {
	if err := validate.EnumCase(path, location, value, objectTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this object type
func (m ObjectType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateObjectTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this object type based on context it is used
func (m ObjectType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
