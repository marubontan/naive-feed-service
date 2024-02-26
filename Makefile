run:
	go run app/cmd/main.go
test:
	go test -count=1 ./...
test-e2e:
	go test -tags=integration -count=1 ./...

gen-swagger:
	swag init -g ./app/server/server.go -o ./app/cmd/docs