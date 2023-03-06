FROM golang:alpine as builder

RUN apk update \
  && apk --no-cache add --virtual build-dependencies \
  zlib-dev build-base linux-headers coreutils

ENV GOPROXY=https://goproxy.io,direct

COPY . /opt/

RUN cd /opt \
  && ls -al \
  && make build-static

FROM alpine
COPY --from=builder /opt/mdnscli-static /bin/mdnscli
