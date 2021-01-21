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
	"github.com/go-openapi/validate"
)

// V1StageCondition stage condition specification
//
// swagger:model v1StageCondition
type V1StageCondition struct {

	// last transition time
	// Format: date-time
	LastTransitionTime strfmt.DateTime `json:"last_transition_time,omitempty"`

	// last update time
	// Format: date-time
	LastUpdateTime strfmt.DateTime `json:"last_update_time,omitempty"`

	// Status message
	Message string `json:"message,omitempty"`

	// Status reason
	Reason string `json:"reason,omitempty"`

	// Status state
	Status string `json:"status,omitempty"`

	// Status type
	Type *V1Stages `json:"type,omitempty"`
}

// Validate validates this v1 stage condition
func (m *V1StageCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1StageCondition) validateLastTransitionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("last_transition_time", "body", "date-time", m.LastTransitionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1StageCondition) validateLastUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("last_update_time", "body", "date-time", m.LastUpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1StageCondition) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 stage condition based on the context it is used
func (m *V1StageCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1StageCondition) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1StageCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1StageCondition) UnmarshalBinary(b []byte) error {
	var res V1StageCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
