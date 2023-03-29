package handlers

import (
	"net/http"
	
	"github.com/dhnikolas/mock/pkg/response"
)

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	response.WriteBody(w, http.StatusOK, []byte{})
	return
}
