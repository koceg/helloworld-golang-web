FROM golang:1.13-alpine

ENV GOPATH /go
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org

COPY app.go /go/helloworld-golang-web/

RUN cd helloworld-golang-web \
    && go build app.go

FROM alpine:3.11

RUN adduser -g '' -h /client -u 1000 -D -s /sbin/nologin golang

EXPOSE 8080

USER golang

COPY --from=0 --chown=1000:1000 /go/helloworld-golang-web/app /client/app

CMD ["/client/app"]
