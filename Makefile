linux-build:
	GOOS=linux GOARCH=amd64 go build -o installer main.go