test: build
	go test -v ./...

build:
	go build

run: build
	./ll3

doc:
	godoc -http=:6060
