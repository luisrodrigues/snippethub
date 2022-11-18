help:
	go run ./cmd/web -help

run:
	go run ./cmd/web -addr=":4000"

PHONY:
	run help