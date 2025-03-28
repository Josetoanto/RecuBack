package main

import (
    "recuAPI/handlers"
    "recuAPI/infra"
    "recuAPI/repository"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.Use(infra.CORSMiddleware())

    repo := repository.NewProductoRepository()
    handler := handlers.NewProductoHandler(repo)

    r.POST("/addProduct", handler.AddProduct)
    r.GET("/getTemporaryProducts", handler.GetTemporaryProducts)
    r.GET("/countProductInDiscount", handler.CountProductInDiscount)

    r.Run(":8080")
}
