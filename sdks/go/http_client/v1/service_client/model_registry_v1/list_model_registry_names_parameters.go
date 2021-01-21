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
	"github.com/go-openapi/swag"
)

// NewListModelRegistryNamesParams creates a new ListModelRegistryNamesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListModelRegistryNamesParams() *ListModelRegistryNamesParams {
	return &ListModelRegistryNamesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListModelRegistryNamesParamsWithTimeout creates a new ListModelRegistryNamesParams object
// with the ability to set a timeout on a request.
func NewListModelRegistryNamesParamsWithTimeout(timeout time.Duration) *ListModelRegistryNamesParams {
	return &ListModelRegistryNamesParams{
		timeout: timeout,
	}
}

// NewListModelRegistryNamesParamsWithContext creates a new ListModelRegistryNamesParams object
// with the ability to set a context for a request.
func NewListModelRegistryNamesParamsWithContext(ctx context.Context) *ListModelRegistryNamesParams {
	return &ListModelRegistryNamesParams{
		Context: ctx,
	}
}

// NewListModelRegistryNamesParamsWithHTTPClient creates a new ListModelRegistryNamesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListModelRegistryNamesParamsWithHTTPClient(client *http.Client) *ListModelRegistryNamesParams {
	return &ListModelRegistryNamesParams{
		HTTPClient: client,
	}
}

/* ListModelRegistryNamesParams contains all the parameters to send to the API endpoint
   for the list model registry names operation.

   Typically these are written to a http.Request.
*/
type ListModelRegistryNamesParams struct {

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

// WithDefaults hydrates default values in the list model registry names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListModelRegistryNamesParams) WithDefaults() *ListModelRegistryNamesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list model registry names params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListModelRegistryNamesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list model registry names params
func (o *ListModelRegistryNamesParams) WithTimeout(timeout time.Duration) *ListModelRegistryNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list model registry names params
func (o *ListModelRegistryNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list model registry names params
func (o *ListModelRegistryNamesParams) WithContext(ctx context.Context) *ListModelRegistryNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list model registry names params
func (o *ListModelRegistryNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list model registry names params
func (o *ListModelRegistryNamesParams) WithHTTPClient(client *http.Client) *ListModelRegistryNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list model registry names params
func (o *ListModelRegistryNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list model registry names params
func (o *ListModelRegistryNamesParams) WithLimit(limit *int32) *ListModelRegistryNamesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list model registry names params
func (o *ListModelRegistryNamesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the list model registry names params
func (o *ListModelRegistryNamesParams) WithOffset(offset *int32) *ListModelRegistryNamesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list model registry names params
func (o *ListModelRegistryNamesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOwner adds the owner to the list model registry names params
func (o *ListModelRegistryNamesParams) WithOwner(owner string) *ListModelRegistryNamesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the list model registry names params
func (o *ListModelRegistryNamesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithQuery adds the query to the list model registry names params
func (o *ListModelRegistryNamesParams) WithQuery(query *string) *ListModelRegistryNamesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list model registry names params
func (o *ListModelRegistryNamesParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the list model registry names params
func (o *ListModelRegistryNamesParams) WithSort(sort *string) *ListModelRegistryNamesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list model registry names params
func (o *ListModelRegistryNamesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListModelRegistryNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
