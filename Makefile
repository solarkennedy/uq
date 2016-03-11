run: uq
	@cat itest/example.yaml | ./uq
	@./uq itest/example.yaml

uq: *.go
	@go build .

clean:
	rm uq

deps:
	go get .

test:
	go test -v .

deb: uq
	fpm -s dir -t deb --prefix=/usr/bin/ --name=uq \
	  --deb-user=root --deb-group=root \
	  uq
