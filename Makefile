.PHONY: clean build

build: clean dist/gridfsmount

dist/gridfsmount:
	mkdir -p dist/
	cd dist/; \
		go build -o gridfsmount ..

test:
	go list ./... | grep -v /vendor  | xargs go test

clean:
	rm -Rf dist/
