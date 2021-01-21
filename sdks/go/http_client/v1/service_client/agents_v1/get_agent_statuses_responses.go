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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetAgentStatusesReader is a Reader for the GetAgentStatuses structure.
type GetAgentStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAgentStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAgentStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetAgentStatusesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAgentStatusesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAgentStatusesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAgentStatusesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAgentStatusesOK creates a GetAgentStatusesOK with default headers values
func NewGetAgentStatusesOK() *GetAgentStatusesOK {
	return &GetAgentStatusesOK{}
}

/* GetAgentStatusesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAgentStatusesOK struct {
	Payload *service_model.V1Status
}

func (o *GetAgentStatusesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] getAgentStatusesOK  %+v", 200, o.Payload)
}
func (o *GetAgentStatusesOK) GetPayload() *service_model.V1Status {
	return o.Payload
}

func (o *GetAgentStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgentStatusesNoContent creates a GetAgentStatusesNoContent with default headers values
func NewGetAgentStatusesNoContent() *GetAgentStatusesNoContent {
	return &GetAgentStatusesNoContent{}
}

/* GetAgentStatusesNoContent describes a response with status code 204, with default header values.

No content.
*/
type GetAgentStatusesNoContent struct {
	Payload interface{}
}

func (o *GetAgentStatusesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] getAgentStatusesNoContent  %+v", 204, o.Payload)
}
func (o *GetAgentStatusesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAgentStatusesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgentStatusesForbidden creates a GetAgentStatusesForbidden with default headers values
func NewGetAgentStatusesForbidden() *GetAgentStatusesForbidden {
	return &GetAgentStatusesForbidden{}
}

/* GetAgentStatusesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type GetAgentStatusesForbidden struct {
	Payload interface{}
}

func (o *GetAgentStatusesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] getAgentStatusesForbidden  %+v", 403, o.Payload)
}
func (o *GetAgentStatusesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAgentStatusesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgentStatusesNotFound creates a GetAgentStatusesNotFound with default headers values
func NewGetAgentStatusesNotFound() *GetAgentStatusesNotFound {
	return &GetAgentStatusesNotFound{}
}

/* GetAgentStatusesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type GetAgentStatusesNotFound struct {
	Payload interface{}
}

func (o *GetAgentStatusesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] getAgentStatusesNotFound  %+v", 404, o.Payload)
}
func (o *GetAgentStatusesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAgentStatusesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAgentStatusesDefault creates a GetAgentStatusesDefault with default headers values
func NewGetAgentStatusesDefault(code int) *GetAgentStatusesDefault {
	return &GetAgentStatusesDefault{
		_statusCode: code,
	}
}

/* GetAgentStatusesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAgentStatusesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get agent statuses default response
func (o *GetAgentStatusesDefault) Code() int {
	return o._statusCode
}

func (o *GetAgentStatusesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{uuid}/statuses][%d] GetAgentStatuses default  %+v", o._statusCode, o.Payload)
}
func (o *GetAgentStatusesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetAgentStatusesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
