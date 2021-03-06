// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VolumeEncryptionKey VolumeEncryptionKey
// swagger:model volumeEncryptionKey
type VolumeEncryptionKey struct {

	// The CRN for this key
	Crn string `json:"crn,omitempty"`
}

// Validate validates this volume encryption key
func (m *VolumeEncryptionKey) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeEncryptionKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeEncryptionKey) UnmarshalBinary(b []byte) error {
	var res VolumeEncryptionKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
