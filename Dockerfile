FROM golang:alpine as builder
MAINTAINER awgh <awgh@awgh.org>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	bash \
	ca-certificates

COPY . /go/src/github.com/awgh/metaloot

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		git \
		gcc \
		libc-dev \
		libgcc \
		make \
	&& cd /go/src/github.com/awgh/metaloot \
	&& CGO_ENABLED=0 go build -ldflags "-s -w" -o metaloot.bin \
	&& mv metaloot.bin /usr/bin/metaloot \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."

FROM alpine:latest

COPY --from=builder /usr/bin/metaloot /usr/bin/metaloot
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

ENTRYPOINT [ "metaloot" ]
