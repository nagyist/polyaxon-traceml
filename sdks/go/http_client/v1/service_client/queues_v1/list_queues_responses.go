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

package queues_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListQueuesReader is a Reader for the ListQueues structure.
type ListQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListQueuesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListQueuesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListQueuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListQueuesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListQueuesOK creates a ListQueuesOK with default headers values
func NewListQueuesOK() *ListQueuesOK {
	return &ListQueuesOK{}
}

/* ListQueuesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListQueuesOK struct {
	Payload *service_model.V1ListQueuesResponse
}

func (o *ListQueuesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{name}/queues][%d] listQueuesOK  %+v", 200, o.Payload)
}
func (o *ListQueuesOK) GetPayload() *service_model.V1ListQueuesResponse {
	return o.Payload
}

func (o *ListQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListQueuesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQueuesNoContent creates a ListQueuesNoContent with default headers values
func NewListQueuesNoContent() *ListQueuesNoContent {
	return &ListQueuesNoContent{}
}

/* ListQueuesNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListQueuesNoContent struct {
	Payload interface{}
}

func (o *ListQueuesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{name}/queues][%d] listQueuesNoContent  %+v", 204, o.Payload)
}
func (o *ListQueuesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListQueuesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQueuesForbidden creates a ListQueuesForbidden with default headers values
func NewListQueuesForbidden() *ListQueuesForbidden {
	return &ListQueuesForbidden{}
}

/* ListQueuesForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListQueuesForbidden struct {
	Payload interface{}
}

func (o *ListQueuesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{name}/queues][%d] listQueuesForbidden  %+v", 403, o.Payload)
}
func (o *ListQueuesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListQueuesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQueuesNotFound creates a ListQueuesNotFound with default headers values
func NewListQueuesNotFound() *ListQueuesNotFound {
	return &ListQueuesNotFound{}
}

/* ListQueuesNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListQueuesNotFound struct {
	Payload interface{}
}

func (o *ListQueuesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{name}/queues][%d] listQueuesNotFound  %+v", 404, o.Payload)
}
func (o *ListQueuesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListQueuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQueuesDefault creates a ListQueuesDefault with default headers values
func NewListQueuesDefault(code int) *ListQueuesDefault {
	return &ListQueuesDefault{
		_statusCode: code,
	}
}

/* ListQueuesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListQueuesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list queues default response
func (o *ListQueuesDefault) Code() int {
	return o._statusCode
}

func (o *ListQueuesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/agents/{name}/queues][%d] ListQueues default  %+v", o._statusCode, o.Payload)
}
func (o *ListQueuesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListQueuesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
