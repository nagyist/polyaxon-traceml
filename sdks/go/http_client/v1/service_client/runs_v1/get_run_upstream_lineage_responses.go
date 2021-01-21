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

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetRunUpstreamLineageReader is a Reader for the GetRunUpstreamLineage structure.
type GetRunUpstreamLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunUpstreamLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunUpstreamLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunUpstreamLineageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunUpstreamLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunUpstreamLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRunUpstreamLineageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunUpstreamLineageOK creates a GetRunUpstreamLineageOK with default headers values
func NewGetRunUpstreamLineageOK() *GetRunUpstreamLineageOK {
	return &GetRunUpstreamLineageOK{}
}

/* GetRunUpstreamLineageOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetRunUpstreamLineageOK struct {
	Payload *service_model.V1ListRunEdgesResponse
}

func (o *GetRunUpstreamLineageOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/upstream][%d] getRunUpstreamLineageOK  %+v", 200, o.Payload)
}
func (o *GetRunUpstreamLineageOK) GetPayload() *service_model.V1ListRunEdgesResponse {
	return o.Payload
}

func (o *GetRunUpstreamLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunEdgesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunUpstreamLineageNoContent creates a GetRunUpstreamLineageNoContent with default headers values
func NewGetRunUpstreamLineageNoContent() *GetRunUpstreamLineageNoContent {
	return &GetRunUpstreamLineageNoContent{}
}

/* GetRunUpstreamLineageNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetRunUpstreamLineageNoContent struct {
	Payload interface{}
}

func (o *GetRunUpstreamLineageNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/upstream][%d] getRunUpstreamLineageNoContent  %+v", 204, o.Payload)
}
func (o *GetRunUpstreamLineageNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunUpstreamLineageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunUpstreamLineageForbidden creates a GetRunUpstreamLineageForbidden with default headers values
func NewGetRunUpstreamLineageForbidden() *GetRunUpstreamLineageForbidden {
	return &GetRunUpstreamLineageForbidden{}
}

/* GetRunUpstreamLineageForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetRunUpstreamLineageForbidden struct {
	Payload interface{}
}

func (o *GetRunUpstreamLineageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/upstream][%d] getRunUpstreamLineageForbidden  %+v", 403, o.Payload)
}
func (o *GetRunUpstreamLineageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunUpstreamLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunUpstreamLineageNotFound creates a GetRunUpstreamLineageNotFound with default headers values
func NewGetRunUpstreamLineageNotFound() *GetRunUpstreamLineageNotFound {
	return &GetRunUpstreamLineageNotFound{}
}

/* GetRunUpstreamLineageNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetRunUpstreamLineageNotFound struct {
	Payload interface{}
}

func (o *GetRunUpstreamLineageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/upstream][%d] getRunUpstreamLineageNotFound  %+v", 404, o.Payload)
}
func (o *GetRunUpstreamLineageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunUpstreamLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunUpstreamLineageDefault creates a GetRunUpstreamLineageDefault with default headers values
func NewGetRunUpstreamLineageDefault(code int) *GetRunUpstreamLineageDefault {
	return &GetRunUpstreamLineageDefault{
		_statusCode: code,
	}
}

/* GetRunUpstreamLineageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetRunUpstreamLineageDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get run upstream lineage default response
func (o *GetRunUpstreamLineageDefault) Code() int {
	return o._statusCode
}

func (o *GetRunUpstreamLineageDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/upstream][%d] GetRunUpstreamLineage default  %+v", o._statusCode, o.Payload)
}
func (o *GetRunUpstreamLineageDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetRunUpstreamLineageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
