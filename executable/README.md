sudo apt-get update && sudo apt-get -y install golang-go 
ubuntu commands
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
    GOOS=windows GOARCH=amd64 go build
    GOOS=darwin GOARCH=amd64 go build