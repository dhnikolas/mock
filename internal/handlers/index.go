package handlers

import (
	"mock/pkg/jsonconfig"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	ConfigMap map[string][]*jsonconfig.Mock
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	url := strings.Trim(r.URL.Path, "/")
	m, ok := h.ConfigMap[url]
	if ok {
		for _, cm := range m {
			if strings.ToUpper(cm.Method) == r.Method {
				h.mockResponse(w, cm)
				return
			}
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("page not found"))
	return
}

func (h *Handler) mockResponse(w http.ResponseWriter, m *jsonconfig.Mock) {
	if len(m.Headers) > 0 {
		for _, h := range m.Headers {
			w.Header().Set(h.Name, h.Value)
		}
	}
	if len(m.ContentType) > 0 {
		w.Header().Set("Content-Type", m.ContentType)
	}
	statusInt, errStatus := strconv.Atoi(m.Status)
	if m.Status == "" || errStatus != nil{
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(statusInt)
	}

	w.Write([]byte(m.Body))
}
