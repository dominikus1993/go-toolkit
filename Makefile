build:
	go build -v ./...

test:
	go test ./...

vet: 
	go vet ./...
	
buildandtest: build test

upgrade:
	go get -u