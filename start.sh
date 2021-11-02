#!/bin/sh

# build and start service
echo '=====> Build && start services'
docker-compose up -d --build
