FROM golang:1.11-alpine AS builder

RUN apk add --update \
        ca-certificates \
        make \
        git

WORKDIR /app
COPY go.mod go.sum Makefile /app/
RUN make deps

COPY . /app/
RUN make build_prod

FROM scratch

ENV PORT=3000 \
    DB_HOST=postgres \
    DB_USER=postgres \
    DB_NAME=dev.echo-sandbox

COPY --from=builder /app/build/echo-sandbox /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./echo-sandbox"]
