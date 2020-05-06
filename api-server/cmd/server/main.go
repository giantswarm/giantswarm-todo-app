package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"contrib.go.opencensus.io/exporter/ocagent"
	"github.com/go-chi/chi"
	"github.com/piontec/go-chi-middleware-server/pkg/server"
	"github.com/piontec/go-chi-middleware-server/pkg/server/middleware"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/trace"
	"go.opencensus.io/zpages"

	"github.com/giantswarm/giantswarm-todo-app/api-server/pkg/todo"
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

func initTracing(config *todo.Config) {
	oce, err := ocagent.NewExporter(
		ocagent.WithInsecure(),
		ocagent.WithReconnectionPeriod(5*time.Second),
		ocagent.WithAddress(config.OcAgentHost),
		ocagent.WithServiceName("api-server"))
	if err != nil {
		log.Fatalf("Failed to create ocagent-exporter: %v", err)
	}
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(oce)
}

func main() {
	config := todo.NewConfig()
	initTracing(config)

	go func() {
		mux := http.NewServeMux()
		zpages.Handle(mux, "/debug")
		log.Fatal(http.ListenAndServe(":8081", mux))
	}()

	server := server.NewChiServer(func(r *chi.Mux) {
		r.Use(func(handler http.Handler) http.Handler {
			return &ochttp.Handler{
				Handler:          handler,
				IsPublicEndpoint: false,
				Propagation:      &b3.HTTPFormat{},
				IsHealthEndpoint: func(r *http.Request) bool {
					if r.URL.Path == "/ping" {
						return true
					}
					return false
				},
			}
		})
		r.Route("/v1", func(r chi.Router) {
			r.Mount("/todo",
				todo.NewRouter(config.TodoURL).GetRouter())
		})
	}, &server.ChiServerOptions{
		HTTPPort:              8080,
		DisableOIDCMiddleware: true,
		LoggerFields: logrus.Fields{
			"ver": version,
		},
		LoggerFieldFuncs: middleware.LogrusFieldFuncs{
			"traceID": func(r *http.Request) string {
				span := trace.FromContext(r.Context())
				return span.SpanContext().TraceID.String()
			},
		},
	})
	printVersion(server.GetLogger())
	server.Run()
}
