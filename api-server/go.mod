module github.com/giantswarm/giantswarm-todo-app/api-server

go 1.13

require (
	contrib.go.opencensus.io/exporter/ocagent v0.6.0
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/golang/protobuf v1.3.2
	github.com/piontec/go-chi-middleware-server v0.1.1
	github.com/sirupsen/logrus v1.4.2
	go.opencensus.io v0.22.1
	google.golang.org/grpc v1.24.0
)
