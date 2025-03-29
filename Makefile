createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:rootpassword@tcp(localhost:3306)/crytohub" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:rootpassword@tcp(localhost:3306)/crytohub" -verbose down	

v:
	migrate -path=sql/migrations -database "mysql://root:rootpassword@tcp(localhost:3306)/crytohub" -verbose version

clean:
	migrate -path=sql/migrations -database "mysql://root:rootpassword@tcp(localhost:3306)/crytohub" -verbose force 2		

.PHONY: migrate createmigration migratedown v clean