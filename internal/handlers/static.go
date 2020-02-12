package handlers

import (
	"mock/public"
	"net/http"
	"net/url"
	"strings"
)

func (h *Handler) Static(w http.ResponseWriter, r *http.Request) {
	prefix := "/static/"
	fp := public.Path() + prefix
	fs := http.FileServer(http.Dir(fp))
	if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		r2.URL.Path = p
		fs.ServeHTTP(w, r2)
	} else {
		http.NotFound(w, r)
	}
}
