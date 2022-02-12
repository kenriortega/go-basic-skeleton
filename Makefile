git_hash := $(shell git rev-parse --short HEAD)

app:
	SOME_TEXT="Hello world" go run cmd/main.go

# Use the dockerfile principal
docker-build:
	docker build -t username/basic-go:${git_hash} \
		-f Dockerfile .
docker-app:
	docker run --name app --rm -m 128m --cpus="0.25" \
	 -e SOME_TEXT="Hello world" username/basic-go:${git_hash}