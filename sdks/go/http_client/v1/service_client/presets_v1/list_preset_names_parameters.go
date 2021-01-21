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
	"github.com/go-openapi/swag"
)

// NewListPresetNamesParams creates a new ListPresetNamesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPresetNamesParams() *ListPresetNamesParams {
	return &ListPresetNamesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPresetNamesParamsWithTimeout creates a new ListPresetNamesParams object
// with the ability to set a timeout on a request.
func NewListPresetNamesParamsWithTimeout(timeout time.Duration) *ListPresetNamesParams {
	return &ListPresetNamesParams{
		timeout: timeout,
	}
}

// NewListPresetNamesParamsWithContext creates a new ListPresetNamesParams object
// with the ability to set a context for a request.
func NewListPresetNamesParamsWithContext(ctx context.Context) *ListPresetNamesParams {
	return &ListPresetNamesParams{
		Context: ctx,
	}
}

// NewListPresetNamesParamsWithHTTPClient creates a new ListPresetNamesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPresetNamesParamsWithHTTPClient(client *http.Client) *ListPresetNamesParams {
	return &ListPresetNamesParams{
		HTTPClient: client,
	}
}

/* ListPresetNamesParams contains all the parameters to send to the API endpoint
   for the list preset names operation.

   Typically these are written to a http.Request.
*/
type ListPresetNamesParams struct {

	/* Limit.

	   Limit size.

	   Format: int32
	*/
	Limit *int32

	/* Offset.

	   Pagination offset.

	   Format: int32
	*/
	Offset *int32

	/* Owner.

	   Owner of the namespace
	*/
	Owner string

	/* Query.

	   Query filter the search.
	*/
	Query *string

	/* Sort.

	   Sort to order the search.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list preset names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPresetNamesParams) WithDefaults() *ListPresetNamesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list preset names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPresetNamesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list preset names params
func (o *ListPresetNamesParams) WithTimeout(timeout time.Duration) *ListPresetNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list preset names params
func (o *ListPresetNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list preset names params
func (o *ListPresetNamesParams) WithContext(ctx context.Context) *ListPresetNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list preset names params
func (o *ListPresetNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list preset names params
func (o *ListPresetNamesParams) WithHTTPClient(client *http.Client) *ListPresetNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list preset names params
func (o *ListPresetNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list preset names params
func (o *ListPresetNamesParams) WithLimit(limit *int32) *ListPresetNamesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list preset names params
func (o *ListPresetNamesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list preset names params
func (o *ListPresetNamesParams) WithOffset(offset *int32) *ListPresetNamesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list preset names params
func (o *ListPresetNamesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOwner adds the owner to the list preset names params
func (o *ListPresetNamesParams) WithOwner(owner string) *ListPresetNamesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the list preset names params
func (o *ListPresetNamesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithQuery adds the query to the list preset names params
func (o *ListPresetNamesParams) WithQuery(query *string) *ListPresetNamesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list preset names params
func (o *ListPresetNamesParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the list preset names params
func (o *ListPresetNamesParams) WithSort(sort *string) *ListPresetNamesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list preset names params
func (o *ListPresetNamesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListPresetNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
