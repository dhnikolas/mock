package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"github.com/dhnikolas/mock/internal/dto"
	"github.com/dhnikolas/mock/pkg/response"
	"github.com/go-chi/chi"
)

func (h *Handler) DeleteMock(w http.ResponseWriter, r *http.Request) {

	rb := &dto.Mock{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(body, rb)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, "unmarshal json error ")
		return
	}

	if rb.Id == "" {
		response.JSONError(w, http.StatusBadRequest, "Id is empty ")
		return
	}
	
	isDeleted, err := h.Mock.Delete(rb.Id)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, "Delete mock error:  " + err.Error())
		return
	}
	
	if !isDeleted{
		response.JSONError(w, http.StatusBadRequest, "nothing to delete")
		return
	}

	response.JSON(w, http.StatusOK, "Mock successfully deleted")
	return
}

func (h *Handler) DeleteLogRequests(w http.ResponseWriter, r *http.Request) {
	mockId := chi.URLParam(r, "mockId")
	if len(mockId) < 1 {
		response.WriteBody(w, http.StatusNotFound, []byte{})
		return
	}
	
	count, err := h.LogRequest.DeleteByMockId(mockId)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, "Delete log requests error:  " + err.Error())
		return
	}
	
	response.JSON(w, http.StatusOK, "Log requests successfully deleted: " + fmt.Sprint(count))
	return
}

func (h *Handler) DeleteLog(w http.ResponseWriter, r *http.Request) {
	mockId := chi.URLParam(r, "mockId")
	if len(mockId) < 1 {
		response.WriteBody(w, http.StatusNotFound, []byte{})
		return
	}
	
	logId := chi.URLParam(r, "logId")
	if len(logId) < 1 {
		response.WriteBody(w, http.StatusNotFound, []byte{})
		return
	}
	
	_, err := h.LogRequest.DeleteLog(mockId, logId)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, "Delete log requests error:  " + err.Error())
		return
	}
	
	response.JSON(w, http.StatusOK, "Log requests successfully deleted")
	return
}
