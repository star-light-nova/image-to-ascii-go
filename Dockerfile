FROM golang:1.22.1-alpine3.19

WORKDIR /usr/src/itaii

RUN apk update \
    && apk upgrade \
    && apk add \
    build-base \
    vips-dev \
    && rm -rf /var/cache/apk/*

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

ENV GIN_MODE="release"

RUN go build -v -o itaii

EXPOSE 8080

CMD [ "sh", "./entrypoint.sh" ]
