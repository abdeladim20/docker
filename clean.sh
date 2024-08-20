#!/bin/bash

# Remove all containers (stopped and running)
docker container rm -f $(docker container ls -aq)

# Remove all unused images
docker image prune -a -f

# Remove all unused volumes
docker volume prune -f

# Remove all unused networks
docker network prune -f