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

package project_searches_v1

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

// NewUpdateProjectSearchParams creates a new UpdateProjectSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProjectSearchParams() *UpdateProjectSearchParams {
	return &UpdateProjectSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectSearchParamsWithTimeout creates a new UpdateProjectSearchParams object
// with the ability to set a timeout on a request.
func NewUpdateProjectSearchParamsWithTimeout(timeout time.Duration) *UpdateProjectSearchParams {
	return &UpdateProjectSearchParams{
		timeout: timeout,
	}
}

// NewUpdateProjectSearchParamsWithContext creates a new UpdateProjectSearchParams object
// with the ability to set a context for a request.
func NewUpdateProjectSearchParamsWithContext(ctx context.Context) *UpdateProjectSearchParams {
	return &UpdateProjectSearchParams{
		Context: ctx,
	}
}

// NewUpdateProjectSearchParamsWithHTTPClient creates a new UpdateProjectSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProjectSearchParamsWithHTTPClient(client *http.Client) *UpdateProjectSearchParams {
	return &UpdateProjectSearchParams{
		HTTPClient: client,
	}
}

/* UpdateProjectSearchParams contains all the parameters to send to the API endpoint
   for the update project search operation.

   Typically these are written to a http.Request.
*/
type UpdateProjectSearchParams struct {

	/* Body.

	   Search body
	*/
	Body *service_model.V1Search

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Project.

	   Project under namesapce
	*/
	Project string

	/* SearchUUID.

	   UUID
	*/
	SearchUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update project search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectSearchParams) WithDefaults() *UpdateProjectSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update project search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update project search params
func (o *UpdateProjectSearchParams) WithTimeout(timeout time.Duration) *UpdateProjectSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project search params
func (o *UpdateProjectSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project search params
func (o *UpdateProjectSearchParams) WithContext(ctx context.Context) *UpdateProjectSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project search params
func (o *UpdateProjectSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project search params
func (o *UpdateProjectSearchParams) WithHTTPClient(client *http.Client) *UpdateProjectSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project search params
func (o *UpdateProjectSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update project search params
func (o *UpdateProjectSearchParams) WithBody(body *service_model.V1Search) *UpdateProjectSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update project search params
func (o *UpdateProjectSearchParams) SetBody(body *service_model.V1Search) {
	o.Body = body
}

// WithOwner adds the owner to the update project search params
func (o *UpdateProjectSearchParams) WithOwner(owner string) *UpdateProjectSearchParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update project search params
func (o *UpdateProjectSearchParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the update project search params
func (o *UpdateProjectSearchParams) WithProject(project string) *UpdateProjectSearchParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update project search params
func (o *UpdateProjectSearchParams) SetProject(project string) {
	o.Project = project
}

// WithSearchUUID adds the searchUUID to the update project search params
func (o *UpdateProjectSearchParams) WithSearchUUID(searchUUID string) *UpdateProjectSearchParams {
	o.SetSearchUUID(searchUUID)
	return o
}

// SetSearchUUID adds the searchUuid to the update project search params
func (o *UpdateProjectSearchParams) SetSearchUUID(searchUUID string) {
	o.SearchUUID = searchUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param search.uuid
	if err := r.SetPathParam("search.uuid", o.SearchUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
