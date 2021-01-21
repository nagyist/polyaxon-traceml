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

// NewPatchModelVersionParams creates a new PatchModelVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchModelVersionParams() *PatchModelVersionParams {
	return &PatchModelVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchModelVersionParamsWithTimeout creates a new PatchModelVersionParams object
// with the ability to set a timeout on a request.
func NewPatchModelVersionParamsWithTimeout(timeout time.Duration) *PatchModelVersionParams {
	return &PatchModelVersionParams{
		timeout: timeout,
	}
}

// NewPatchModelVersionParamsWithContext creates a new PatchModelVersionParams object
// with the ability to set a context for a request.
func NewPatchModelVersionParamsWithContext(ctx context.Context) *PatchModelVersionParams {
	return &PatchModelVersionParams{
		Context: ctx,
	}
}

// NewPatchModelVersionParamsWithHTTPClient creates a new PatchModelVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchModelVersionParamsWithHTTPClient(client *http.Client) *PatchModelVersionParams {
	return &PatchModelVersionParams{
		HTTPClient: client,
	}
}

/* PatchModelVersionParams contains all the parameters to send to the API endpoint
   for the patch model version operation.

   Typically these are written to a http.Request.
*/
type PatchModelVersionParams struct {

	/* Body.

	   Model version body
	*/
	Body *service_model.V1ModelVersion

	/* Model.

	   Model name
	*/
	Model string

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* VersionName.

	   Optional component name, should be a valid fully qualified value: name[:version]
	*/
	VersionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch model version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchModelVersionParams) WithDefaults() *PatchModelVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch model version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchModelVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch model version params
func (o *PatchModelVersionParams) WithTimeout(timeout time.Duration) *PatchModelVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch model version params
func (o *PatchModelVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch model version params
func (o *PatchModelVersionParams) WithContext(ctx context.Context) *PatchModelVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch model version params
func (o *PatchModelVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch model version params
func (o *PatchModelVersionParams) WithHTTPClient(client *http.Client) *PatchModelVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch model version params
func (o *PatchModelVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch model version params
func (o *PatchModelVersionParams) WithBody(body *service_model.V1ModelVersion) *PatchModelVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch model version params
func (o *PatchModelVersionParams) SetBody(body *service_model.V1ModelVersion) {
	o.Body = body
}

// WithModel adds the model to the patch model version params
func (o *PatchModelVersionParams) WithModel(model string) *PatchModelVersionParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the patch model version params
func (o *PatchModelVersionParams) SetModel(model string) {
	o.Model = model
}

// WithOwner adds the owner to the patch model version params
func (o *PatchModelVersionParams) WithOwner(owner string) *PatchModelVersionParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch model version params
func (o *PatchModelVersionParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithVersionName adds the versionName to the patch model version params
func (o *PatchModelVersionParams) WithVersionName(versionName string) *PatchModelVersionParams {
	o.SetVersionName(versionName)
	return o
}

// SetVersionName adds the versionName to the patch model version params
func (o *PatchModelVersionParams) SetVersionName(versionName string) {
	o.VersionName = versionName
}

// WriteToRequest writes these params to a swagger request
func (o *PatchModelVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param version.name
	if err := r.SetPathParam("version.name", o.VersionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
