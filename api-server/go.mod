module github.com/giantswarm/giantswarm-todo-app/api-server

go 1.13

require (
	contrib.go.opencensus.io/exporter/ocagent v0.7.0
	github.com/albertogviana/prometheus-middleware v0.0.1
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/golang/protobuf v1.4.3
	github.com/piontec/go-chi-middleware-server v0.1.2
	github.com/prometheus/client_golang v1.11.1
	github.com/sirupsen/logrus v1.6.0
	go.opencensus.io v0.22.3
	google.golang.org/api v0.26.0 // indirect
	google.golang.org/genproto v0.0.0-20200604104852-0b0486081ffb // indirect
	google.golang.org/grpc v1.29.1
)
