package handlers

import (
	"encoding/json"
	"io/ioutil"
	"mock/pkg/jsonconfig"
	"mock/pkg/response"
	"net/http"
)

func (h *Handler) UpdateMock(w http.ResponseWriter, r *http.Request) {

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
		response.JSONError(w, http.StatusBadRequest, "Empty Id")
		return
	}

	if rb.Url == "" {
		response.JSONError(w, http.StatusBadRequest, "url is empty ")
		return
	}

	if rb.Body == "" {
		response.JSONError(w, http.StatusBadRequest, "body is empty ")
		return
	}

	if rb.Method == "" {
		response.JSONError(w, http.StatusBadRequest, "method is empty ")
		return
	}

	updatedMocks, isUpdated, err := jsonconfig.UpdateConfig(rb)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !isUpdated {
		response.JSONError(w, http.StatusBadRequest, "nothing to update")
		return
	}

	h.ConfigMap = jsonconfig.ConfigMap(updatedMocks)

	jsonBody, err := json.Marshal(rb)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteBody(w, http.StatusOK, jsonBody)
	return

}
