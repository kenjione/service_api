FROM golang:latest as builder

ARG SSH_PRIVATE_KEY

RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ stretch-pgdg main"| tee /etc/apt/sources.list.d/pgdg.list
RUN curl -sS https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
RUN apt-get update -qq && apt-get install -y build-essential apt-transport-https apt-utils \
  postgresql-client-10

WORKDIR /builder/
ADD . /builder/
RUN mkdir -p ~/.ssh && umask 0077 && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
  && git config --global url."git@bitbucket.org:".insteadOf https://bitbucket.org/ \
  && git config --global url."git@github.com:".insteadOf https://github.com/ \
  && ssh-keyscan bitbucket.org >> ~/.ssh/known_hosts \
  && ssh-keyscan github.com >> ~/.ssh/known_hosts

COPY . .

RUN go get -d -v ./...