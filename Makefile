
build:
	go mod verify
	env GOOS=linux go build -ldflags="-s -w" -o bin/http cmd/http/*

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

destroy:
	sls remove --verbose
