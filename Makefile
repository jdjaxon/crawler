BIN := crawler


.PHONY: build
build:
	go build -o $(BIN)


.PHONY: test
test:
	go test ./... -v


.PHONY: clean
clean:
	$(RM) $(BIN)
