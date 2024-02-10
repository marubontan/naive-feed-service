run:
	go run app/cmd/main.go
test:
	go test ./...

gen-swagger:
	swag init -g ./app/server/server.go -o ./app/cmd/docs