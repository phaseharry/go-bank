# runs go build and output project binary as bin/gobank
build:
	@go build -o bin/gobank

# runs build and then the command below it.
# adding @ in front comments will not have the commands to printed out when ran
run: build
	@./bin/gobank

test:
	@go test -v ./..