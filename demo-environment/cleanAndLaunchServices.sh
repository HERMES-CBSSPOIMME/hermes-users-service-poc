#!/usr/bin/env bash

# TODO: Change credentials handling method

# Constants
PROJECT=hermes-demo

printf "\n"
echo "====================================================================================================="
echo "Removing previous containers ..."
echo "====================================================================================================="

# Stop and remove previous container
docker-compose down

printf "\n"
echo "====================================================================================================="
echo "Building & starting containers ..."
echo "====================================================================================================="
# Build & start services
docker-compose build && docker-compose up -d
