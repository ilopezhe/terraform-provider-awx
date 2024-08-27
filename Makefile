HOSTNAME ?= github.com
NAMESPACE ?= ilopezhe
NAME ?= awx
BINARY ?= ./build/terraform-provider-${NAME}
VERSION ?= 24.6.1
OS_ARCH ?= darwin_arm64 # Can be overridden to linux_amd64, darwin_amd64, etc.

default: generate build

.PHONY: generate-config
generate-config:
	node tools/config-merge.js $(shell pwd)/resources/config $(shell pwd)/resources/api/$(VERSION)

.PHONY: download-api
download-api:
	mkdir -p resources/api/$(VERSION)/config resources/api/$(VERSION)/gen-data
	go run ./tools/generator/cmd/generator/main.go fetch-api-resources resources/api/$(VERSION) --host http://localhost:8080 --password changeme --username test

.PHONY: generate-configs
generate-configs: resources/api/*
	@for file in $^ ; do \
		node tools/config-merge.js $(shell pwd)/resources/config $(shell pwd)/$${file} ; \
	done

.PHONY: generate-awx
generate-awx: generate-config
	rm -f internal/awx/gen_*.go
	rm -rf cmd/provider/docs/*
	go run ./tools/generator/cmd/generator/main.go template resources/api/$(VERSION) internal/awx
	goimports -w internal/awx/*.go
	gofmt -s -w internal/awx/*.go

.PHONY: generate-tfplugindocs
generate-tfplugindocs:
	rm -rf docs
	mkdir -p cmd/provider/docs
	tfplugindocs generate --examples-dir examples --provider-name awx --provider-dir ./cmd/provider
	mv cmd/provider/docs .

.PHONY: generate
generate: generate-awx generate-tfplugindocs

.PHONY: build-cover
build-cover:
	go build -cover -trimpath -o ./build/terraform-provider-awx ./cmd/provider

.PHONY: build
build:
	go build -trimpath -o ${BINARY} -ldflags "-s -w" ./cmd/provider

.PHONY: test
test:
	go test ./internal/... -count=1 -parallel=4 -cover -coverprofile=build/coverage.out
	go tool cover -html=build/coverage.out -o build/coverage.html

.PHONY: testacc
testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -tags=full -v ./...

.PHONY: testfull
testfull:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...; \
	rm -rf ./examples/full/.terraform* ./examples/full/terraform.tfstate* && \
	rm -rf ./examples/full/.terraform* ./examples/full/terraform.tfstate*

.PHONY: install
install: build
	go install ./cmd/provider/
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

.PHONY: dev
dev:
	cd ./tools/mage && go run mage.go -v reCreate && cd ../..

.PHONY: dev-cleanup
dev-cleanup:
	cd ./tools/mage && go run mage.go -v delete && cd ../..

.PHONY: port-forward
port-forward:
	kubectl port-forward service/awx-service -n ansible-awx 8080:80 > /dev/null 2>&1 &

.PHONY: no-port-forward
no-port-forward:
	pkill -f "kubectl port-forward"

