package controller

import (
	"github.com/emicklei/go-restful"
	"log"

	"github.com/stensonb/cidrd/config"
	"github.com/stensonb/cidrd/model"

	// the controller endpoints
	_ "github.com/stensonb/cidrd/controller/class"
	_ "github.com/stensonb/cidrd/controller/netblock"
	"github.com/stensonb/cidrd/controller/plugin"
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
