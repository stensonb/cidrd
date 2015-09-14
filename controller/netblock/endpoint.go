package netblock

import (
	"github.com/emicklei/go-restful"

	"github.com/stensonb/cidrd/config"
	"github.com/stensonb/cidrd/controller/plugin"
	"github.com/stensonb/cidrd/model"
)

const path_name = "netblock"
const param_name = path_name + "-id"
const param_name_curly = "{" + param_name + "}"

type netblockEndpoint struct {
	Model *model.Model
}

func init() {
	e := &netblockEndpoint{}
	plugin.RegisterEndpoint(e)
}

func (nbe *netblockEndpoint) Name() string {
	return path_name
}

func (nbe *netblockEndpoint) Register(c *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/" + path_name).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	// define routes and their handlers
	ws.Route(ws.GET("/").To(nbe.getAllNetblocks))
	ws.Route(ws.GET("/" + param_name_curly).To(nbe.getNetblock))
	ws.Route(ws.POST("/").To(nbe.createNetblock))
	ws.Route(ws.POST("/" + param_name_curly).To(nbe.updateNetblock))
	ws.Route(ws.PUT("/" + param_name_curly).To(nbe.updateNetblock))
	ws.Route(ws.DELETE("/" + param_name_curly).To(nbe.removeNetblock))

	// add this webservice to the container
	c.Add(ws)
}

func (nbe *netblockEndpoint) Configure(model *model.Model, config *config.Config) {
	nbe.Model = model
}
