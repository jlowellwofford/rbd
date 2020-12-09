// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Rbd rbd describes an RBD map.  To successfully map, at least one monitor, pool and image must be specified.
// Additionally, you will need options.name and options.secret specified.
//
//
// swagger:model rbd
type Rbd struct {

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// image
	// Required: true
	// Min Length: 1
	Image *string `json:"image"`

	// monitors
	// Required: true
	Monitors []strfmt.IPv4 `json:"monitors"`

	// options
	Options *RbdOptions `json:"options,omitempty"`

	// pool
	// Required: true
	// Min Length: 1
	Pool *string `json:"pool"`

	// refs
	// Read Only: true
	Refs int64 `json:"refs,omitempty"`

	// snapshot
	Snapshot string `json:"snapshot,omitempty"`
}

// Validate validates this rbd
func (m *Rbd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rbd) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if err := validate.MinLength("image", "body", string(*m.Image), 1); err != nil {
		return err
	}

	return nil
}

func (m *Rbd) validateMonitors(formats strfmt.Registry) error {

	if err := validate.Required("monitors", "body", m.Monitors); err != nil {
		return err
	}

	for i := 0; i < len(m.Monitors); i++ {

		if err := validate.Pattern("monitors"+"."+strconv.Itoa(i), "body", string(m.Monitors[i]), `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`); err != nil {
			return err
		}

		if err := validate.FormatOf("monitors"+"."+strconv.Itoa(i), "body", "ipv4", m.Monitors[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Rbd) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *Rbd) validatePool(formats strfmt.Registry) error {

	if err := validate.Required("pool", "body", m.Pool); err != nil {
		return err
	}

	if err := validate.MinLength("pool", "body", string(*m.Pool), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rbd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rbd) UnmarshalBinary(b []byte) error {
	var res Rbd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
