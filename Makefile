
.PHONY: clean
clean:
	rm -rf dist/webdist

gitTime=$(shell date +%Y%m%d%H%M%S)
gitCID=$(shell git rev-parse HEAD)
.PHONY: bin
bin:
	@CGO_ENABLED=0 go build -mod vendor -ldflags "-s -w -X main.build=${gitTime}.${gitCID}" -o bin/event cmd/*.go
	@echo "[OK] bin binary was created!"
