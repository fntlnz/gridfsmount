.PHONY: clean build

dist/gridfsmount:
	mkdir -p dist/
	cd dist/; \
		go build -o gridfsmount ..

build: dist/gridfsmount

clean:
	rm -Rf dist/
