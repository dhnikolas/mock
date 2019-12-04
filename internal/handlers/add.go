package handlers

import (
	"encoding/json"
	"io/ioutil"
	"mock/pkg/jsonconfig"
	"mock/pkg/response"
	"mock/third_party/utils"
	"net/http"
	"strings"
)

func (h *Handler) AddMock(w http.ResponseWriter, r *http.Request) {

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

	if rb.Body == "" {
		response.JSONError(w, http.StatusBadRequest, "body is empty ")
		return
	}

	if rb.Method == "" {
		response.JSONError(w, http.StatusBadRequest, "method is empty ")
		return
	}

	methods := []string{http.MethodDelete,
		http.MethodGet,
		http.MethodPost,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodPut,
		http.MethodOptions,
		http.MethodTrace,}

	if !utils.Contains(methods, rb.Method) {
		response.JSONError(w, http.StatusBadRequest, "Method can be only " + strings.Join(methods, ", "))
		return
	}

	if rb.Status == 0 {
		response.JSONError(w, http.StatusBadRequest, "status is empty ")
		return
	}

	url := strings.Trim(rb.Url, "/")

	_, ok := h.ConfigMap[url]

	if ok {
		h.ConfigMap[url] = append(h.ConfigMap[url], rb)
	} else {
		h.ConfigMap[url] = []*jsonconfig.Mock{rb}
	}

	err = jsonconfig.AddToConfig(rb)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "New mock successfully added")
	return
}
