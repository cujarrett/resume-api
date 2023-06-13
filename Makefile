clean:
	go clean

install:
	go install -v -t -d ./...

test:
	go test -v ./...

compile:
	GOOS=linux go build -o resume-api main.go
