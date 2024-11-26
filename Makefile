.PHONY: default
default: build

all: clean get-deps build test

version := "0.1.0"

build:
	mkdir -p bin
	go build -o bin/service-sonar cmd/api/main.go

test: build
	go test -cover ./... -coverprofile="bin/cover.out"
	go tool cover -func="bin/cover.out"

clean:
	rm -rf ./bin

sonar: test
	sonar-scanner -Dsonar.projectVersion="$(version)"

start-sonar:
	docker run --name sonarqube -p 9000:9000 sonarqube