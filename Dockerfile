FROM golang:1.9-alpine AS builder

COPY . /go/src/github.com/tsub/echo-sandbox
WORKDIR /go/src/github.com/tsub/echo-sandbox

RUN apk add --update \
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

COPY --from=builder /go/src/github.com/tsub/echo-sandbox/build/echo-sandbox /echo-sandbox

CMD ["./echo-sandbox"]
