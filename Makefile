mimic:
	go build -n main.go

go:
	go build -ldflags="-s -w" -o bin/psite main.go

docker:
	docker build -t johnodonn/psite:3 .
