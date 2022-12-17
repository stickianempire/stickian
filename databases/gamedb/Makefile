build:
	docker build -t stickianempire/gamedb:0.1 .

dev-run: build
	docker run --name dev-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d -p 5432:5432 stickianempire/gamedb:0.1
	@echo "Your DB is running with localhost:5432 and username/password: postgres/postgres"
	@echo "If you wish to exec into it use \"docker exec -it dev-postgres bash\""


dev-clean:
	docker ps -a| grep "stickianempire" | awk '{print $$1}' | xargs -d '\n' bash -c 'for arg do docker stop $$arg; docker rm $$arg; done' _
