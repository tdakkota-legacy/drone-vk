FROM golang:1.13 AS builder

MAINTAINER tdakkota

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build
ENV CGO_ENABLED=0
RUN go generate ./... && go fmt ./...
RUN go build -v -o drone-vk cmd/plugin.go

FROM plugins/base:linux-amd64

LABEL maintainer="tdakkota"
LABEL org.label-schema.name="Drone VK"
LABEL org.label-schema.vendor="tdakkota"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.vcs-url="https://github.com/tdakkota/drone-vk"

COPY --from=builder /app/drone-vk /bin/drone-vk

# Run the binary
ENTRYPOINT ["/bin/drone-vk"]