// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Reference Reference specification
//
// swagger:model v1Reference
type V1Reference struct {

	// dag reference
	DagReference *V1DagRef `json:"dag_reference,omitempty"`

	// hub reference
	HubReference *V1HubRef `json:"hub_reference,omitempty"`

	// path reference
	PathReference *V1PathRef `json:"path_reference,omitempty"`

	// url reference
	URLReference *V1URLRef `json:"url_reference,omitempty"`
}

// Validate validates this v1 reference
func (m *V1Reference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDagReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHubReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePathReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURLReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Reference) validateDagReference(formats strfmt.Registry) error {
	if swag.IsZero(m.DagReference) { // not required
		return nil
	}

	if m.DagReference != nil {
		if err := m.DagReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dag_reference")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) validateHubReference(formats strfmt.Registry) error {
	if swag.IsZero(m.HubReference) { // not required
		return nil
	}

	if m.HubReference != nil {
		if err := m.HubReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hub_reference")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) validatePathReference(formats strfmt.Registry) error {
	if swag.IsZero(m.PathReference) { // not required
		return nil
	}

	if m.PathReference != nil {
		if err := m.PathReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("path_reference")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) validateURLReference(formats strfmt.Registry) error {
	if swag.IsZero(m.URLReference) { // not required
		return nil
	}

	if m.URLReference != nil {
		if err := m.URLReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("url_reference")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 reference based on the context it is used
func (m *V1Reference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDagReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHubReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePathReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURLReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Reference) contextValidateDagReference(ctx context.Context, formats strfmt.Registry) error {

	if m.DagReference != nil {
		if err := m.DagReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dag_reference")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) contextValidateHubReference(ctx context.Context, formats strfmt.Registry) error {

	if m.HubReference != nil {
		if err := m.HubReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hub_reference")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) contextValidatePathReference(ctx context.Context, formats strfmt.Registry) error {

	if m.PathReference != nil {
		if err := m.PathReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("path_reference")
			}
			return err
		}
	}

	return nil
}

func (m *V1Reference) contextValidateURLReference(ctx context.Context, formats strfmt.Registry) error {

	if m.URLReference != nil {
		if err := m.URLReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("url_reference")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Reference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Reference) UnmarshalBinary(b []byte) error {
	var res V1Reference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
