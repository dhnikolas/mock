package routes

import (
	"log"
	"net/http"
	"os"
	
	"github.com/dhnikolas/mock/internal/handlers"
	"github.com/dhnikolas/mock/internal/repository/logrequest"
	"github.com/dhnikolas/mock/internal/repository/mock"
	"github.com/dhnikolas/mock/pkg/middleware"
	
	"github.com/go-chi/chi"
)

func Init(m *mock.Repository, l *logrequest.Repository) {
	r := chi.NewRouter()
	h := &handlers.Handler{Mock: m, LogRequest: l}
	r.Use(middleware.Cors)

	//API
	r.Get("/v1/mock/", h.ListMock)
	r.Get("/v1/mock/{mockId}/mock-requests", h.ListLogRequests)
	r.Post("/v1/mock/", h.AddMock)
	r.Patch("/v1/mock/", h.UpdateMock)
	r.Delete("/v1/mock/", h.DeleteMock)
	r.Delete("/v1/mock/{mockId}/mock-requests", h.DeleteLogRequests)
	r.Delete("/v1/mock/{mockId}/mock-requests/{logId}", h.DeleteLog)

	//UI
	r.Get("/ui/", h.UserInterface)
	r.Get("/static/*", h.Static)

	//Cors policy avoid
	r.Options("/*", h.Options)

	r.Get("/*", h.Index)
	r.Post("/*", h.Index)
	r.Patch("/*", h.Index)
	r.Put("/*", h.Index)
	r.Trace("/*", h.Index)
	r.Head("/*", h.Index)
	r.Delete("/*", h.Index)

	port := os.Getenv("CURRENT_PORT")
	found := len(port) > 0
	if !found {
		port = "8111"
	}

	log.Fatal(http.ListenAndServe(":" + port, r))
}
