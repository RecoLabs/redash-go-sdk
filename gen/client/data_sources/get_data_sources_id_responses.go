// Code generated by go-swagger; DO NOT EDIT.

package data_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/recolabs/redash-go-sdk/gen/models"
)

// GetDataSourcesIDReader is a Reader for the GetDataSourcesID structure.
type GetDataSourcesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataSourcesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataSourcesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDataSourcesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDataSourcesIDOK creates a GetDataSourcesIDOK with default headers values
func NewGetDataSourcesIDOK() *GetDataSourcesIDOK {
	return &GetDataSourcesIDOK{}
}

/* GetDataSourcesIDOK describes a response with status code 200, with default header values.

Get data source by ID
*/
type GetDataSourcesIDOK struct {
	Payload *models.DataSource
}

func (o *GetDataSourcesIDOK) Error() string {
	return fmt.Sprintf("[GET /data_sources/{id}][%d] getDataSourcesIdOK  %+v", 200, o.Payload)
}
func (o *GetDataSourcesIDOK) GetPayload() *models.DataSource {
	return o.Payload
}

func (o *GetDataSourcesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataSourcesIDDefault creates a GetDataSourcesIDDefault with default headers values
func NewGetDataSourcesIDDefault(code int) *GetDataSourcesIDDefault {
	return &GetDataSourcesIDDefault{
		_statusCode: code,
	}
}

/* GetDataSourcesIDDefault describes a response with status code -1, with default header values.

error
*/
type GetDataSourcesIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get data sources ID default response
func (o *GetDataSourcesIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDataSourcesIDDefault) Error() string {
	return fmt.Sprintf("[GET /data_sources/{id}][%d] GetDataSourcesID default  %+v", o._statusCode, o.Payload)
}
func (o *GetDataSourcesIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDataSourcesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
