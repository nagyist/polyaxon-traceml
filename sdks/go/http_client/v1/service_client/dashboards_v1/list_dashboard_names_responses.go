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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListDashboardNamesReader is a Reader for the ListDashboardNames structure.
type ListDashboardNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDashboardNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDashboardNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListDashboardNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListDashboardNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListDashboardNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListDashboardNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDashboardNamesOK creates a ListDashboardNamesOK with default headers values
func NewListDashboardNamesOK() *ListDashboardNamesOK {
	return &ListDashboardNamesOK{}
}

/* ListDashboardNamesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListDashboardNamesOK struct {
	Payload *service_model.V1ListDashboardsResponse
}

func (o *ListDashboardNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/names][%d] listDashboardNamesOK  %+v", 200, o.Payload)
}
func (o *ListDashboardNamesOK) GetPayload() *service_model.V1ListDashboardsResponse {
	return o.Payload
}

func (o *ListDashboardNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListDashboardsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDashboardNamesNoContent creates a ListDashboardNamesNoContent with default headers values
func NewListDashboardNamesNoContent() *ListDashboardNamesNoContent {
	return &ListDashboardNamesNoContent{}
}

/* ListDashboardNamesNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListDashboardNamesNoContent struct {
	Payload interface{}
}

func (o *ListDashboardNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/names][%d] listDashboardNamesNoContent  %+v", 204, o.Payload)
}
func (o *ListDashboardNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListDashboardNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDashboardNamesForbidden creates a ListDashboardNamesForbidden with default headers values
func NewListDashboardNamesForbidden() *ListDashboardNamesForbidden {
	return &ListDashboardNamesForbidden{}
}

/* ListDashboardNamesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListDashboardNamesForbidden struct {
	Payload interface{}
}

func (o *ListDashboardNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/names][%d] listDashboardNamesForbidden  %+v", 403, o.Payload)
}
func (o *ListDashboardNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListDashboardNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDashboardNamesNotFound creates a ListDashboardNamesNotFound with default headers values
func NewListDashboardNamesNotFound() *ListDashboardNamesNotFound {
	return &ListDashboardNamesNotFound{}
}

/* ListDashboardNamesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListDashboardNamesNotFound struct {
	Payload interface{}
}

func (o *ListDashboardNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/names][%d] listDashboardNamesNotFound  %+v", 404, o.Payload)
}
func (o *ListDashboardNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListDashboardNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDashboardNamesDefault creates a ListDashboardNamesDefault with default headers values
func NewListDashboardNamesDefault(code int) *ListDashboardNamesDefault {
	return &ListDashboardNamesDefault{
		_statusCode: code,
	}
}

/* ListDashboardNamesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListDashboardNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list dashboard names default response
func (o *ListDashboardNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListDashboardNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/dashboards/names][%d] ListDashboardNames default  %+v", o._statusCode, o.Payload)
}
func (o *ListDashboardNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListDashboardNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
