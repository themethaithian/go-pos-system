run:
	go run cmd/app/main.go

migratecreate:
	migrate create -ext sql -dir database/migrations -seq ${SEQ}

migrateup:
	migrate -database 'postgres://root:password@127.0.0.1:5432/go-pos?sslmode=disable' -path database/migrations up

migratedown:
	migrate -database 'postgres://root:password@127.0.0.1:5432/go-pos?sslmode=disable' -path database/migrations down
