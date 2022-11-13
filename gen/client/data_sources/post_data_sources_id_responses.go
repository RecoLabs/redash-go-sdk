// Code generated by go-swagger; DO NOT EDIT.

package data_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/recolabs/redash-go-sdk/gen/models"
)

// PostDataSourcesIDReader is a Reader for the PostDataSourcesID structure.
type PostDataSourcesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDataSourcesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDataSourcesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostDataSourcesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostDataSourcesIDOK creates a PostDataSourcesIDOK with default headers values
func NewPostDataSourcesIDOK() *PostDataSourcesIDOK {
	return &PostDataSourcesIDOK{}
}

/* PostDataSourcesIDOK describes a response with status code 200, with default header values.

OK
*/
type PostDataSourcesIDOK struct {
	Payload *models.DataSource
}

func (o *PostDataSourcesIDOK) Error() string {
	return fmt.Sprintf("[POST /data_sources/{id}][%d] postDataSourcesIdOK  %+v", 200, o.Payload)
}
func (o *PostDataSourcesIDOK) GetPayload() *models.DataSource {
	return o.Payload
}

func (o *PostDataSourcesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDataSourcesIDDefault creates a PostDataSourcesIDDefault with default headers values
func NewPostDataSourcesIDDefault(code int) *PostDataSourcesIDDefault {
	return &PostDataSourcesIDDefault{
		_statusCode: code,
	}
}

/* PostDataSourcesIDDefault describes a response with status code -1, with default header values.

error
*/
type PostDataSourcesIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post data sources ID default response
func (o *PostDataSourcesIDDefault) Code() int {
	return o._statusCode
}

func (o *PostDataSourcesIDDefault) Error() string {
	return fmt.Sprintf("[POST /data_sources/{id}][%d] PostDataSourcesID default  %+v", o._statusCode, o.Payload)
}
func (o *PostDataSourcesIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostDataSourcesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostDataSourcesIDBody post data sources ID body
swagger:model PostDataSourcesIDBody
*/
type PostDataSourcesIDBody struct {

	// name
	Name string `json:"name,omitempty"`

	// options
	Options interface{} `json:"options,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this post data sources ID body
func (o *PostDataSourcesIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post data sources ID body based on context it is used
func (o *PostDataSourcesIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDataSourcesIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDataSourcesIDBody) UnmarshalBinary(b []byte) error {
	var res PostDataSourcesIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}