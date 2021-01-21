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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1RunEdgeKind v1 run edge kind
//
// swagger:model v1RunEdgeKind
type V1RunEdgeKind string

const (

	// V1RunEdgeKindAction captures enum value "action"
	V1RunEdgeKindAction V1RunEdgeKind = "action"

	// V1RunEdgeKindEvent captures enum value "event"
	V1RunEdgeKindEvent V1RunEdgeKind = "event"

	// V1RunEdgeKindHook captures enum value "hook"
	V1RunEdgeKindHook V1RunEdgeKind = "hook"

	// V1RunEdgeKindDag captures enum value "dag"
	V1RunEdgeKindDag V1RunEdgeKind = "dag"

	// V1RunEdgeKindJoin captures enum value "join"
	V1RunEdgeKindJoin V1RunEdgeKind = "join"

	// V1RunEdgeKindRun captures enum value "run"
	V1RunEdgeKindRun V1RunEdgeKind = "run"
)

// for schema
var v1RunEdgeKindEnum []interface{}

func init() {
	var res []V1RunEdgeKind
	if err := json.Unmarshal([]byte(`["action","event","hook","dag","join","run"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1RunEdgeKindEnum = append(v1RunEdgeKindEnum, v)
	}
}

func (m V1RunEdgeKind) validateV1RunEdgeKindEnum(path, location string, value V1RunEdgeKind) error {
	if err := validate.EnumCase(path, location, value, v1RunEdgeKindEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 run edge kind
func (m V1RunEdgeKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1RunEdgeKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 run edge kind based on context it is used
func (m V1RunEdgeKind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
