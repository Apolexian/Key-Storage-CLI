build:
	~/go/bin/golint cobra/*.go
	gofmt main.go
	go build -o ../bin/main main.go
	gofmt cobra/*.go
	godoc -html ../cmd/cobra/ > cobra/docs/cmd_docs.html
