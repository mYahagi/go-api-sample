package main

import (
    "./src/app/infrastructure/database"
    "github.com/gin-gonic/gin"
    //"github.com/mYahagi/src/app/infrastructure/database"
)

func main() {
    r := gin.Default()
    db := database.GormConnect()

    defer db.Close()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
