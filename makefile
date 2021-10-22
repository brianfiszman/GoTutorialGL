db-up:	
	migrate -path db/migrations -database "postgresql://root:123123@localhost:5432/testdb?sslmode=disable" -verbose up

db-down:
	migrate -path db/migrations -database "postgresql://root:123123@localhost:5432/testdb?sslmode=disable" -verbose down