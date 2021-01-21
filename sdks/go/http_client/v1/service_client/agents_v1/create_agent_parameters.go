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

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewCreateAgentParams creates a new CreateAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAgentParams() *CreateAgentParams {
	return &CreateAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAgentParamsWithTimeout creates a new CreateAgentParams object
// with the ability to set a timeout on a request.
func NewCreateAgentParamsWithTimeout(timeout time.Duration) *CreateAgentParams {
	return &CreateAgentParams{
		timeout: timeout,
	}
}

// NewCreateAgentParamsWithContext creates a new CreateAgentParams object
// with the ability to set a context for a request.
func NewCreateAgentParamsWithContext(ctx context.Context) *CreateAgentParams {
	return &CreateAgentParams{
		Context: ctx,
	}
}

// NewCreateAgentParamsWithHTTPClient creates a new CreateAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAgentParamsWithHTTPClient(client *http.Client) *CreateAgentParams {
	return &CreateAgentParams{
		HTTPClient: client,
	}
}

/* CreateAgentParams contains all the parameters to send to the API endpoint
   for the create agent operation.

   Typically these are written to a http.Request.
*/
type CreateAgentParams struct {

	/* Body.

	   Agent body
	*/
	Body *service_model.V1Agent

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAgentParams) WithDefaults() *CreateAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create agent params
func (o *CreateAgentParams) WithTimeout(timeout time.Duration) *CreateAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create agent params
func (o *CreateAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create agent params
func (o *CreateAgentParams) WithContext(ctx context.Context) *CreateAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create agent params
func (o *CreateAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create agent params
func (o *CreateAgentParams) WithHTTPClient(client *http.Client) *CreateAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create agent params
func (o *CreateAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create agent params
func (o *CreateAgentParams) WithBody(body *service_model.V1Agent) *CreateAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create agent params
func (o *CreateAgentParams) SetBody(body *service_model.V1Agent) {
	o.Body = body
}

// WithOwner adds the owner to the create agent params
func (o *CreateAgentParams) WithOwner(owner string) *CreateAgentParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create agent params
func (o *CreateAgentParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
