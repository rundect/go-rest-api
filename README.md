# Go, gin, swag, gorm

`go mod init go-rest-api`

`go mod tidy`

## Install gin

`go get -u github.com/gin-gonic/gin`

`go get -u github.com/gin-contrib/cors`

## Install swag

`go get -u github.com/swaggo/gin-swagger`

`go get -u github.com/swaggo/swag`

`go get -u github.com/swaggo/files`

## Need for `swag init`

`echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc`

`source ~/.bashrc`

`swag init`

## Install other packages

`go get -u github.com/gofrs/uuid`

## Create postgres container

[https://docs.docker.com/engine/install/ubuntu/]

[https://github.com/LessonDump/DockerPostgresPgAdmin/blob/master/docker-compose.yaml]

## Install gorm

`go get -u gorm.io/gorm`

`go get -u gorm.io/driver/postgres`

[https://github.com/YugabyteDB-Samples/orm-examples/blob/master/golang/gorm/src/controller/base.go]

[https://gist.github.com/mashingan/4212d447f857cfdfbbba4f5436b779ac]

## Install golang-migrate

[https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md]

`go get github.com/golang-migrate/migrate/v4`

`export POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable'`

[https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md]

[https://www.geeksforgeeks.org/how-to-install-golang-migrate-on-ubuntu]

`migrate create -ext sql -dir migrations -seq create_users_table`
