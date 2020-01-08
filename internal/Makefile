build:
	~/go/bin/golint encryption/encryption.go
	gofmt -d -l encryption/encryption.go
	go build -o ../bin/encryption encryption/encryption.go
	~/go/bin/golint vault/storage.go
	gofmt -d -l vault/storage.go
	go build -o ../bin/storage vault/storage.go
