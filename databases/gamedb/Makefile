build:
	docker build -t stickianempire/gamedb:0.1 .

dev-run: build
	docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d -p 5433:5432 stickianempire/gamedb:0.1
	@echo "Your DB is running with localhost:5432 and username/password: postgres/postgres"
	@echo "If you wish to exec into it use \"docker exec -it <contained-id> bash\""


dev-clean:
	docker ps | grep "stickianempire" | awk '{system("docker kill "$$1)}'
