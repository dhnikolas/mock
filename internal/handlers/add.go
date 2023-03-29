package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strings"
	
	"github.com/dhnikolas/mock/internal/dto"
	"github.com/dhnikolas/mock/pkg/response"
	"github.com/dhnikolas/mock/third_party/utils"
	"github.com/google/uuid"
)

func (h *Handler) AddMock(w http.ResponseWriter, r *http.Request) {

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

	if rb.Status == "" {
		response.JSONError(w, http.StatusBadRequest, "status is empty ")
		return
	}

	url := strings.TrimRight(rb.Url, "/")
	rb.Url = url
	if len(rb.Id) < 1 {
		id := uuid.New()
		rb.Id = id.String()
	}
	
	_, ok, _ := h.Mock.GetByUrlAndMethod(url, rb.Method)
	if ok {
		response.JSONError(w, http.StatusBadRequest, "Mock already exist")
		return
	}
	err = h.Mock.Add(rb)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	jsonBody, err := json.Marshal(rb)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteBody(w, http.StatusOK, jsonBody)
	return
}
