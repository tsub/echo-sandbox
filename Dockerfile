FROM golang:1.9-alpine AS builder

COPY . /go/src/github.com/tsub/echo-sandbox
WORKDIR /go/src/github.com/tsub/echo-sandbox

RUN apk add --update \
        ca-certificates \
        make \
        git \
        glide && \
    make deps && \
    make build_prod

FROM scratch

ENV PORT=3000 \
    DB_HOST=postgres \
    DB_USER=postgres \
    DB_NAME=dev.echo-sandbox

COPY --from=builder /go/src/github.com/tsub/echo-sandbox/build/echo-sandbox /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./echo-sandbox"]
