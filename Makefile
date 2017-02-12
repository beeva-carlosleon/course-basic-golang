TEST_UNIT?=$(shell go list ./... | grep -v vendor)
export PATH := $(GOPATH)/bin:$(PATH)

all: vendor build

install:
	@echo "Install godep and restoring dependencies"
	go get github.com/tools/godep

vendor:
	godep save ./...

build:
	@echo "building..."
	go build -o build/bin/helloworld

fmt:
	@echo "Formatting all go code..."
	go fmt `go list ./... | grep -v vendor`

test:
	@echo "Running the tests"
	@go test $(TEST_UNIT) -v
	@go vet $(TEST_UNIT) ; if [ $$? -eq 1 ]; then \
		echo "ERROR: Vet found problems in the code."; \
		exit 1; \
		fi

clean:
	@echo "Cleaning binaries..."
	@rm -rf $(GOPATH)/bin $(GOPATH)/pkg

.PHONY: all install vendor build fmt test clean