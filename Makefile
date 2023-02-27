.PHONY : build
build:
	go build -o bin/server .

.PHONY: test
test:
	go test ./... -v

.PHONY: clean
clean:
	rm bin/server