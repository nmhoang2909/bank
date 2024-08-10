mysqlup:
	podman run --name mysql -e MYSQL_ROOT_PASSWORD=secret -d -p 3306:3306 -v mysql-data:/var/lib/mysql mysql
createdb:
	podman exec -it mysql mysql -u root -p -e 'create database bank'
dropdb:
	podman exec -it mysql mysql -u root -p -e 'drop database bank'
migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp/bank" -verbose up
migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp/bank" -verbose down
test:
	go test -v -cover ./...

.PHONY: mysqlup createdb dropdb migrateup migratedown test
