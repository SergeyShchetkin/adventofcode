FROM golang:1.17.6-alpine3.14

# Maintainer Info
LABEL maintainer="Sergey Shchetkin <mrschetkin@gmail.com>"

# Install suuport tools
RUN apk add --no-cache \
        bash nano curl netcat-openbsd iputils

# Move to working directory
WORKDIR /go/src/app

# Copy the code into the container
COPY . .

# Download packages
RUN apk add --no-cache --virtual .ext-deps build-base \
    && go mod download \
    && go mod vendor \
    && apk del .ext-deps