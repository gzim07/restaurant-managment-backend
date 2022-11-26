package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/gzim07/restaurant-managment-backend/controllers"
)

func FoodRouter(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id)", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id)", controller.UpdateFood())
}
