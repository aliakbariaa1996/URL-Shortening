build:
ifeq ($(OS),Windows_NT)
	go build -o build/auth-profile-backend.exe main.go
else
	go build -o build/auth-profile-backend main.go
endif

install:
ifeq ($(OS),Windows_NT)
	go install main.go
else
	go install main.go
endif

.PHONY: build install

test:
	go test -v ./...

swag:
	swag init -g main.go