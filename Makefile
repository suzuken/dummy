gen:
	go run cmd/dummy/dummy.go

test:
	go test -v

fmt:
	gofmt -w *.go

bench:
	go test -bench .

build:
	go build cmd/dummy/*.go

ci:
	make fmt
	git add -u *
	git commit

csv:
	go run cmd/dummy/dummy.go -f "str|10,str|10"
