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

package connections_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteConnectionReader is a Reader for the DeleteConnection structure.
type DeleteConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteConnectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteConnectionOK creates a DeleteConnectionOK with default headers values
func NewDeleteConnectionOK() *DeleteConnectionOK {
	return &DeleteConnectionOK{}
}

/* DeleteConnectionOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteConnectionOK struct {
}

func (o *DeleteConnectionOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/connections/{uuid}][%d] deleteConnectionOK ", 200)
}

func (o *DeleteConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConnectionNoContent creates a DeleteConnectionNoContent with default headers values
func NewDeleteConnectionNoContent() *DeleteConnectionNoContent {
	return &DeleteConnectionNoContent{}
}

/* DeleteConnectionNoContent describes a response with status code 204, with default header values.

No content.
*/
type DeleteConnectionNoContent struct {
	Payload interface{}
}

func (o *DeleteConnectionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/connections/{uuid}][%d] deleteConnectionNoContent  %+v", 204, o.Payload)
}
func (o *DeleteConnectionNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteConnectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConnectionForbidden creates a DeleteConnectionForbidden with default headers values
func NewDeleteConnectionForbidden() *DeleteConnectionForbidden {
	return &DeleteConnectionForbidden{}
}

/* DeleteConnectionForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type DeleteConnectionForbidden struct {
	Payload interface{}
}

func (o *DeleteConnectionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/connections/{uuid}][%d] deleteConnectionForbidden  %+v", 403, o.Payload)
}
func (o *DeleteConnectionForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConnectionNotFound creates a DeleteConnectionNotFound with default headers values
func NewDeleteConnectionNotFound() *DeleteConnectionNotFound {
	return &DeleteConnectionNotFound{}
}

/* DeleteConnectionNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type DeleteConnectionNotFound struct {
	Payload interface{}
}

func (o *DeleteConnectionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/connections/{uuid}][%d] deleteConnectionNotFound  %+v", 404, o.Payload)
}
func (o *DeleteConnectionNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConnectionDefault creates a DeleteConnectionDefault with default headers values
func NewDeleteConnectionDefault(code int) *DeleteConnectionDefault {
	return &DeleteConnectionDefault{
		_statusCode: code,
	}
}

/* DeleteConnectionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteConnectionDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete connection default response
func (o *DeleteConnectionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteConnectionDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/connections/{uuid}][%d] DeleteConnection default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteConnectionDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
