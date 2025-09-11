# Multi-stage build for cross-platform support
FROM --platform=$BUILDPLATFORM golang:1.21-alpine AS builder

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH GOARM=${TARGETVARIANT#v} \
    go build -ldflags='-w -s' -o rpi_exporter .

# Final stage
FROM scratch

LABEL org.opencontainers.image.title="rpi_exporter"
LABEL org.opencontainers.image.description="Prometheus exporter for Raspberry Pi hardware metrics"
LABEL org.opencontainers.image.source="https://github.com/eurusik/rpi_export"
LABEL org.opencontainers.image.authors="cavaliercoder@github, eurusik@github"

EXPOSE 9110

COPY --from=builder /src/rpi_exporter /rpi_exporter

ENTRYPOINT ["/rpi_exporter"]
CMD ["-addr=:9110"]
