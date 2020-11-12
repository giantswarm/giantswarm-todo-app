module github.com/giantswarm/giantswarm-todo-app/api-server

go 1.13

require (
	contrib.go.opencensus.io/exporter/ocagent v0.7.0
	github.com/albertogviana/prometheus-middleware v0.0.1
	github.com/giantswarm/giantswarm-todo-app/todo-manager v0.0.0-20201112102441-ba1c9188359a // indirect
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/piontec/go-chi-middleware-server v0.1.2
	github.com/prometheus/client_golang v1.7.1
	github.com/sirupsen/logrus v1.4.2
	go.opencensus.io v0.22.3
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0 // indirect
)
