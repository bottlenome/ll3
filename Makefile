run:
	go build; ./ll3
doc:
	godoc -http=:6060
test:
	go test -v ./...
