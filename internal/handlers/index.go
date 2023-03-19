package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"
	
	"mock/pkg/jsonconfig"
)

type Handler struct {
	ConfigMap map[string][]*jsonconfig.Mock
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	h.dumpRequest(r)
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
	if m.Status == "" || errStatus != nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(statusInt)
	}
	
	w.Write([]byte(m.Body))
}

func (h *Handler) dumpRequest(r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	currentTime := time.Now().Format(time.RFC3339Nano)
	fmt.Printf("Request %s:\n%s\n\n", currentTime, string(reqDump))
}
