package main

import (
    "log"

    _ "go-rest-api/docs"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
    router := gin.Default()
    router.Use(cors.Default())
    router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    err := router.Run("localhost:8080")
    if err != nil {
        log.Println("Error running router: ", err)
    }

}