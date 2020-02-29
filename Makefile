NAME := gipr
VERSION := $(shell cat VERSION)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := "-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

export GO111MODULE ?= on

build:
	go build -ldflags=$(LDFLAGS) -o bin/$(NAME) ./cmd/gipr

bin/$(NAME): $(SRCS)
	go build -ldflags=$(LDFLAGS) -o bin/$(NAME) ./cmd/gipr 

.PHONY: gox
gox:
	gox -ldflags=$(LDFLAGS) -output="bin/$(NAME)_{{.OS}}_{{.Arch}}"  ./cmd/gipr

.PHONY: zip
zip:
	cd bin ; \
	for file in *; do \
		zip $${file}.zip $${file} ; \
	done

.PHONY: gox_with_zip
gox_with_zip: clean gox zip

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: tag
tag:
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push --tags

generate:
	go generate ./...
