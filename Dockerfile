ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}-alpine AS build-env

ENV CGO_ENABLED 0

ENV GOOS linux

WORKDIR /src

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -installsuffix 'cgo' -o main .

FROM golang:${GO_VERSION}-alpine

RUN mkdir /opt/ddbooking
WORKDIR /opt/ddbooking

COPY --from=build-env /src/main /opt/ddbooking

RUN addgroup -g 1001 dev
RUN adduser -S -u 1001 -G dev dev

RUN chown -R dev:dev /opt/ddbooking
RUN chown -R dev:dev /var/log

EXPOSE 8080

USER dev:dev

ENTRYPOINT ["/opt/ddbooking/main"]
