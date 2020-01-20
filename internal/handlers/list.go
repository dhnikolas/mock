package handlers

import (
	"mock/pkg/jsonconfig"
	"mock/pkg/response"
	"net/http"
)

func (h *Handler) ListMock(w http.ResponseWriter, r *http.Request) {
	responseBody, err := jsonconfig.GetConfigFileBody()
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteBody(w, http.StatusOK, responseBody)
	return
}
