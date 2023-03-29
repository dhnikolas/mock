package handlers

import (
	"html/template"
	"net/http"
	
	"github.com/dhnikolas/mock/public"
)

func (h *Handler) UserInterface(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(public.Path() + "/index.html")
	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}

