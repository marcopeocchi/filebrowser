package main

import (
	"embed"
	"flag"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	fsDomain "github.com/marcopeocchi/filebrowser/internal/fs"
	"github.com/marcopeocchi/filebrowser/internal/middlewares"
)

var (
	//go:embed app/dist
	vue  embed.FS
	root string
)

func init() {
	flag.StringVar(&root, "r", ".", "Root directory")
	flag.Parse()
}

func main() {
	runBlocking()
}

func runBlocking() {
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

	http.ListenAndServe(":8080", r)
}

/*
func gracefulShutdown(app *fiber.App) {
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()
		log.Println("shutdown signal received")

		defer func() {
			db.Persist()
			stop()
			app.ShutdownWithTimeout(time.Second * 5)
		}()
	}()
}
*/
