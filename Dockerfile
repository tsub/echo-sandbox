FROM golang:1.9-alpine AS builder

RUN apk add --update \
        ca-certificates \
        make \
        git \
        glide

WORKDIR /go/src/github.com/tsub/echo-sandbox
COPY glide* Makefile /go/src/github.com/tsub/echo-sandbox/
RUN make deps

COPY . /go/src/github.com/tsub/echo-sandbox/
RUN make build_prod

FROM scratch

ENV PORT=3000 \
    DB_HOST=postgres \
    DB_USER=postgres \
    DB_NAME=dev.echo-sandbox

COPY --from=builder /go/src/github.com/tsub/echo-sandbox/build/echo-sandbox /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./echo-sandbox"]
