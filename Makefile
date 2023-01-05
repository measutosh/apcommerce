build :
	@go build -o bin/apcommerce
run : build	
	@./bin/apcommerce
test: 
	@go test -v ./...