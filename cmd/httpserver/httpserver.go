package httpserver

import (
	"github.com/Trafilea/cmd/dependencies"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router     *chi.Mux
	httpServer *http.Server
}

func Start(d dependencies.Definition) {
	log.Info("server started")

	router := SetupRouter(d)

	server := SetupHttpServer(router)

	server.ListenAndServe()
}

func SetupRouter(d dependencies.Definition) *chi.Mux {
	router := chi.NewRouter()

	router.Use(render.SetContentType(render.ContentTypeJSON))

	Routes(router, d)

	return router
}

func SetupHttpServer(router *chi.Mux) *Server {
	return &Server{
		router: router,
		httpServer: &http.Server{
			Addr:    ":8000",
			Handler: router,
		},
	}
}

func (s *Server) ListenAndServe() {
	if err := http.ListenAndServe(s.httpServer.Addr, s.router); err != nil {
		log.Info("server stopped")
		log.Fatal(err)
	}
}
