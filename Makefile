
## meta
NAME := gin-sample

setup:
	glide up
go:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags netgo -a -v -o main *.go
docker:
	docker build --rm=false -t gin-sample:latest .
