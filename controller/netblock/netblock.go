package netblock

import (
	"github.com/emicklei/go-restful"
	"net/http"

	"github.com/stensonb/cidrd/model"
)

func (nbe *netblockEndpoint) getAllNetblocks(req *restful.Request, res *restful.Response) {
	err := validateGetAllRequest(req)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
	}

	ans, err := nbe.Model.GetAllNetblocks()
	if err != nil {
		res.WriteError(http.StatusNotFound, err)
	} else {
		res.WriteEntity(ans)
	}
}

func validateGetAllRequest(req *restful.Request) error {
	return nil
}

func (nbe *netblockEndpoint) getNetblock(req *restful.Request, res *restful.Response) {
	err := validateGetRequest(req)
	if err != nil {
		//error, bad request
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	ans, err := nbe.Model.GetNetblockByUUID(req.PathParameter(param_name))
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

func (nbe *netblockEndpoint) updateNetblock(req *restful.Request, res *restful.Response) {
	err := validateUpdateRequest(req)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	newnetblock := new(model.Netblock)
	err = req.ReadEntity(newnetblock)
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	filterNetblock(newnetblock)

	newnetblock.Uuid = req.PathParameter(param_name)

	nbe.Model.StoreNetblock(newnetblock)
	res.WriteHeader(http.StatusAccepted)
	res.WriteEntity(newnetblock)
}

func validateUpdateRequest(res *restful.Request) error {
	// this is for PUT request validation only
	// keep model validated with the model
	return nil
}

// this function removes data from the Netblock object
// before saving (fields which can never be set by
// the user...uuid, created, modified...)
func filterNetblock(n *model.Netblock) {
	_newnetblock := new(model.Netblock)
	n.Created = _newnetblock.Created
	n.Modified = _newnetblock.Modified
	n.Uuid = _newnetblock.Uuid
}

func (nbe *netblockEndpoint) createNetblock(req *restful.Request, res *restful.Response) {
	err := validateCreateRequest(req)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	newnetblock := new(model.Netblock)
	err = req.ReadEntity(newnetblock)
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	// ensure created isn't set (users can never set this data)
	filterNetblock(newnetblock)

	err = nbe.Model.StoreNetblock(newnetblock)
	if err != nil {
		res.WriteError(http.StatusInternalServerError, err)
	}

	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(newnetblock)
}

func validateCreateRequest(req *restful.Request) error {
	// this is for POST request validation only
	// keep model validated with the model
	return nil
}

func (nbe *netblockEndpoint) removeNetblock(req *restful.Request, res *restful.Response) {
	err := validateDeleteRequest(req)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}

	err = nbe.Model.DeleteNetblockByUUID(req.PathParameter(param_name))
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
