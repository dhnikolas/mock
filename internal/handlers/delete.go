package handlers

import (
	"encoding/json"
	"io/ioutil"
	"mock/pkg/jsonconfig"
	"mock/pkg/response"
	"net/http"
	"strings"
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

	if rb.Url == "" {
		response.JSONError(w, http.StatusBadRequest, "url is empty ")
		return
	}

	if rb.Method == "" {
		response.JSONError(w, http.StatusBadRequest, "method is empty ")
		return
	}

	url := strings.Trim(rb.Url, "/")
	delete(h.ConfigMap, url)

	isDeleted, err := jsonconfig.RemoveFromConfig(url, rb.Method)
	if !isDeleted{
		response.JSONError(w, http.StatusBadRequest, "nothing to delete")
		return
	}

	response.JSON(w, http.StatusOK, "Mock successfully deleted")
	return

}
