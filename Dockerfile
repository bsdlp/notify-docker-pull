FROM golang:1.15-alpine AS builder
ADD . /src/
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags='-extldflags=-static' -o /bin/entrypoint

FROM alpine:3.10

COPY LICENSE README.md /

COPY --from=builder /bin/entrypoint /bin/entrypoint

ENTRYPOINT ["bin/entrypoint"]
