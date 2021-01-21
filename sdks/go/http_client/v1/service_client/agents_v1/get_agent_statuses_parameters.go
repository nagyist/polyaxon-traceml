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

package agents_v1

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
)

// NewGetAgentStatusesParams creates a new GetAgentStatusesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAgentStatusesParams() *GetAgentStatusesParams {
	return &GetAgentStatusesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentStatusesParamsWithTimeout creates a new GetAgentStatusesParams object
// with the ability to set a timeout on a request.
func NewGetAgentStatusesParamsWithTimeout(timeout time.Duration) *GetAgentStatusesParams {
	return &GetAgentStatusesParams{
		timeout: timeout,
	}
}

// NewGetAgentStatusesParamsWithContext creates a new GetAgentStatusesParams object
// with the ability to set a context for a request.
func NewGetAgentStatusesParamsWithContext(ctx context.Context) *GetAgentStatusesParams {
	return &GetAgentStatusesParams{
		Context: ctx,
	}
}

// NewGetAgentStatusesParamsWithHTTPClient creates a new GetAgentStatusesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAgentStatusesParamsWithHTTPClient(client *http.Client) *GetAgentStatusesParams {
	return &GetAgentStatusesParams{
		HTTPClient: client,
	}
}

/* GetAgentStatusesParams contains all the parameters to send to the API endpoint
   for the get agent statuses operation.

   Typically these are written to a http.Request.
*/
type GetAgentStatusesParams struct {

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* UUID.

	   Uuid identifier of the entity
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get agent statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentStatusesParams) WithDefaults() *GetAgentStatusesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get agent statuses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAgentStatusesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get agent statuses params
func (o *GetAgentStatusesParams) WithTimeout(timeout time.Duration) *GetAgentStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent statuses params
func (o *GetAgentStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent statuses params
func (o *GetAgentStatusesParams) WithContext(ctx context.Context) *GetAgentStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent statuses params
func (o *GetAgentStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent statuses params
func (o *GetAgentStatusesParams) WithHTTPClient(client *http.Client) *GetAgentStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent statuses params
func (o *GetAgentStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get agent statuses params
func (o *GetAgentStatusesParams) WithOwner(owner string) *GetAgentStatusesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get agent statuses params
func (o *GetAgentStatusesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get agent statuses params
func (o *GetAgentStatusesParams) WithUUID(uuid string) *GetAgentStatusesParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get agent statuses params
func (o *GetAgentStatusesParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
