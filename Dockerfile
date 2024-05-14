FROM golang:1.22-alpine@sha256:2a882244fb51835ebbd8313bffee83775b0c076aaf56b497b43d8a4c72db65e1 as build
COPY . /code
WORKDIR /code
ENV CGO_ENABLED=0
ENV GOPROXY=https://proxy.golang.org
RUN go mod download && \
    go build -o app

FROM alpine:3.18@sha256:11e21d7b981a59554b3f822c49f6e9f57b6068bb74f49c4cd5cc4c663c7e5160
RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /code/app /app/app
WORKDIR /app
ENTRYPOINT [ "/app/app" ]
CMD [ "web" ]
