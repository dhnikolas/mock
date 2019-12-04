package handlers

import (
	"mock/pkg/jsonconfig"
	"net/http"
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
	if len(m.ContentType) < 1 {
		w.Header().Set("Content-Type", "application/json")
	} else {
		w.Header().Set("Content-Type", m.ContentType)
	}

	if m.Status == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(m.Status)
	}

	w.Write([]byte(m.Body))
}
