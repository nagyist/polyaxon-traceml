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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// BookmarkProjectReader is a Reader for the BookmarkProject structure.
type BookmarkProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookmarkProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookmarkProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewBookmarkProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBookmarkProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBookmarkProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBookmarkProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBookmarkProjectOK creates a BookmarkProjectOK with default headers values
func NewBookmarkProjectOK() *BookmarkProjectOK {
	return &BookmarkProjectOK{}
}

/* BookmarkProjectOK describes a response with status code 200, with default header values.

A successful response.
*/
type BookmarkProjectOK struct {
}

func (o *BookmarkProjectOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectOK ", 200)
}

func (o *BookmarkProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBookmarkProjectNoContent creates a BookmarkProjectNoContent with default headers values
func NewBookmarkProjectNoContent() *BookmarkProjectNoContent {
	return &BookmarkProjectNoContent{}
}

/* BookmarkProjectNoContent describes a response with status code 204, with default header values.

No content.
*/
type BookmarkProjectNoContent struct {
	Payload interface{}
}

func (o *BookmarkProjectNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectNoContent  %+v", 204, o.Payload)
}
func (o *BookmarkProjectNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkProjectForbidden creates a BookmarkProjectForbidden with default headers values
func NewBookmarkProjectForbidden() *BookmarkProjectForbidden {
	return &BookmarkProjectForbidden{}
}

/* BookmarkProjectForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type BookmarkProjectForbidden struct {
	Payload interface{}
}

func (o *BookmarkProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectForbidden  %+v", 403, o.Payload)
}
func (o *BookmarkProjectForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkProjectNotFound creates a BookmarkProjectNotFound with default headers values
func NewBookmarkProjectNotFound() *BookmarkProjectNotFound {
	return &BookmarkProjectNotFound{}
}

/* BookmarkProjectNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type BookmarkProjectNotFound struct {
	Payload interface{}
}

func (o *BookmarkProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] bookmarkProjectNotFound  %+v", 404, o.Payload)
}
func (o *BookmarkProjectNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkProjectDefault creates a BookmarkProjectDefault with default headers values
func NewBookmarkProjectDefault(code int) *BookmarkProjectDefault {
	return &BookmarkProjectDefault{
		_statusCode: code,
	}
}

/* BookmarkProjectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BookmarkProjectDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the bookmark project default response
func (o *BookmarkProjectDefault) Code() int {
	return o._statusCode
}

func (o *BookmarkProjectDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{name}/bookmark][%d] BookmarkProject default  %+v", o._statusCode, o.Payload)
}
func (o *BookmarkProjectDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *BookmarkProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
