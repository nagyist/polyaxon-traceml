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

package presets_v1

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

// NewPatchPresetParams creates a new PatchPresetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchPresetParams() *PatchPresetParams {
	return &PatchPresetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPresetParamsWithTimeout creates a new PatchPresetParams object
// with the ability to set a timeout on a request.
func NewPatchPresetParamsWithTimeout(timeout time.Duration) *PatchPresetParams {
	return &PatchPresetParams{
		timeout: timeout,
	}
}

// NewPatchPresetParamsWithContext creates a new PatchPresetParams object
// with the ability to set a context for a request.
func NewPatchPresetParamsWithContext(ctx context.Context) *PatchPresetParams {
	return &PatchPresetParams{
		Context: ctx,
	}
}

// NewPatchPresetParamsWithHTTPClient creates a new PatchPresetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchPresetParamsWithHTTPClient(client *http.Client) *PatchPresetParams {
	return &PatchPresetParams{
		HTTPClient: client,
	}
}

/* PatchPresetParams contains all the parameters to send to the API endpoint
   for the patch preset operation.

   Typically these are written to a http.Request.
*/
type PatchPresetParams struct {

	/* Body.

	   Preset body
	*/
	Body *service_model.V1Preset

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* PresetUUID.

	   UUID
	*/
	PresetUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch preset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPresetParams) WithDefaults() *PatchPresetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch preset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchPresetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch preset params
func (o *PatchPresetParams) WithTimeout(timeout time.Duration) *PatchPresetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch preset params
func (o *PatchPresetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch preset params
func (o *PatchPresetParams) WithContext(ctx context.Context) *PatchPresetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch preset params
func (o *PatchPresetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch preset params
func (o *PatchPresetParams) WithHTTPClient(client *http.Client) *PatchPresetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch preset params
func (o *PatchPresetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch preset params
func (o *PatchPresetParams) WithBody(body *service_model.V1Preset) *PatchPresetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch preset params
func (o *PatchPresetParams) SetBody(body *service_model.V1Preset) {
	o.Body = body
}

// WithOwner adds the owner to the patch preset params
func (o *PatchPresetParams) WithOwner(owner string) *PatchPresetParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch preset params
func (o *PatchPresetParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithPresetUUID adds the presetUUID to the patch preset params
func (o *PatchPresetParams) WithPresetUUID(presetUUID string) *PatchPresetParams {
	o.SetPresetUUID(presetUUID)
	return o
}

// SetPresetUUID adds the presetUuid to the patch preset params
func (o *PatchPresetParams) SetPresetUUID(presetUUID string) {
	o.PresetUUID = presetUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPresetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param preset.uuid
	if err := r.SetPathParam("preset.uuid", o.PresetUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
