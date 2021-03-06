package class

import (
	"github.com/emicklei/go-restful"
	"net/http"

	"github.com/stensonb/cidrd/model"
)

func (ce *classEndpoint) getAllClasses(req *restful.Request, res *restful.Response) {
	err := validateGetAllRequest(req)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
	}

	ans, err := ce.Model.GetAllClasses()
	if err != nil {
		res.WriteError(http.StatusNotFound, err)
	} else {
		res.WriteEntity(ans)
	}
}

func validateGetAllRequest(req *restful.Request) error {
	return nil
}

func (ce *classEndpoint) getClass(req *restful.Request, res *restful.Response) {
	err := validateGetRequest(req)
	if err != nil {
		//error, bad request
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	ans, err := ce.Model.GetClassByUUID(req.PathParameter(param_name))
	if err != nil {
		//error from model
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

func (ce *classEndpoint) updateClass(req *restful.Request, res *restful.Response) {
	err := validateUpdateRequest(req)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	newclass := new(model.Class)
	err = req.ReadEntity(newclass)
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	filterClass(newclass)

	newclass.Uuid = req.PathParameter(param_name)

	ce.Model.StoreClass(newclass)
	res.WriteHeader(http.StatusAccepted)
	res.WriteEntity(newclass)
}

func validateUpdateRequest(res *restful.Request) error {
	// this is for PUT request validation only
	// keep model validated with the model
	return nil
}

// this function removes data from the Class object
// before saving (fields which can never be set by
// the user...uuid, created, modified...)
func filterClass(c *model.Class) {
	_newclass := new(model.Class)
	c.Created = _newclass.Created
	c.Modified = _newclass.Modified
	c.Uuid = _newclass.Uuid
}

func (ce *classEndpoint) createClass(req *restful.Request, res *restful.Response) {
	err := validateCreateRequest(req)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	newclass := new(model.Class)
	err = req.ReadEntity(newclass)
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	// ensure created isn't set (users can never set this data)
	filterClass(newclass)

	ce.Model.StoreClass(newclass)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(newclass)
}

func validateCreateRequest(req *restful.Request) error {
	// this is for POST request validation only
	// keep model validated with the model
	return nil
}

func (ce *classEndpoint) removeClass(req *restful.Request, res *restful.Response) {
	err := validateDeleteRequest(req)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	err = ce.Model.DeleteClassByUUID(req.PathParameter(param_name))
	if err != nil {
		//error from model
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
