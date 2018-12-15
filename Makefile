test: build
	go test -v -race -covermode=atomic -coverprofile=coverage.txt ./...

build:
	go build

run: build
	./ll3

doc:
	godoc -http=:6060
