help:
	go run ./cmd/web -help

run:
	go run ./cmd/web -addr=":4000"

mysql:
	mysql -D snippethub -u web -p

verify:
	go mod verify

download:
	go mod download

tidy:
	go mod tidy

test-create:
	curl -iL -X POST http://localhost:4000/snippet/create

test:
	go test -v ./...

test-cover:
	go test -covermode=count -coverprofile=/tmp/profile.out ./...

test-report:
	go tool cover -html=/tmp/profile.out

PHONY:
	run help mysql verify download tidy test-create test test-cover test-report