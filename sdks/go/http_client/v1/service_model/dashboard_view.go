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

// DashboardView Dashboard view
//
// - any: Any view
//  - single: Single run dashboard
//  - compare: Compare runs dashboard
//
// swagger:model DashboardView
type DashboardView string

const (

	// DashboardViewAny captures enum value "any"
	DashboardViewAny DashboardView = "any"

	// DashboardViewSingle captures enum value "single"
	DashboardViewSingle DashboardView = "single"

	// DashboardViewCompare captures enum value "compare"
	DashboardViewCompare DashboardView = "compare"
)

// for schema
var dashboardViewEnum []interface{}

func init() {
	var res []DashboardView
	if err := json.Unmarshal([]byte(`["any","single","compare"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dashboardViewEnum = append(dashboardViewEnum, v)
	}
}

func (m DashboardView) validateDashboardViewEnum(path, location string, value DashboardView) error {
	if err := validate.EnumCase(path, location, value, dashboardViewEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this dashboard view
func (m DashboardView) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDashboardViewEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this dashboard view based on context it is used
func (m DashboardView) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
