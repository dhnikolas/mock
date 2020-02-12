package handlers

import (
	"fmt"
	"html/template"
	"mock/public"
	"net/http"
)

func (h *Handler) UserInterface(w http.ResponseWriter, r *http.Request) {
	fmt.Println(public.Path() + "/index.html")
	t, err := template.ParseFiles(public.Path() + "/index.html")
	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}

