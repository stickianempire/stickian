compile:
	CGO_ENABLED=0 go build -buildvcs=false -o ./bin/main .

dev-run: compile
	./bin/main
	@echo "Your game is running on localhost:4000"