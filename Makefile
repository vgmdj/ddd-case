AUTOMATIONPATH := $(shell pwd)

run: checkEnv
	@export AUTOMATIONPATH=$(AUTOMATIONPATH)
	@go run main.go

install: clean
	@go build -mod=vendor -o automation main.go

linux:
	@export CGO_ENABLE=0
	@export GOOS=linux
	@export GOARCH=amd64
	@go build -mod=vendor -o automation main.go

clean:
	@go clean

checkEnv:
ifeq ($(env), )
	@echo "please set env first"
	@echo "eg : make install env=dev"
	@exit 1
endif


.PHONY: checkEnv run
.ONESHELL: checkEnv run
