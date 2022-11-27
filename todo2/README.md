docker pull postgres
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
docker run --name=todo-rest_db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
docker exec -it  todo-rest_db /bin/bash
psql -U postgres
смена с dirty
select * from schema_migrations;
update  schema_migrations set version ='000001',dirty=false;

что бы читать с /env файла используется библиотека  go get -u github.com/joho/godotenv

