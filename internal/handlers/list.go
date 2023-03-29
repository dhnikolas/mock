package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/dhnikolas/mock/pkg/response"
	"github.com/go-chi/chi"
)

func (h *Handler) ListMock(w http.ResponseWriter, r *http.Request) {
	mocks, err := h.Mock.GetAll()
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	responseBody, err := json.Marshal(mocks)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteBody(w, http.StatusOK, responseBody)
	return
}

func (h *Handler) ListLogRequests(w http.ResponseWriter, r *http.Request) {
	mockId := chi.URLParam(r, "mockId")
	if len(mockId) < 1 {
		response.WriteBody(w, http.StatusNotFound, []byte{})
		return
	}
	
	logRequests, err := h.LogRequest.GetByMockId(mockId)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	responseBody, err := json.Marshal(logRequests)
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	response.WriteBody(w, http.StatusOK, responseBody)
	return
}
