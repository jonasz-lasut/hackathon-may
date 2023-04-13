ifndef $(GOPATH)
	GOPATH=$(shell go env GOPATH)
	export GOPATH
endif

build:
	go build -o dist/main main.go

documentation:
	go run main.go -docs

development:
	docker compose -f devel/docker-compose.yaml up -d

run:
	go run main.go
