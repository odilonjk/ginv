build:
	@echo 'Building a ginv binary file...'
	@go build -o ginv cmd/cli.go
	@echo 'Done!'