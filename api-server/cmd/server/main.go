package main

import (
	"os"
	"runtime"

	"github.com/go-chi/chi"
	"github.com/piontec/go-chi-middleware-server/pkg/server"
	"github.com/sirupsen/logrus"

	"github.com/giantswarm/blog-i-want-it-all/api-server/pkg/todo"
)

var (
	version = "v0.1.0-dev-build"
	commit  = "none"
	date    = "unknown"
)

func printVersion(l *logrus.Logger) {
	l.Infof("apiserver version: %s, commit: %s, build date: %s", version, commit, date)
	l.Infof("apiserver Go Version: %s, OS/Arch: %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func main() {
	todoURL := os.Getenv("TODO_URL")
	if todoURL == "" {
		panic("Required environment variable 'TODO_URL' not set")
	}
	server := server.NewChiServer(func(r *chi.Mux) {
		r.Route("/v1", func(r chi.Router) {
			r.Mount("/todo", todo.NewRouter(todoURL).GetRouter())
		})
	}, &server.ChiServerOptions{
		HTTPPort:              8080,
		DisableOIDCMiddleware: true,
		LoggerFields: logrus.Fields{
			"ver": version,
		},
	})
	printVersion(server.GetLogger())
	server.Run()
}
