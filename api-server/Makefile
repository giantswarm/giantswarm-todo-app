# Image URL to use all building/pushing image targets
BIN ?= apiserver
IMG ?= ${BIN}:latest

all: test build

# Run tests
test: generate fmt vet
	go test ./... -coverprofile cover.out

# Build binary
build: generate fmt vet compile

compile:
	go build \
		-ldflags "\
		-X main.version=`git describe 2>/dev/null || echo "0.0.0"` \
		-X main.commit=`git rev-parse HEAD` \
		-X main.date=`date '+%FT%T%:z'`" \
		-o bin/${BIN} ./cmd/server/main.go

# Run directly
run: generate fmt vet
	go run ./cmd/server/main.go

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet -composites=false ./pkg/... ./cmd/...

# Generate code
generate:
	go generate ./pkg/... ./cmd/...

# Build the docker image from lacally built binary
docker-build: build
	docker build -t ${IMG} .

# Push the docker image
docker-push:
	docker push ${IMG}
