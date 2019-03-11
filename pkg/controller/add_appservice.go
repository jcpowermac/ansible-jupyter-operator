package controller

import (
	"github.com/jcpowermac/ansible-jupyter-operator/pkg/controller/appservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, appservice.Add)
}
