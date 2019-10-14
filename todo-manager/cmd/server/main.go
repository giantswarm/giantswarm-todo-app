package main

import (
	"runtime"

	grpcserver "github.com/piontec/grpc-middleware-server/pkg/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	todomgr "github.com/giantswarm/blog-i-want-it-all/todo-manager/pkg/server"
	todomgrpb "github.com/giantswarm/blog-i-want-it-all/todo-manager/pkg/proto"
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
	server := grpcserver.NewGrpcServer(func(server *grpc.Server) {
		todoMgr := &todomgr.TodoManagerServer{}
		todomgrpb.RegisterTodoManagerServer(server, todoMgr)
	}, &grpcserver.GrpcServerOptions{
		LoggerFields: logrus.Fields{
			"ver": version,
		},
	})
	printVersion(server.GetLogger())
	server.Run()
}
