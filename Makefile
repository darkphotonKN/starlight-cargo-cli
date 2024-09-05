.PHONY: test 
	
build:
	 @go build -o bin/starlight-cargo-cli ./cmd/app/
	
run: build
	@./bin/starlight-cargo-cli

test:
	@go test ./...

