// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Visualization visualization
//
// swagger:model visualization
type Visualization struct {

	// description
	// Min Length: 1
	Description string `json:"description,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// is draft
	IsDraft bool `json:"is_draft,omitempty"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// options
	// Required: true
	Options interface{} `json:"options"`

	// query id
	// Required: true
	QueryID *int64 `json:"query_id"`

	// type
	// Required: true
	// Enum: [BOXPLOT CHART CHOROPLETH COHORT COUNTER DETAILS FUNNEL MAP PIVOT SANKEY SUNBURST_SEQUENCE TABLE WORD_CLOUD]
	Type *string `json:"type"`
}

// Validate validates this visualization
func (m *Visualization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Visualization) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", m.Description, 1); err != nil {
		return err
	}

	return nil
}

func (m *Visualization) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *Visualization) validateOptions(formats strfmt.Registry) error {

	if m.Options == nil {
		return errors.Required("options", "body", nil)
	}

	return nil
}

func (m *Visualization) validateQueryID(formats strfmt.Registry) error {

	if err := validate.Required("query_id", "body", m.QueryID); err != nil {
		return err
	}

	return nil
}

var visualizationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BOXPLOT","CHART","CHOROPLETH","COHORT","COUNTER","DETAILS","FUNNEL","MAP","PIVOT","SANKEY","SUNBURST_SEQUENCE","TABLE","WORD_CLOUD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		visualizationTypeTypePropEnum = append(visualizationTypeTypePropEnum, v)
	}
}

const (

	// VisualizationTypeBOXPLOT captures enum value "BOXPLOT"
	VisualizationTypeBOXPLOT string = "BOXPLOT"

	// VisualizationTypeCHART captures enum value "CHART"
	VisualizationTypeCHART string = "CHART"

	// VisualizationTypeCHOROPLETH captures enum value "CHOROPLETH"
	VisualizationTypeCHOROPLETH string = "CHOROPLETH"

	// VisualizationTypeCOHORT captures enum value "COHORT"
	VisualizationTypeCOHORT string = "COHORT"

	// VisualizationTypeCOUNTER captures enum value "COUNTER"
	VisualizationTypeCOUNTER string = "COUNTER"

	// VisualizationTypeDETAILS captures enum value "DETAILS"
	VisualizationTypeDETAILS string = "DETAILS"

	// VisualizationTypeFUNNEL captures enum value "FUNNEL"
	VisualizationTypeFUNNEL string = "FUNNEL"

	// VisualizationTypeMAP captures enum value "MAP"
	VisualizationTypeMAP string = "MAP"

	// VisualizationTypePIVOT captures enum value "PIVOT"
	VisualizationTypePIVOT string = "PIVOT"

	// VisualizationTypeSANKEY captures enum value "SANKEY"
	VisualizationTypeSANKEY string = "SANKEY"

	// VisualizationTypeSUNBURSTSEQUENCE captures enum value "SUNBURST_SEQUENCE"
	VisualizationTypeSUNBURSTSEQUENCE string = "SUNBURST_SEQUENCE"

	// VisualizationTypeTABLE captures enum value "TABLE"
	VisualizationTypeTABLE string = "TABLE"

	// VisualizationTypeWORDCLOUD captures enum value "WORD_CLOUD"
	VisualizationTypeWORDCLOUD string = "WORD_CLOUD"
)

// prop value enum
func (m *Visualization) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, visualizationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Visualization) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this visualization based on the context it is used
func (m *Visualization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Visualization) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Visualization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Visualization) UnmarshalBinary(b []byte) error {
	var res Visualization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}