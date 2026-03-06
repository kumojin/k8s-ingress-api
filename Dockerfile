FROM golang:1.26-alpine@sha256:2389ebfa5b7f43eeafbd6be0c3700cc46690ef842ad962f6c5bd6be49ed82039 as build
COPY . /code
WORKDIR /code
ENV CGO_ENABLED=0
ENV GOPROXY=https://proxy.golang.org
RUN go mod download && \
    go build -o app

FROM alpine:3.23@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659
RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /code/app /app/app
WORKDIR /app
ENTRYPOINT [ "/app/app" ]
CMD [ "web" ]
