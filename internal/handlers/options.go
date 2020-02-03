package handlers

import (
	"mock/pkg/response"
	"net/http"
)

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	response.WriteBody(w, http.StatusOK, []byte{})
	return
}
