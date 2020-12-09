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

// ContainerNamespace Linux namespace
//
// swagger:model container_namespace
type ContainerNamespace string

const (

	// ContainerNamespaceCgroup captures enum value "cgroup"
	ContainerNamespaceCgroup ContainerNamespace = "cgroup"

	// ContainerNamespaceIpc captures enum value "ipc"
	ContainerNamespaceIpc ContainerNamespace = "ipc"

	// ContainerNamespaceNet captures enum value "net"
	ContainerNamespaceNet ContainerNamespace = "net"

	// ContainerNamespaceMnt captures enum value "mnt"
	ContainerNamespaceMnt ContainerNamespace = "mnt"

	// ContainerNamespacePid captures enum value "pid"
	ContainerNamespacePid ContainerNamespace = "pid"

	// ContainerNamespaceTime captures enum value "time"
	ContainerNamespaceTime ContainerNamespace = "time"

	// ContainerNamespaceUser captures enum value "user"
	ContainerNamespaceUser ContainerNamespace = "user"

	// ContainerNamespaceUts captures enum value "uts"
	ContainerNamespaceUts ContainerNamespace = "uts"
)

// for schema
var containerNamespaceEnum []interface{}

func init() {
	var res []ContainerNamespace
	if err := json.Unmarshal([]byte(`["cgroup","ipc","net","mnt","pid","time","user","uts"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerNamespaceEnum = append(containerNamespaceEnum, v)
	}
}

func (m ContainerNamespace) validateContainerNamespaceEnum(path, location string, value ContainerNamespace) error {
	if err := validate.EnumCase(path, location, value, containerNamespaceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this container namespace
func (m ContainerNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContainerNamespaceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this container namespace based on context it is used
func (m ContainerNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
