build-local:
	docker-compose build

build-prod:
	docker buildx build --platform linux/amd64 -t go-vetryx -f build/docker/prod/Dockerfile --load .

run-local:
	docker-compose up -d

run-prod:
	docker run -p 8080:8080 go-vetryx

test: |
	go test -v ./... -covermode=count -coverprofile=coverage.out && go tool cover -func=coverage.out -o=coverage.out

html-cov: 
	go test -v ./... -covermode=count -coverprofile=coverage.out && go tool cover -func=coverage.out && go tool cover -html=coverage.out

run-txt:
	go run cmd/filerunner/main.go test.txt