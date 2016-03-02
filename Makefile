run: uq
	@cat itest/example.yaml | ./uq

uq: *.go
	@go build .

clean:
	rm uq

test:
	go test -v .
