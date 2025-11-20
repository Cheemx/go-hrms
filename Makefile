# Up Migrations in mysql with goose
migrationUp:
	cd model/schema && goose mysql "mysql:mysql@tcp(localhost:3307)/hrms?parseTime=true" up 

# Down migrations in mysql with goose
migrationDown:
	cd model/schema && goose mysql "mysql:mysql@tcp(localhost:3307)/hrms?parseTime=true" down

# To connect to hrms DB in CLI
dbDikha:
	docker exec -it go-hrms-mysql-1 mysql -u mysql -pmysql -D hrms