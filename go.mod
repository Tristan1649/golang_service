module test_go

go 1.16

require (
	github.com/go-redis/redis/v8 v8.10.0
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgx/v4 v4.14.1
	github.com/nats-io/nats-server/v2 v2.10.11 // indirect
	github.com/nats-io/nats.go v1.33.0
)

replace github.com/nats-io/nats.go => github.com/nats-io/nats.go v1.11.1-0.20210825181159-53cc936f74e9
