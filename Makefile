ifndef $(GOPATH)
	GOPATH=$(shell go env GOPATH)
	export GOPATH
endif

build:
	go build -o main main.go

documentation:
	go run main.go -docs

development: development-teardown
	docker compose up -d

development-teardown:
	docker compose down --rmi local

run:
	go run main.go
