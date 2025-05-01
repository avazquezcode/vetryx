.PHONY: build
build:
	go build -o build/filerunner cmd/filerunner/main.go

test: |
	go test -v ./... -covermode=count -coverprofile=coverage.out && go tool cover -func=coverage.out -o=coverage.out

html-cov: 
	go test -v ./... -covermode=count -coverprofile=coverage.out && go tool cover -func=coverage.out && go tool cover -html=coverage.out

run-txt:
	go run cmd/filerunner/main.go test.txt

.PHONY: build-wasm
build-wasm:
	rm -rf build/wasm
	mkdir -p build/wasm
	cp -a web/ build/wasm
	GOOS=js GOARCH=wasm go build -o build/wasm/main.wasm cmd/wasm/main.go
	cp "$(shell go env GOROOT)/lib/wasm/wasm_exec.js" build/wasm/

.PHONY: serve-wasm
serve-wasm: build-wasm stop-wasm 
	cd build/wasm && python3 -m http.server 8080

.PHONY: stop-wasm
stop-wasm:
	@lsof -ti:8080 | xargs kill -9 2>/dev/null || true
