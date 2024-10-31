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

## Install gorm

`go get -u gorm.io/gorm`

`go get -u gorm.io/driver/postgres`
