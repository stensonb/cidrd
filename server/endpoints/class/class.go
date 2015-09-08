package class

import (
  "net/http"
  "github.com/emicklei/go-restful"
  dao "github.com/stensonb/cidrd/models/class"
)

type ClassEndpoint struct {
}

const path_name = "class"
const param_name = path_name + "-id"
const param_name_curly = "{" + param_name + "}"

func New() *ClassEndpoint {
  return &ClassEndpoint{}
}

func (ce *ClassEndpoint) Register(c *restful.Container) {
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

func (ce *ClassEndpoint) getAllClasses(req *restful.Request, res *restful.Response) {
  err := validateGetAllRequest(req)
  if err != nil {
    res.WriteError(http.StatusBadRequest, err)
  }

  ans, err := dao.GetAllClasses()
  if err != nil {
    res.WriteError(http.StatusNotFound, err)
  } else {
    res.WriteEntity(ans)
  }
}

func validateGetAllRequest(req *restful.Request) error {
  return nil
}

func (ce *ClassEndpoint) getClass(req *restful.Request, res *restful.Response) {
  err := validateGetRequest(req)
  if err != nil {
    //error, bad request
    res.WriteError(http.StatusBadRequest, err)
    return
  }

  ans, err := dao.GetClassByUUID(req.PathParameter(param_name))
  if err != nil {
    //error from dao
    res.WriteError(http.StatusNotFound, err)
  } else {
    res.WriteEntity(ans)
  }
}

func validateGetRequest(req *restful.Request) error {
  // this is for GET request validation only
  // keep model validation with the model
  return nil
}

func (ce *ClassEndpoint) updateClass(req *restful.Request, res *restful.Response) {
  err := validateUpdateRequest(req)
  if err != nil {
    res.WriteError(http.StatusBadRequest, err)
    return
  }

  newclass := new(dao.Class)
	err = req.ReadEntity(newclass)
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
  newclass.SetUUID(req.PathParameter(param_name))

  dao.Update(newclass)
	res.WriteHeader(http.StatusAccepted)
	res.WriteEntity(newclass)
}

func validateUpdateRequest(res *restful.Request) error {
  // this is for PUT request validation only
  // keep model validated with the model
  return nil
}

func (ce *ClassEndpoint) createClass(req *restful.Request, res *restful.Response) {
  err := validateCreateRequest(req)
  if err != nil {
    res.WriteError(http.StatusBadRequest, err)
    return
  }

  newclass := new(dao.Class)
	err = req.ReadEntity(newclass)
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

  dao.Create(newclass)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(newclass)
}

func validateCreateRequest(req *restful.Request) error {
  // this is for POST request validation only
  // keep model validated with the model
  return nil
}

func (ce *ClassEndpoint) removeClass(req *restful.Request, res *restful.Response) {
  err := validateDeleteRequest(req)
  if err != nil {
    res.WriteError(http.StatusBadRequest, err)
    return
  }

  err = dao.DeleteClassByUUID(req.PathParameter(param_name))
  if err != nil {
    //error from dao
    res.WriteError(http.StatusNotFound, err)
  } else {
	  res.WriteHeader(http.StatusAccepted)
    res.Write([]byte("Deleted."))
  }
}

func validateDeleteRequest(req *restful.Request) error {
  // this is for DELETE request validation only
  // keep model validated with the model
  return nil
}
