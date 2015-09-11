package controller

import (
	"log"
  "github.com/emicklei/go-restful"

	"github.com/stensonb/cidrd/config"
	"github.com/stensonb/cidrd/model"

  // the controller endpoints
	"github.com/stensonb/cidrd/controller/plugin"
	_ "github.com/stensonb/cidrd/controller/class"
)

// Controller contains the HTTP controller with all necessary dependencies
type Controller struct {
	config    *config.Config
	model     *model.Model
	eps       []plugin.Endpoint
	container *restful.Container
}

// New will create a new Controller
func New(config *config.Config, model *model.Model, container *restful.Container) *Controller {
	var f []plugin.Endpoint

	a := &Controller{
		config:    config,
		model:     model,
		eps:       f,
		container: container,
	}

	for _, ep := range plugin.Endpoints() {
		ep.Register(container)
		ep.Configure(model, config)
		a.eps = append(a.eps, ep)
		log.Println("Plugin configured:", ep.Name())
	}
	return a
}
