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
	"github.com/go-openapi/swag"
)

// NewListProjectSearchesParams creates a new ListProjectSearchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectSearchesParams() *ListProjectSearchesParams {
	return &ListProjectSearchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectSearchesParamsWithTimeout creates a new ListProjectSearchesParams object
// with the ability to set a timeout on a request.
func NewListProjectSearchesParamsWithTimeout(timeout time.Duration) *ListProjectSearchesParams {
	return &ListProjectSearchesParams{
		timeout: timeout,
	}
}

// NewListProjectSearchesParamsWithContext creates a new ListProjectSearchesParams object
// with the ability to set a context for a request.
func NewListProjectSearchesParamsWithContext(ctx context.Context) *ListProjectSearchesParams {
	return &ListProjectSearchesParams{
		Context: ctx,
	}
}

// NewListProjectSearchesParamsWithHTTPClient creates a new ListProjectSearchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectSearchesParamsWithHTTPClient(client *http.Client) *ListProjectSearchesParams {
	return &ListProjectSearchesParams{
		HTTPClient: client,
	}
}

/* ListProjectSearchesParams contains all the parameters to send to the API endpoint
   for the list project searches operation.

   Typically these are written to a http.Request.
*/
type ListProjectSearchesParams struct {

	/* Limit.

	   Limit size.

	   Format: int32
	*/
	Limit *int32

	/* Name.

	   Entity managing the resource
	*/
	Name string

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

// WithDefaults hydrates default values in the list project searches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectSearchesParams) WithDefaults() *ListProjectSearchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project searches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectSearchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project searches params
func (o *ListProjectSearchesParams) WithTimeout(timeout time.Duration) *ListProjectSearchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project searches params
func (o *ListProjectSearchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project searches params
func (o *ListProjectSearchesParams) WithContext(ctx context.Context) *ListProjectSearchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project searches params
func (o *ListProjectSearchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project searches params
func (o *ListProjectSearchesParams) WithHTTPClient(client *http.Client) *ListProjectSearchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project searches params
func (o *ListProjectSearchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list project searches params
func (o *ListProjectSearchesParams) WithLimit(limit *int32) *ListProjectSearchesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list project searches params
func (o *ListProjectSearchesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithName adds the name to the list project searches params
func (o *ListProjectSearchesParams) WithName(name string) *ListProjectSearchesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list project searches params
func (o *ListProjectSearchesParams) SetName(name string) {
	o.Name = name
}

// WithOffset adds the offset to the list project searches params
func (o *ListProjectSearchesParams) WithOffset(offset *int32) *ListProjectSearchesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list project searches params
func (o *ListProjectSearchesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOwner adds the owner to the list project searches params
func (o *ListProjectSearchesParams) WithOwner(owner string) *ListProjectSearchesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the list project searches params
func (o *ListProjectSearchesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithQuery adds the query to the list project searches params
func (o *ListProjectSearchesParams) WithQuery(query *string) *ListProjectSearchesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list project searches params
func (o *ListProjectSearchesParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the list project searches params
func (o *ListProjectSearchesParams) WithSort(sort *string) *ListProjectSearchesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list project searches params
func (o *ListProjectSearchesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectSearchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
