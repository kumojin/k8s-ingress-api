FROM golang:1.24-alpine@sha256:b4f875e650466fa0fe62c6fd3f02517a392123eea85f1d7e69d85f780e4db1c1 as build
COPY . /code
WORKDIR /code
ENV CGO_ENABLED=0
ENV GOPROXY=https://proxy.golang.org
RUN go mod download && \
    go build -o app

FROM alpine:3.22@sha256:8a1f59ffb675680d47db6337b49d22281a139e9d709335b492be023728e11715
RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /code/app /app/app
WORKDIR /app
ENTRYPOINT [ "/app/app" ]
CMD [ "web" ]
