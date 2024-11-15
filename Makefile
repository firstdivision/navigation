

update:
	go get -t -u ./...
	go mod tidy
	go mod download