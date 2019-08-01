TESTTARGETS := $(shell go list -e ./... | egrep -v "/(vendor)/")
# ex, -v
TESTOPTS :=

default: gobuild

.PHONY: vendor
vendor:
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
	go build .