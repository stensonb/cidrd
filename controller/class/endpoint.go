package class

import (
  "github.com/emicklei/go-restful"

  "github.com/stensonb/cidrd/config"
  "github.com/stensonb/cidrd/model"
  "github.com/stensonb/cidrd/controller/plugin"
)

const path_name = "class"
const param_name = path_name + "-id"
const param_name_curly = "{" + param_name + "}"

type classEndpoint struct {
  Model *model.Model
}

func init() {
	e := &classEndpoint{}
	plugin.RegisterEndpoint(e)
}

func (ce *classEndpoint) Name() string {
  return path_name
}

func (ce *classEndpoint) Register(c *restful.Container) {
  ws := new(restful.WebService)
	ws.
		Path("/" + path_name).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

  // define routes and their handlers
	ws.Route(ws.GET("/").To(ce.getAllClasses))
	ws.Route(ws.GET("/" + param_name_curly).To(ce.getClass))
	ws.Route(ws.POST("/").To(ce.createClass))
	ws.Route(ws.POST("/" + param_name_curly).To(ce.updateClass))
	ws.Route(ws.PUT("/" + param_name_curly).To(ce.updateClass))
	ws.Route(ws.DELETE("/" + param_name_curly).To(ce.removeClass))

  // add this webservice to the container
	c.Add(ws)
}

func (ce *classEndpoint) Configure(model *model.Model, config *config.Config) {
  ce.Model = model
}
