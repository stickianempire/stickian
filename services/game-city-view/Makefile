compile:
	CGO_ENABLED=0 go build -buildvcs=false -o ./bin/main ./cmd/

build: compile
	docker build -t stickianempire/game-city-view:0.1 .

test: compile
	go test ./...

dev-run: compile
	@echo "Will attempt to run local gamedb on localhost:5432"
	docker run --name dev-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d -p 5432:5432 stickianempire/gamedb:0.1
	@echo "Awaiting for the database to stabelize a bit..."
	sleep 1
	@echo "Running the dummy data files..."
	cat ./testdata/city0001.sql | docker exec -i dev-postgres psql -U postgres -d postgres
	@echo "Your game-city-view is running on localhost:4000"
	./bin/main

dev-clean:
	rm -r ./bin/
	docker ps -a | grep "stickianempire" | awk '{print $$1}' | xargs -d '\n' bash -c 'for arg do docker stop $$arg; docker rm $$arg; done' _

