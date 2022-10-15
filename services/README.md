# Services

The services are the core back end of the game. They are devided between micro services directed to the game, and micro services directed to the user experience. The following prefix MUST be utilized when creating a micro service:

- `ga-` when related with game logic
- `fo-` when related with a forum logic
- `*` anything when it is not related with forum or game

Each service folder should be build using the following micro service struct:

- `main.go` where the service is launched and http server is started
- `handler.go` where the endpoint handlers do request based logic
- `service.go` where the service logic is held, using local interfaces for external requests
- (optional) `db.go` where the interface implementation is present for DB accesses, if appliable
- `*.go` within reasonable separation of concerns (interfaces vs implementations)

Every file should also have a unit test counterpart `*_test.go` that tests essential logic in the service. Note that the purpose is to have MICROSERVICES, so the number of files should be reduced, and they should be only responsible for a single topic. Common base logic should, and will, live on the `pkg` level at root.

Work in progress... :warning: :construction_worker:
