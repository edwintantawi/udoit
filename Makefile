COVERAGE_FILE=coverage.out
BIN_DIR=bin

all: build

prepare:
	@echo "Preparing ..."
	@go install github.com/onsi/ginkgo/v2/ginkgo@latest
	@go mod download

build:
	@echo "Building ..."
	@go build -o ${BIN_DIR}/udoit cmd/main.go

run:
	@go run cmd/main.go

test:
	@echo "Running tests ..."
	@ginkgo -r --cover --coverprofile=${COVERAGE_FILE}

cover: test
	@echo "Running tests coverage ..."
	@go tool cover -html=${COVERAGE_FILE}

clean:
	@echo "Cleaning ..."
	@rm -rf ${BIN_DIR}
	@rm -rf ${COVERAGE_FILE}