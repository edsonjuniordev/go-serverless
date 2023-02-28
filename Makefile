clean:
	go clean
	rm -rf ./bin

build: clean
	go mod tidy
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello-world main.go

docker-compose:
	sudo docker compose up -d

dev: docker-compose
	sls offline --useDocker start --host 0.0.0.0