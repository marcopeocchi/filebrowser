package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	fsDomain "github.com/marcopeocchi/filebrowser/internal/fs"
	"github.com/marcopeocchi/filebrowser/internal/middlewares"
)

var (
	//go:embed app/dist
	vue  embed.FS
	root string
	port int
)

func init() {
	flag.StringVar(&root, "r", ".", "Root directory")
	flag.IntVar(&port, "p", 8080, "Port to listen at")
	flag.Parse()
}

func main() {
	s := newServer()
	go gracefulShutdown(s)

	fmt.Printf("Listening on http://localhost:%d\n", port)
	s.ListenAndServe()
}

func newServer() *http.Server {
	r := chi.NewRouter()
	r.Use(middlewares.CORS)
	r.Use(middleware.Logger)

	build, _ := fs.Sub(vue, "app/dist")

	sh := middlewares.SpaHandler{
		Entrypoint: "index.html",
		Filesystem: &build,
	}

	fsContainer := fsDomain.Container(root)

	r.Route("/api", func(r chi.Router) {
		r.Post("/walk", fsContainer.WalkDir())
		r.Get("/open/{id}", fsContainer.OpenFile())
		r.Get("/basepath/length", fsContainer.GetBasePathLength())
	})

	r.Get("/*", sh.Handler())

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
}

func gracefulShutdown(s *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()
		fmt.Println()
		log.Println("shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)

		defer func() {
			stop()
			cancel()
			log.Println("shutdown completed")
		}()

		s.Shutdown(ctxTimeout)
	}()
}
