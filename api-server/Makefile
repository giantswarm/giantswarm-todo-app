# Image URL to use all building/pushing image targets
BIN ?= apiserver
IMG ?= quay.io/giantswarm/giantswarm-todo-${BIN}

export VER ?= $(shell git describe 2>/dev/null || echo "0.0.0")
export COMMIT ?= $(shell git rev-parse HEAD 2>/dev/null || echo "0000000")
export SHORT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "0000000")
export DATE ?= $(shell date '+%FT%T%:z')

IMG_VER ?= ${VER}-${SHORT_COMMIT}

all: test build

# Run tests
test: fmt vet
	go test ./... -coverprofile cover.out

# Build binary
build: fmt vet compile

compile:
	go build \
		-ldflags "\
		-X main.version=${VER} \
		-X main.commit=${COMMIT} \
		-X main.date=${DATE}" \
		-o bin/${BIN} ./cmd/server/main.go

# Run directly
run: fmt vet
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

# Build the docker image from locally built binary
docker-build: 
	docker build --build-arg ver=${VER} --build-arg commit=${COMMIT} . -t ${IMG}:latest -t ${IMG}:${IMG_VER}
	echo ${IMG}:${IMG_VER} > .${BIN}.version

# Push the docker image
docker-push:
	docker push ${IMG}:${IMG_VER}
