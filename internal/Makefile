build:
	~/go/bin/golint encryption/encryption.go
	gofmt -d -l encryption/encryption.go
	go build -o ../bin/encryption encryption/encryption.go
	~/go/bin/golint storage/storage.go
	gofmt -d -l storage/storage.go
	go build -o ../bin/storage storage/storage.go
	godoc -html ../internal/encryption/ > encryption_doc.html
	godoc -html ../internal/storage/ > storage_doc.html

