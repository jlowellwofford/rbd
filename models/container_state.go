// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ContainerState Valid container states
//
// swagger:model container_state
type ContainerState string

const (

	// ContainerStateCreated captures enum value "created"
	ContainerStateCreated ContainerState = "created"

	// ContainerStateRunning captures enum value "running"
	ContainerStateRunning ContainerState = "running"

	// ContainerStateRestarting captures enum value "restarting"
	ContainerStateRestarting ContainerState = "restarting"

	// ContainerStatePaused captures enum value "paused"
	ContainerStatePaused ContainerState = "paused"

	// ContainerStateExited captures enum value "exited"
	ContainerStateExited ContainerState = "exited"

	// ContainerStateDead captures enum value "dead"
	ContainerStateDead ContainerState = "dead"
)

// for schema
var containerStateEnum []interface{}

func init() {
	var res []ContainerState
	if err := json.Unmarshal([]byte(`["created","running","restarting","paused","exited","dead"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerStateEnum = append(containerStateEnum, v)
	}
}

func (m ContainerState) validateContainerStateEnum(path, location string, value ContainerState) error {
	if err := validate.EnumCase(path, location, value, containerStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this container state
func (m ContainerState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContainerStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this container state based on context it is used
func (m ContainerState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
