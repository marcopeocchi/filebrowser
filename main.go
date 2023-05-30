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
	authDomain "github.com/marcopeocchi/filebrowser/internal/auth"
	"github.com/marcopeocchi/filebrowser/internal/db"
	fsDomain "github.com/marcopeocchi/filebrowser/internal/fs"
	"github.com/marcopeocchi/filebrowser/internal/middlewares"
)

var (
	//go:embed app/dist
	vue    embed.FS
	dbpath string
	root   string
	port   int
)

func init() {
	flag.StringVar(&dbpath, "db", "./db.json", "JSON database file path")
	flag.StringVar(&root, "r", ".", "Root directory")
	flag.IntVar(&port, "p", 8080, "Port to listen at")
	flag.Parse()
}

func main() {
	jdb, err := db.NewJsonDB(dbpath)
	if err != nil {
		log.Fatalln(err)
	}

	s := newServer(jdb)
	go gracefulShutdown(s)

	fmt.Printf("Listening on http://localhost:%d\n", port)
	s.ListenAndServe()
}

func newServer(jdb *db.JsonDB) *http.Server {
	r := chi.NewRouter()
	r.Use(middlewares.CORS)
	r.Use(middleware.Logger)

	build, _ := fs.Sub(vue, "app/dist")

	sh := middlewares.SpaHandler{
		Entrypoint: "index.html",
		Filesystem: &build,
	}

	sh.AddClientRoute("/files")

	fsContainer := fsDomain.Container(root)
	authContainer := authDomain.Container(jdb)

	r.Route("/api", func(r chi.Router) {
		// require auth routes
		r.With(middlewares.Authenticated).Post("/walk", fsContainer.WalkDir())
		r.With(middlewares.Authenticated).Get("/open/{id}", fsContainer.OpenFile())
		// unauthenticated routes
		r.Get("/basepath/length", fsContainer.GetBasePathLength())
		r.Post("/login", authContainer.Login())
		r.Get("/logout", authContainer.Logout())
	})

	r.Get("/*", sh.Handler())

	return &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
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
