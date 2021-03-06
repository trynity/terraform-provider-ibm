// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostSecurityGroupsParamsBodyVpc VPCIdentityByName
//
// The VPC the server is to be a part of
// swagger:model postSecurityGroupsParamsBodyVpc
type PostSecurityGroupsParamsBodyVpc struct {

	// The CRN for this VPC
	Crn string `json:"crn,omitempty"`

	// The unique identifier for this VPC
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this VPC
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`
}

// Validate validates this post security groups params body vpc
func (m *PostSecurityGroupsParamsBodyVpc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostSecurityGroupsParamsBodyVpc) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostSecurityGroupsParamsBodyVpc) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostSecurityGroupsParamsBodyVpc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSecurityGroupsParamsBodyVpc) UnmarshalBinary(b []byte) error {
	var res PostSecurityGroupsParamsBodyVpc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
