ARG GO_VERSION=1.22

FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /build
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/restapi

FROM debian:bookworm-slim AS final
COPY --from=builder /build/api_server /opt/app/
ENTRYPOINT ["/opt/app/sample-api"]