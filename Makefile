.PHONY: itest clean

itest: uq
	cat itest/example.yaml | ./uq >/dev/null
	./uq itest/example.yaml >/dev/null
	cat itest/example.json | ./uq >/dev/null
	./uq itest/example.json >/dev/null
	cat itest/example.xml | ./uq -s xml >/dev/null
	./uq itest/example.xml >/dev/null

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
