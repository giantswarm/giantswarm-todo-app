package main

import (
	"net/http"
	"runtime"
	"time"

	"contrib.go.opencensus.io/exporter/ocagent"
	grpcserver "github.com/piontec/grpc-middleware-server/pkg/server"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"go.opencensus.io/zpages"
	"google.golang.org/grpc"

	todomgrpb "github.com/giantswarm/giantswarm-todo-app/todo-manager/pkg/proto"
	todomgr "github.com/giantswarm/giantswarm-todo-app/todo-manager/pkg/server"
)

var (
	version = "v0.1.0-dev-build"
	commit  = "none"
	date    = "unknown"
)

func printVersion(l *log.Logger) {
	l.Infof("todomanager version: %s, commit: %s, build date: %s", version, commit, date)
	l.Infof("todomanager Go Version: %s, OS/Arch: %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func initTracing(config *todomgr.Config) {
	oce, err := ocagent.NewExporter(
		ocagent.WithInsecure(),
		ocagent.WithReconnectionPeriod(5*time.Second),
		ocagent.WithAddress(config.OcAgentHost),
		ocagent.WithServiceName("todo-manager"))
	if err != nil {
		log.Fatalf("Failed to create ocagent-exporter: %v", err)
	}
	trace.RegisterExporter(oce)

	go func() {
		mux := http.NewServeMux()
		zpages.Handle(mux, "/debug")
		log.Fatal(http.ListenAndServe(":8081", mux))
	}()
}

func main() {
	config := todomgr.NewConfig()
	if config.EnableTracing {
		initTracing(config)
	}
	todoMgr := todomgr.NewTodoManagerServer(config)

	server := grpcserver.NewGrpcServer(func(server *grpc.Server) {
		todomgrpb.RegisterTodoManagerServer(server, todoMgr)
	}, &grpcserver.GrpcServerOptions{
		LoggerFields: log.Fields{
			"ver": version,
		},
		AdditionalOptions: []grpc.ServerOption{
			grpc.StatsHandler(&ocgrpc.ServerHandler{}),
		},
		MetricsPort: 8080,
	})
	printVersion(server.GetLogger())
	if config.EnableFailures {
		server.GetLogger().Warn("Failures Middleware is enabled")
	}
	server.GetLogger().Infof("Tracing instrumentation is %v", config.EnableTracing)
	server.Run()
	todoMgr.Stop()
}
