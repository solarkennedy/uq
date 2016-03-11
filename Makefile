run: uq
	@cat itest/example.yaml | ./uq
	@./uq itest/example.yaml

uq: *.go
	@go build .

clean:
	rm uq

test:
	go test -v .
