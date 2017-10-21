FROM golang:1.9-alpine AS builder

COPY . /go/src/github.com/tsub/echo-sandbox
WORKDIR /go/src/github.com/tsub/echo-sandbox

RUN apk add --update \
        make \
        git \
        glide && \
    make deps && \
    make build

FROM alpine:3.6

ENV PORT=3000

COPY --from=builder /go/src/github.com/tsub/echo-sandbox/build/echo-sandbox /echo-sandbox

CMD ["./echo-sandbox"]
