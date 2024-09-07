FROM golang:1.23-alpine@sha256:ac67716dd016429be8d4c2c53a248d7bcdf06d34127d3dc451bda6aa5a87bc06 as build
COPY . /code
WORKDIR /code
ENV CGO_ENABLED=0
ENV GOPROXY=https://proxy.golang.org
RUN go mod download && \
    go build -o app

FROM alpine:3.20@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d
RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /code/app /app/app
WORKDIR /app
ENTRYPOINT [ "/app/app" ]
CMD [ "web" ]
