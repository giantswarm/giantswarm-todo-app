FROM golang:1.13.12-alpine as builder
ARG ver
ARG commit
RUN apk add make gcc build-base
WORKDIR /tmp/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV VER=$ver COMMIT=$commit
RUN make

FROM alpine:3.10  
WORKDIR /
RUN wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.0/grpc_health_probe-linux-amd64
RUN mv grpc_health_probe-linux-amd64 /usr/bin/grpc_health_probe
RUN chmod 755 /usr/bin/grpc_health_probe
RUN adduser -u 1010 -D -H nouser
USER 1010
COPY --from=builder /tmp/src/bin/todomanager /todomanager
CMD ["/todomanager"]
