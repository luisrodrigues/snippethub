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

PHONY:
	run help mysql verify download tidy