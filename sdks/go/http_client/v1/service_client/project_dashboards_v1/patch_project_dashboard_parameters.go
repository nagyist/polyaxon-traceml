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

package project_dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewPatchProjectDashboardParams creates a new PatchProjectDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchProjectDashboardParams() *PatchProjectDashboardParams {
	return &PatchProjectDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchProjectDashboardParamsWithTimeout creates a new PatchProjectDashboardParams object
// with the ability to set a timeout on a request.
func NewPatchProjectDashboardParamsWithTimeout(timeout time.Duration) *PatchProjectDashboardParams {
	return &PatchProjectDashboardParams{
		timeout: timeout,
	}
}

// NewPatchProjectDashboardParamsWithContext creates a new PatchProjectDashboardParams object
// with the ability to set a context for a request.
func NewPatchProjectDashboardParamsWithContext(ctx context.Context) *PatchProjectDashboardParams {
	return &PatchProjectDashboardParams{
		Context: ctx,
	}
}

// NewPatchProjectDashboardParamsWithHTTPClient creates a new PatchProjectDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchProjectDashboardParamsWithHTTPClient(client *http.Client) *PatchProjectDashboardParams {
	return &PatchProjectDashboardParams{
		HTTPClient: client,
	}
}

/* PatchProjectDashboardParams contains all the parameters to send to the API endpoint
   for the patch project dashboard operation.

   Typically these are written to a http.Request.
*/
type PatchProjectDashboardParams struct {

	/* Body.

	   Dashboard body
	*/
	Body *service_model.V1Dashboard

	/* DashboardUUID.

	   UUID
	*/
	DashboardUUID string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project under namesapce
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch project dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchProjectDashboardParams) WithDefaults() *PatchProjectDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch project dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchProjectDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch project dashboard params
func (o *PatchProjectDashboardParams) WithTimeout(timeout time.Duration) *PatchProjectDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch project dashboard params
func (o *PatchProjectDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch project dashboard params
func (o *PatchProjectDashboardParams) WithContext(ctx context.Context) *PatchProjectDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch project dashboard params
func (o *PatchProjectDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch project dashboard params
func (o *PatchProjectDashboardParams) WithHTTPClient(client *http.Client) *PatchProjectDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch project dashboard params
func (o *PatchProjectDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch project dashboard params
func (o *PatchProjectDashboardParams) WithBody(body *service_model.V1Dashboard) *PatchProjectDashboardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch project dashboard params
func (o *PatchProjectDashboardParams) SetBody(body *service_model.V1Dashboard) {
	o.Body = body
}

// WithDashboardUUID adds the dashboardUUID to the patch project dashboard params
func (o *PatchProjectDashboardParams) WithDashboardUUID(dashboardUUID string) *PatchProjectDashboardParams {
	o.SetDashboardUUID(dashboardUUID)
	return o
}

// SetDashboardUUID adds the dashboardUuid to the patch project dashboard params
func (o *PatchProjectDashboardParams) SetDashboardUUID(dashboardUUID string) {
	o.DashboardUUID = dashboardUUID
}

// WithOwner adds the owner to the patch project dashboard params
func (o *PatchProjectDashboardParams) WithOwner(owner string) *PatchProjectDashboardParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch project dashboard params
func (o *PatchProjectDashboardParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the patch project dashboard params
func (o *PatchProjectDashboardParams) WithProject(project string) *PatchProjectDashboardParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch project dashboard params
func (o *PatchProjectDashboardParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PatchProjectDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param dashboard.uuid
	if err := r.SetPathParam("dashboard.uuid", o.DashboardUUID); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
