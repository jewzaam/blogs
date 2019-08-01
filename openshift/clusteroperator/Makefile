TESTTARGETS := $(shell go list -e ./... | egrep -v "/(vendor)/")
# ex, -v
TESTOPTS :=

GO_VERSION:=1.11.12

default: gobuild

Gopkg.lock:
	rm -f Gopkg.toml
	dep init

.PHONY: vendor
vendor: Gopkg.lock
	dep ensure -v

.PHONY: gocheck
gocheck:
	gofmt -s -w $(shell go list -f '{{ .Dir }}' ./... )
	go vet ./pkg/...

.PHONY: gotest
gotest:
	go test $(TESTOPTS) $(TESTTARGETS)

.PHONY: gobuild
gobuild: gocheck gotest
	docker run --rm -v $$GOPATH/src/:/go/src/ -v `pwd -P`:`pwd -P` golang:${GO_VERSION} /bin/sh -c "cd `pwd`; go build ./pkg/..."

.PHONY: setup
setup:
	sudo chcon -Rt svirt_sandbox_file_t .