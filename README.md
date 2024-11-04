# Go, gin, swag, gorm, atlas

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

## Create postgres container

[https://docs.docker.com/engine/install/ubuntu/]

[https://github.com/LessonDump/DockerPostgresPgAdmin/blob/master/docker-compose.yaml]

## Install gorm

`go get -u gorm.io/gorm`

`go get -u gorm.io/driver/postgres`

[https://github.com/YugabyteDB-Samples/orm-examples/blob/master/golang/gorm/src/controller/base.go]

[https://gist.github.com/mashingan/4212d447f857cfdfbbba4f5436b779ac]

## Install atlas

[https://atlasgo.io/guides/orms/gorm/standalone]

`go get -u ariga.io/atlas-provider-gorm`

`go get -u ariga.io/atlas-provider-gorm/gormschema`

`atlas schema inspect --env gorm --url "env://src"`

`atlas migrate diff --env gorm`

`atlas migrate apply --url "postgres://postgres:postgres@:5432/postgres?sslmode=disable"`
