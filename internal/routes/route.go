package routes

import (
	"log"
	"mock/internal/handlers"
	"mock/pkg/jsonconfig"
	"net/http"
	"syscall"

	"github.com/go-chi/chi"
)

func Init(cm map[string][]*jsonconfig.Mock) {
	r := chi.NewRouter()

	h := &handlers.Handler{ConfigMap: cm}

	r.Get("/v1/mock/", h.ListMock)
	r.Post("/v1/mock/", h.AddMock)
	r.Patch("/v1/mock/", h.UpdateMock)
	r.Delete("/v1/mock/", h.DeleteMock)

	r.Get("/*", h.Index)
	r.Post("/*", h.Index)
	r.Patch("/*", h.Index)
	r.Put("/*", h.Index)
	r.Options("/*", h.Index)
	r.Trace("/*", h.Index)
	r.Head("/*", h.Index)
	r.Delete("/*", h.Index)

	port, found := syscall.Getenv("CURRENT_PORT")
	if !found {
		port = "8111"
	}

	log.Fatal(http.ListenAndServe(":" + port, r))
}
