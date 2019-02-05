#!/usr/bin/env bash

# Constants

MONGODB_PORT=27017
MONGODB_ROOT_USER=root
MONGODB_ROOT_PASS=example

WAVE_USER=wave-demo-user
WAVE_USER_PASSWORD=example


WAVE_DB_NAME=waveDemoDB

echo 'Creating application user and db'

mongo ${WAVE_DB_NAME} \
        --host localhost \
        --port ${MONGODB_PORT} \
        -u ${MONGODB_ROOT_USER}  \
        -p ${MONGODB_ROOT_PASS} \
        --authenticationDatabase admin \
        --eval "db.createUser({user: '${WAVE_USER}', pwd: '${WAVE_USER_PASSWORD}', roles:[{role:'readWrite', db: '${WAVE_DB_NAME}'}]});"
