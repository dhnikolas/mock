package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"
	
	"mock/internal/dto"
	"mock/internal/repository/logrequest"
	"mock/internal/repository/mock"
	"mock/pkg/response"
)

type Handler struct {
	Mock *mock.Repository
	LogRequest *logrequest.Repository
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	dump := h.dumpRequest(r)
	url := strings.TrimRight(r.URL.Path, "/")
	mockResult, err := h.Mock.GetByUrl(url)
	if err != nil {
		response.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(*mockResult) > 0 {
		for _, cm := range *mockResult {
			if strings.ToUpper(cm.Method) == r.Method {
				go func(dump string) {
					err := h.LogRequest.Add(&dto.LogRequest{
						MockId: cm.Id,
						Body:   dump,
					})
					if err != nil {
						fmt.Println(err)
					}
				}(dump)
				h.mockResponse(w, cm)
				return
			}
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	fmt.Println(dump)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("page not found"))
	return
}

func (h *Handler) mockResponse(w http.ResponseWriter, m *dto.Mock) {
	if len(m.Headers) > 0 {
		for _, h := range m.Headers {
			w.Header().Set(h.Name, h.Value)
		}
	}
	if len(m.ContentType) > 0 {
		w.Header().Set("Content-Type", m.ContentType)
	}
	statusInt, errStatus := strconv.Atoi(m.Status)
	if m.Status == "" || errStatus != nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(statusInt)
	}
	
	w.Write([]byte(m.Body))
}

func (h *Handler) dumpRequest(r *http.Request) string {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	currentTime := time.Now().Format(time.RFC3339Nano)
	return fmt.Sprintf("Request %s:\n%s\n\n", currentTime, string(reqDump))
}
