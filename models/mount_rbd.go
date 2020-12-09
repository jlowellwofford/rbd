// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MountRbd mount_rbd describes an RBD mount.  This must have at least and RBD ID associated with it (which becomes the mount's ID), and a provided filesystem type.
//
// swagger:model mount_rbd
type MountRbd struct {

	// fs type
	// Required: true
	FsType *string `json:"fs_type"`

	// must be a valid rbd device id
	// Required: true
	ID *int64 `json:"id"`

	// mount options
	MountOptions []string `json:"mount_options"`

	// mountpoint
	// Read Only: true
	Mountpoint string `json:"mountpoint,omitempty"`

	// ref
	// Read Only: true
	Ref int64 `json:"ref,omitempty"`
}

// Validate validates this mount rbd
func (m *MountRbd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountRbd) validateFsType(formats strfmt.Registry) error {

	if err := validate.Required("fs_type", "body", m.FsType); err != nil {
		return err
	}

	return nil
}

func (m *MountRbd) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MountRbd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountRbd) UnmarshalBinary(b []byte) error {
	var res MountRbd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
