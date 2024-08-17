FROM golang:1.22.3-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the Go Modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o ascii-art-web

# Copy static files into the container
COPY Style /app/Style

#Provide defaults for an executing container. If an executable is not specified, then ENTRYPOINT must be specified as well. There can only be one CMD instruction in a Dockerfile.
CMD [ "./ASCII-ART-WEB" ]

# Define the network ports that this container will listen on at runtime.
EXPOSE 8000
