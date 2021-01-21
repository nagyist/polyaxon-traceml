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

package component_hub_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// BookmarkComponentHubReader is a Reader for the BookmarkComponentHub structure.
type BookmarkComponentHubReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookmarkComponentHubReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookmarkComponentHubOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewBookmarkComponentHubNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBookmarkComponentHubForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBookmarkComponentHubNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBookmarkComponentHubDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBookmarkComponentHubOK creates a BookmarkComponentHubOK with default headers values
func NewBookmarkComponentHubOK() *BookmarkComponentHubOK {
	return &BookmarkComponentHubOK{}
}

/* BookmarkComponentHubOK describes a response with status code 200, with default header values.

A successful response.
*/
type BookmarkComponentHubOK struct {
}

func (o *BookmarkComponentHubOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/bookmark][%d] bookmarkComponentHubOK ", 200)
}

func (o *BookmarkComponentHubOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBookmarkComponentHubNoContent creates a BookmarkComponentHubNoContent with default headers values
func NewBookmarkComponentHubNoContent() *BookmarkComponentHubNoContent {
	return &BookmarkComponentHubNoContent{}
}

/* BookmarkComponentHubNoContent describes a response with status code 204, with default header values.

No content.
*/
type BookmarkComponentHubNoContent struct {
	Payload interface{}
}

func (o *BookmarkComponentHubNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/bookmark][%d] bookmarkComponentHubNoContent  %+v", 204, o.Payload)
}
func (o *BookmarkComponentHubNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkComponentHubNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkComponentHubForbidden creates a BookmarkComponentHubForbidden with default headers values
func NewBookmarkComponentHubForbidden() *BookmarkComponentHubForbidden {
	return &BookmarkComponentHubForbidden{}
}

/* BookmarkComponentHubForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type BookmarkComponentHubForbidden struct {
	Payload interface{}
}

func (o *BookmarkComponentHubForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/bookmark][%d] bookmarkComponentHubForbidden  %+v", 403, o.Payload)
}
func (o *BookmarkComponentHubForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkComponentHubForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkComponentHubNotFound creates a BookmarkComponentHubNotFound with default headers values
func NewBookmarkComponentHubNotFound() *BookmarkComponentHubNotFound {
	return &BookmarkComponentHubNotFound{}
}

/* BookmarkComponentHubNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type BookmarkComponentHubNotFound struct {
	Payload interface{}
}

func (o *BookmarkComponentHubNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/bookmark][%d] bookmarkComponentHubNotFound  %+v", 404, o.Payload)
}
func (o *BookmarkComponentHubNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkComponentHubNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkComponentHubDefault creates a BookmarkComponentHubDefault with default headers values
func NewBookmarkComponentHubDefault(code int) *BookmarkComponentHubDefault {
	return &BookmarkComponentHubDefault{
		_statusCode: code,
	}
}

/* BookmarkComponentHubDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BookmarkComponentHubDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the bookmark component hub default response
func (o *BookmarkComponentHubDefault) Code() int {
	return o._statusCode
}

func (o *BookmarkComponentHubDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/hub/{name}/bookmark][%d] BookmarkComponentHub default  %+v", o._statusCode, o.Payload)
}
func (o *BookmarkComponentHubDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *BookmarkComponentHubDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
