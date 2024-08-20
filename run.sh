#!/bin/bash

# Set the directory where the Dockerfile is located
DOCKERFILE_DIR="/home/ajabbour/Desktop/docker"

# Set the name for the Docker image
IMAGE_NAME="ascii-art-web"

# Set the tag for the Docker image
IMAGE_TAG="latest"

# Change to the directory containing the Dockerfile
cd "$DOCKERFILE_DIR"

# Build the Docker image
docker build -t "$IMAGE_NAME:$IMAGE_TAG" .
