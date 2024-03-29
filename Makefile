.PHONY: itest clean

VERSION=$(shell git describe --abbrev=0 --tags)

uq: *.go
	go build -ldflags "-X main.version=$(VERSION)" .

itest: uq
	./uq itest/example.yaml -s yaml | grep -q '"comments": "Late afternoon is best. Backup contact is Nancy Billsmer @ 338-4338.\\n",'
	./uq itest/example.json -s json | grep -q '"parse_time_nanoseconds": 127664,'
	./uq itest/example.xml  -s xml  | grep -q '"description": "Two of our famous Belgian Waffles with plenty of real maple syrup",'
	./uq itest/example.ini  -s ini  | grep -q '"app_mode": "development"'
	./uq itest/example.ini  -s ini  | grep -q ''
	./uq itest/example.toml -s toml | grep -q '"title": "TOML Example"'
	@echo "itest pass"

clean:
	rm uq

test:
	go test -v .

deb: uq
	fpm -s dir -t deb --prefix=/usr/bin/ --name=uq \
	  --deb-user=root --deb-group=root \
	  --version=$(VERSION) \
	  uq

fmt:
	go fmt *.go
