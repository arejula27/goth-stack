build-app:
	@go build -o bin/app ./cmd/app/

css: 
	@npm run css

run: build-app
	@./bin/app



clean: 
	@rm -rf bin