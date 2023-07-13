build-app:
	@go build -o bin/app ./cmd/app/

css: 
	@npm run css

run: build-app
	@./bin/app

dev: css
	air


clean: 
	@rm -rf bin