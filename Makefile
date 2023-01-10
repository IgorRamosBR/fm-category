build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create cmd/create-category/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list cmd/list-categories/main.go

deploy-dev:
	serverless deploy --aws-profile PERSONAL --stage dev