package routers

import (
    "github.com/gin-gonic/gin"
    
    "personnelMS-server/app/Controllers/Api"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    api := r.Group("/api")
    {
        api.GET("/index", Api.Index)
    }

    return r
}