run: 
	@go build -o ./bin/api ./main.go
	@./bin/api

update:
	go mod tidy
	go mod vendor