TESTTARGETS := $(shell go list -e ./... | egrep -v "/(vendor)/")
# ex, -v
TESTOPTS :=

default: gobuild

.PHONY: vendor
vendor:
	dep init
	dep ensure -v

.PHONY: gocheck
gocheck:
    gofmt -s -l $(shell go list -f '{{ .Dir }}' ./... ) | grep ".*\.go"; if [ "$$?" = "0" ]; then gofmt -s -d $(shell go list -f '{{ .Dir }}' ./... ); exit 1; fi
    go vet ./cmd/... ./pkg/...

.PHONY: gotest
gotest:
    go test $(TESTOPTS) $(TESTTARGETS)

.PHONY: gobuild
gobuild: gocheck gotest
	go build .