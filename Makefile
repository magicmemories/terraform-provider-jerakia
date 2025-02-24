VERSION=0.2.0
OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)
HOSTNAME=registry.terraform.io
NAMESPACE=magicmemories
NAME=jerakia

ifneq (,$(findstring windows,$(OS_ARCH)))
  BINEXT=.exe
  PLUGIN_PATH=$(APPDATA)/terraform.d/plugins
else
  BINEXT=
  PLUGIN_PATH=~/.terraform.d/plugins
endif

TEST?=$$(go list ./... | grep -v 'vendor')

BINARY=terraform-provider-${NAME}${BINEXT}

default: install

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ${PLUGIN_PATH}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ${PLUGIN_PATH}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	docker-compose up -d; \
	export JERAKIA_TOKEN=$$(docker-compose exec -T jerakia jerakia token create terraform --quiet) ; \
	export JERAKIA_URL="http://localhost:19843/v1" ; \
	TF_LOG=DEBUG TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m; \
	status=$$?; \
	docker-compose down; \
	exit $$status

fmtcheck:
	echo "==> Checking that code complies with gofmt requirements..."
	files=$$(go list ./... | grep -v /vendor/ ) ; \
	gofmt_files=`gofmt -l $$files`; \
	if [ -n "$$gofmt_files" ]; then \
		echo 'gofmt needs running on the following files:'; \
		echo "$$gofmt_files"; \
		echo 'You can use the command: \`go fmt $$(go list ./... | grep -v /vendor/)\` to reformat code.'; \
		exit 1; \
	fi

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi