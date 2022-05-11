GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go-yang-generator-generator

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME)
test:
	$(GOTEST) ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run: build
	./$(BINARY_NAME)
localtest: build
	cp $(BINARY_NAME) ~/bin
testcoverage:
	$(GOTEST) -coverprofile coverage.out ./... && go tool cover -html=coverage.out && rm coverage.out
lint:
	staticcheck -f stylish github.com/andresterba/go-yang-generator-generator/...
