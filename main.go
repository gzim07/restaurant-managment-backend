package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gzim07/restaurant-managment-backend/routes"
)

func main() {

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoute(router)

	routes.FoodRouter(router)
	routes.InvoiceRouter(router)
	routes.MenuRouter(router)
	routes.OrderItemRouter(router)
	routes.OrderRouter(router)
	routes.TableRouter(router)
	router.Run(":8080")
}
