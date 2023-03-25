build:
	go build -o bin main.go
mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql
createdb:
	docker exec -it mysql mysql -u root -ppassword -e "CREATE DATABASE rentcar;"
dropdb:
	docker exec -it mysql mysql -u root -ppassword -e "DROP DATABASE rentcar;"
accessdb:
	docker exec -it mysql mysql -u root -ppassword rentcar
migrateup:
	migrate -path database/migration -database "mysql://root:password@tcp(localhost:3306)/rentcar" -verbose up
migratedown:
	migrate -path database/migration -database "mysql://root:password@tcp(localhost:3306)/rentcar" -verbose down
.PHONY: build mysql createdb dropdb accessdb