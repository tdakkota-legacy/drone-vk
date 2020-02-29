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
RUN go build -v cmd/plugin.go

# create scratch image
FROM plugins/base:linux-amd64

# Copy our static executable.
COPY --from=builder ./plugin ./plugin

# Run the hello binary.
ENTRYPOINT ["./plugin"]