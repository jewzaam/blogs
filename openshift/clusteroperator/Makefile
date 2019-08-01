TESTTARGETS := $(shell go list -e ./... | egrep -v "/(vendor)/")
# ex, -v
TESTOPTS :=

GO111MODULE?=on

default: gobuild

.PHONY: gomodhack
gomodhack:
	GO111MODULE=on go get -d github.com/openshift/client-go/operator/informers/externalversions@master
	GO111MODULE=on go get -d github.com/openshift/api@master
	GO111MODULE=on go get -d github.com/openshift/api/config/v1@master
	GO111MODULE=on go get -d github.com/openshift/api/operator/v1alpha1@master
	GO111MODULE=on go get -d github.com/openshift/client-go/config/clientset/versioned@master
	GO111MODULE=on go get -d github.com/openshift/client-go/config/clientset/versioned/typed/config/v1@master
	GO111MODULE=on go get -d github.com/openshift/client-go/config/informers/externalversions@master
	GO111MODULE=on go get -d github.com/openshift/client-go/config/informers/externalversions/config/v1@master
	GO111MODULE=on go get -d github.com/openshift/client-go/config/listers/config/v1@master
	GO111MODULE=on go get -d github.com/openshift/client-go/operator/clientset/versioned@master

.PHONY: gomodtidy
gomodtidy:
	GO111MODULE=${GO111MODULE} go mod tidy

.PHONY: gocheck
gocheck: gomodtidy
	GO111MODULE=${GO111MODULE} gofmt -s -w $(shell go list -f '{{ .Dir }}' ./... )
	GO111MODULE=${GO111MODULE} go vet $(shell go list -f '{{ .Dir }}' ./... )

.PHONY: gotest
gotest:
	GO111MODULE=${GO111MODULE} go test $(TESTOPTS) $(TESTTARGETS)

.PHONY: gobuild
gobuild: gocheck gotest
	GO111MODULE=${GO111MODULE} go build $(shell go list -f '{{ .Dir }}' ./... )
