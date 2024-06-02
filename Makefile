build:
	go build -o bin/uwuify
doc:
	./scripts/generate_docs.sh
run:
	go run main.go

