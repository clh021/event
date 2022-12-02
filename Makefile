
.PHONY: clean
clean:
	rm -rf dist/webdist

.PHONY: bin
bin:
	@CGO_ENABLED=0 go build -mod vendor -ldflags "-s -w -X main.build=${gitTime}.${gitCID}" -o bin/event cmd/*.go
	@echo "[OK] bin binary was created!"
