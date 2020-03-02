package controller

import (
	"github.com/hfuss/dns-operator/pkg/controller/multivaluednsrecord"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, multivaluednsrecord.Add)
}
