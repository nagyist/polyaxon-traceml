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

package model_registry_v1

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

// NewUpdateModelRegistrySettingsParams creates a new UpdateModelRegistrySettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateModelRegistrySettingsParams() *UpdateModelRegistrySettingsParams {
	return &UpdateModelRegistrySettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateModelRegistrySettingsParamsWithTimeout creates a new UpdateModelRegistrySettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateModelRegistrySettingsParamsWithTimeout(timeout time.Duration) *UpdateModelRegistrySettingsParams {
	return &UpdateModelRegistrySettingsParams{
		timeout: timeout,
	}
}

// NewUpdateModelRegistrySettingsParamsWithContext creates a new UpdateModelRegistrySettingsParams object
// with the ability to set a context for a request.
func NewUpdateModelRegistrySettingsParamsWithContext(ctx context.Context) *UpdateModelRegistrySettingsParams {
	return &UpdateModelRegistrySettingsParams{
		Context: ctx,
	}
}

// NewUpdateModelRegistrySettingsParamsWithHTTPClient creates a new UpdateModelRegistrySettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateModelRegistrySettingsParamsWithHTTPClient(client *http.Client) *UpdateModelRegistrySettingsParams {
	return &UpdateModelRegistrySettingsParams{
		HTTPClient: client,
	}
}

/* UpdateModelRegistrySettingsParams contains all the parameters to send to the API endpoint
   for the update model registry settings operation.

   Typically these are written to a http.Request.
*/
type UpdateModelRegistrySettingsParams struct {

	/* Body.

	   Model settings body
	*/
	Body *service_model.V1ModelRegistrySettings

	/* Model.

	   Model name
	*/
	Model string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update model registry settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateModelRegistrySettingsParams) WithDefaults() *UpdateModelRegistrySettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update model registry settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateModelRegistrySettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) WithTimeout(timeout time.Duration) *UpdateModelRegistrySettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) WithContext(ctx context.Context) *UpdateModelRegistrySettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) WithHTTPClient(client *http.Client) *UpdateModelRegistrySettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) WithBody(body *service_model.V1ModelRegistrySettings) *UpdateModelRegistrySettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) SetBody(body *service_model.V1ModelRegistrySettings) {
	o.Body = body
}

// WithModel adds the model to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) WithModel(model string) *UpdateModelRegistrySettingsParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) SetModel(model string) {
	o.Model = model
}

// WithOwner adds the owner to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) WithOwner(owner string) *UpdateModelRegistrySettingsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update model registry settings params
func (o *UpdateModelRegistrySettingsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateModelRegistrySettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param model
	if err := r.SetPathParam("model", o.Model); err != nil {
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
