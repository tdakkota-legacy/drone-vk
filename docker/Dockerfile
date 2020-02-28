FROM golang:1.13

MAINTAINER tdakkota

# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build
ENV CGO_ENABLED=0
RUN go generate ./... && go fmt ./...
RUN go build -v cmd/plugin.go

ENTRYPOINT ["./plugin"]