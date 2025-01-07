FROM golang:1.23-alpine@sha256:13aaa4b92fd4dc81683816b4b62041442e9f685deeb848897ce78c5e2fb03af7 as build
COPY . /code
WORKDIR /code
ENV CGO_ENABLED=0
ENV GOPROXY=https://proxy.golang.org
RUN go mod download && \
    go build -o app

FROM alpine:3.21@sha256:b97e2a89d0b9e4011bb88c02ddf01c544b8c781acf1f4d559e7c8f12f1047ac3
RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /code/app /app/app
WORKDIR /app
ENTRYPOINT [ "/app/app" ]
CMD [ "web" ]
