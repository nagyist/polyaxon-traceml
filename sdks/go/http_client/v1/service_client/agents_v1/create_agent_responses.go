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

// CreateAgentReader is a Reader for the CreateAgent structure.
type CreateAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAgentOK creates a CreateAgentOK with default headers values
func NewCreateAgentOK() *CreateAgentOK {
	return &CreateAgentOK{}
}

/* CreateAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAgentOK struct {
	Payload *service_model.V1Agent
}

func (o *CreateAgentOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentOK  %+v", 200, o.Payload)
}
func (o *CreateAgentOK) GetPayload() *service_model.V1Agent {
	return o.Payload
}

func (o *CreateAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Agent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentNoContent creates a CreateAgentNoContent with default headers values
func NewCreateAgentNoContent() *CreateAgentNoContent {
	return &CreateAgentNoContent{}
}

/* CreateAgentNoContent describes a response with status code 204, with default header values.

No content.
*/
type CreateAgentNoContent struct {
	Payload interface{}
}

func (o *CreateAgentNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentNoContent  %+v", 204, o.Payload)
}
func (o *CreateAgentNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentForbidden creates a CreateAgentForbidden with default headers values
func NewCreateAgentForbidden() *CreateAgentForbidden {
	return &CreateAgentForbidden{}
}

/* CreateAgentForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type CreateAgentForbidden struct {
	Payload interface{}
}

func (o *CreateAgentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentForbidden  %+v", 403, o.Payload)
}
func (o *CreateAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentNotFound creates a CreateAgentNotFound with default headers values
func NewCreateAgentNotFound() *CreateAgentNotFound {
	return &CreateAgentNotFound{}
}

/* CreateAgentNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type CreateAgentNotFound struct {
	Payload interface{}
}

func (o *CreateAgentNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] createAgentNotFound  %+v", 404, o.Payload)
}
func (o *CreateAgentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAgentDefault creates a CreateAgentDefault with default headers values
func NewCreateAgentDefault(code int) *CreateAgentDefault {
	return &CreateAgentDefault{
		_statusCode: code,
	}
}

/* CreateAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateAgentDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the create agent default response
func (o *CreateAgentDefault) Code() int {
	return o._statusCode
}

func (o *CreateAgentDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/agents][%d] CreateAgent default  %+v", o._statusCode, o.Payload)
}
func (o *CreateAgentDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
