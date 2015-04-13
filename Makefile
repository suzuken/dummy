gen:
	go run dummy.go

test:
	go test -v

fmt:
	gofmt -w *.go
