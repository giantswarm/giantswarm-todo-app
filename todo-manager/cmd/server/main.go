package main

import (
	"log"
	"runtime"
	"time"

	"contrib.go.opencensus.io/exporter/ocagent"
	grpcserver "github.com/piontec/grpc-middleware-server/pkg/server"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"

	todomgrpb "github.com/giantswarm/giantswarm-todo-app/todo-manager/pkg/proto"
	todomgr "github.com/giantswarm/giantswarm-todo-app/todo-manager/pkg/server"
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
}

func main() {
	config := todomgr.NewConfig()
	initTracing(config)
	todoMgr := todomgr.NewTodoManagerServer(config)

	server := grpcserver.NewGrpcServer(func(server *grpc.Server) {
		todomgrpb.RegisterTodoManagerServer(server, todoMgr)
	}, &grpcserver.GrpcServerOptions{
		LoggerFields: logrus.Fields{
			"ver": version,
		},
		AdditionalOptions: []grpc.ServerOption{
			grpc.StatsHandler(&ocgrpc.ServerHandler{}),
		},
	})
	printVersion(server.GetLogger())
	server.Run()
	todoMgr.Stop()
}
