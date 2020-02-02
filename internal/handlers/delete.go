package handlers

import (
	"encoding/json"
	"io/ioutil"
	"mock/pkg/jsonconfig"
	"mock/pkg/response"
	"net/http"
)

func (h *Handler) DeleteMock(w http.ResponseWriter, r *http.Request) {

	rb := &jsonconfig.Mock{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, rb)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, "unmarshal json error ")
		return
	}

	if rb.Id == "" {
		response.JSONError(w, http.StatusBadRequest, "Id is empty ")
		return
	}

	newMocks, isDeleted, err := jsonconfig.RemoveFromConfig(rb.Id)
	if !isDeleted{
		response.JSONError(w, http.StatusBadRequest, "nothing to delete")
		return
	}

	h.ConfigMap = jsonconfig.ConfigMap(newMocks)

	response.JSON(w, http.StatusOK, "Mock successfully deleted")
	return

}
