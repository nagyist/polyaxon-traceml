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

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListTeamsReader is a Reader for the ListTeams structure.
type ListTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListTeamsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTeamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListTeamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTeamsOK creates a ListTeamsOK with default headers values
func NewListTeamsOK() *ListTeamsOK {
	return &ListTeamsOK{}
}

/* ListTeamsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListTeamsOK struct {
	Payload *service_model.V1ListTeamsResponse
}

func (o *ListTeamsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams][%d] listTeamsOK  %+v", 200, o.Payload)
}
func (o *ListTeamsOK) GetPayload() *service_model.V1ListTeamsResponse {
	return o.Payload
}

func (o *ListTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListTeamsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsNoContent creates a ListTeamsNoContent with default headers values
func NewListTeamsNoContent() *ListTeamsNoContent {
	return &ListTeamsNoContent{}
}

/* ListTeamsNoContent describes a response with status code 204, with default header values.

No content.
*/
type ListTeamsNoContent struct {
	Payload interface{}
}

func (o *ListTeamsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams][%d] listTeamsNoContent  %+v", 204, o.Payload)
}
func (o *ListTeamsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListTeamsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsForbidden creates a ListTeamsForbidden with default headers values
func NewListTeamsForbidden() *ListTeamsForbidden {
	return &ListTeamsForbidden{}
}

/* ListTeamsForbidden describes a response with status code 403, with default header values.

You don't have permission to access the resource.
*/
type ListTeamsForbidden struct {
	Payload interface{}
}

func (o *ListTeamsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams][%d] listTeamsForbidden  %+v", 403, o.Payload)
}
func (o *ListTeamsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsNotFound creates a ListTeamsNotFound with default headers values
func NewListTeamsNotFound() *ListTeamsNotFound {
	return &ListTeamsNotFound{}
}

/* ListTeamsNotFound describes a response with status code 404, with default header values.

Resource does not exist.
*/
type ListTeamsNotFound struct {
	Payload interface{}
}

func (o *ListTeamsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams][%d] listTeamsNotFound  %+v", 404, o.Payload)
}
func (o *ListTeamsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListTeamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsDefault creates a ListTeamsDefault with default headers values
func NewListTeamsDefault(code int) *ListTeamsDefault {
	return &ListTeamsDefault{
		_statusCode: code,
	}
}

/* ListTeamsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListTeamsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list teams default response
func (o *ListTeamsDefault) Code() int {
	return o._statusCode
}

func (o *ListTeamsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/teams][%d] ListTeams default  %+v", o._statusCode, o.Payload)
}
func (o *ListTeamsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListTeamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
