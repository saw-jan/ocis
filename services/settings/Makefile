SHELL := bash
NAME := settings

include ../../.make/recursion.mk

############ tooling ############
ifneq (, $(shell command -v go 2> /dev/null)) # suppress `command not found warnings` for non go targets in CI
include ../../.bingo/Variables.mk
endif

############ go tooling ############
include ../../.make/go.mk


############ release ############
include ../../.make/release.mk

############ docs generate ############
include ../../.make/docs.mk

.PHONY: docs-generate
docs-generate: config-docs-generate \
			   grpc-docs-generate

############ generate ############
include ../../.make/generate.mk

.PHONY: ci-go-generate
ci-go-generate: protobuf # CI runs ci-node-generate automatically before this target

.PHONY: ci-node-generate
ci-node-generate:

############ protobuf ############
include ../../.make/protobuf.mk

.PHONY: protobuf
protobuf: buf-generate $(PROTOC_GO_INJECT_TAG)
	cd ../../protogen/gen/ && $(PROTOC_GO_INJECT_TAG) -input="ocis/messages/settings/v0/*.pb.go"

############ licenses ############
.PHONY: ci-node-check-licenses
ci-node-check-licenses:

.PHONY: ci-node-save-licenses
ci-node-save-licenses:
