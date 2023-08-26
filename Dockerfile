FROM golang:latest

WORKDIR /usr/src/app

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash

RUN apt-get update
RUN apt-get install -y migrate
RUN apt-get install -y postgresql-client