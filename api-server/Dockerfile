FROM golang:1.13.1-alpine as builder
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
RUN adduser -D -H nouser
USER nouser
COPY --from=builder /tmp/src/bin/apiserver /apiserver
CMD ["/apiserver"]
