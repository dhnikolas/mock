package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	
	"mock/internal/dto"
	"mock/pkg/response"
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
