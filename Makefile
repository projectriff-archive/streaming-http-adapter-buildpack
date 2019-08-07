GO_SOURCES = $(shell find . -type f -name '*.go')

.PHONY: all
all: test build

.PHONY: build
build: artifactory/io/projectriff/streaming-http-adapter/io.projectriff.streaming-http-adapter

.PHONY: test
test:
	go test -v ./...

artifactory/io/projectriff/streaming-http-adapter/io.projectriff.streaming-http-adapter: buildpack.toml $(GO_SOURCES)
	rm -fR $@                           && \
	./ci/package.sh                     && \
	mkdir $@/latest                     && \
	tar -C $@/latest -xzf $@/*/*.tgz

.PHONY: clean
clean:
	rm -fR artifactory/
	rm -fR dependency-cache/
	rm -fR bin/