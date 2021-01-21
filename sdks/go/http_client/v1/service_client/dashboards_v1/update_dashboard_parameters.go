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

package dashboards_v1

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

// NewUpdateDashboardParams creates a new UpdateDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDashboardParams() *UpdateDashboardParams {
	return &UpdateDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDashboardParamsWithTimeout creates a new UpdateDashboardParams object
// with the ability to set a timeout on a request.
func NewUpdateDashboardParamsWithTimeout(timeout time.Duration) *UpdateDashboardParams {
	return &UpdateDashboardParams{
		timeout: timeout,
	}
}

// NewUpdateDashboardParamsWithContext creates a new UpdateDashboardParams object
// with the ability to set a context for a request.
func NewUpdateDashboardParamsWithContext(ctx context.Context) *UpdateDashboardParams {
	return &UpdateDashboardParams{
		Context: ctx,
	}
}

// NewUpdateDashboardParamsWithHTTPClient creates a new UpdateDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDashboardParamsWithHTTPClient(client *http.Client) *UpdateDashboardParams {
	return &UpdateDashboardParams{
		HTTPClient: client,
	}
}

/* UpdateDashboardParams contains all the parameters to send to the API endpoint
   for the update dashboard operation.

   Typically these are written to a http.Request.
*/
type UpdateDashboardParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardParams) WithDefaults() *UpdateDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update dashboard params
func (o *UpdateDashboardParams) WithTimeout(timeout time.Duration) *UpdateDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update dashboard params
func (o *UpdateDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update dashboard params
func (o *UpdateDashboardParams) WithContext(ctx context.Context) *UpdateDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update dashboard params
func (o *UpdateDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update dashboard params
func (o *UpdateDashboardParams) WithHTTPClient(client *http.Client) *UpdateDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update dashboard params
func (o *UpdateDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update dashboard params
func (o *UpdateDashboardParams) WithBody(body *service_model.V1Dashboard) *UpdateDashboardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update dashboard params
func (o *UpdateDashboardParams) SetBody(body *service_model.V1Dashboard) {
	o.Body = body
}

// WithDashboardUUID adds the dashboardUUID to the update dashboard params
func (o *UpdateDashboardParams) WithDashboardUUID(dashboardUUID string) *UpdateDashboardParams {
	o.SetDashboardUUID(dashboardUUID)
	return o
}

// SetDashboardUUID adds the dashboardUuid to the update dashboard params
func (o *UpdateDashboardParams) SetDashboardUUID(dashboardUUID string) {
	o.DashboardUUID = dashboardUUID
}

// WithOwner adds the owner to the update dashboard params
func (o *UpdateDashboardParams) WithOwner(owner string) *UpdateDashboardParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update dashboard params
func (o *UpdateDashboardParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
