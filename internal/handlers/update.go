package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	
	"mock/internal/dto"
	"mock/pkg/response"
)

func (h *Handler) UpdateMock(w http.ResponseWriter, r *http.Request) {

	rb := &dto.Mock{}
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

	_, err = neturl.ParseRequestURI(rb.Url)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if rb.Method == "" {
		response.JSONError(w, http.StatusBadRequest, "method is empty ")
		return
	}
	
	isUpdated, err := h.Mock.Update(rb)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !isUpdated {
		response.JSONError(w, http.StatusBadRequest, "nothing to update")
		return
	}

	jsonBody, err := json.Marshal(rb)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteBody(w, http.StatusOK, jsonBody)
	return

}
