.PHONY: clean build

dist/gridfsmount:
	mkdir -p dist/
	cd dist/; \
		go build -o gridfsmount ..

build: dist/gridfsmount

test:
	go list ./... | grep -v /vendor  | xargs go test

clean:
	rm -Rf dist/
