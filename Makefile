.PHONY: itest clean
uq: *.go
	@go build .

itest: uq
	./uq itest/example.yaml -s yaml | grep -q '"comments": "Late afternoon is best. Backup contact is Nancy Billsmer @ 338-4338.\\n",'
	./uq itest/example.json -s json | grep -q '"parse_time_nanoseconds": 127664,'
	./uq itest/example.xml  -s xml  | grep -q '"description": "Two of our famous Belgian Waffles with plenty of real maple syrup",'
	./uq itest/example.toml -s toml | grep -q '"title": "TOML Example"'
	@echo "itest pass"

clean:
	rm uq

deps:
	dep ensure

test:
	go test -v .

deb: uq
	fpm -s dir -t deb --prefix=/usr/bin/ --name=uq \
	  --deb-user=root --deb-group=root \
	  uq
