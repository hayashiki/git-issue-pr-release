FROM golang:1.13.6-alpine AS build

RUN apk --update add --no-cache alpine-sdk ca-certificates \
 && update-ca-certificates

WORKDIR /work
# not use go modules
# WORKDIR /Users/hayashida/src/github.com/hayashiki/git-issue-pr-release
# COPY ./ ${GOPATH}/src/github.com/hayashiki/

ENV CGO_ENABLED 0
# ENV GOOS linux

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

# FROM scratch

# WORKDIR /work

FROM alpine:3.11

RUN apk update \
	&& apk upgrade \
	&& apk add \
	bash \
	rm -rf /var/cache/apk/*

WORKDIR /work

COPY --from=build  /work/bin/giprdraft /usr/local/bin/giprdraft
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY *.sh /

RUN chmod +x /*.sh

# EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
