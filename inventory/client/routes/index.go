package routes

import (
	"inventory/client/controller"

	"github.com/gin-gonic/gin"
)

func AppRoutes(r *gin.Engine) {
	r.GET("/getitems", controller.HandlerGetAll)
	r.POST("/updateitems", controller.HandlerUpdateItems)
	r.POST("/getitem", controller.HandlerGetItem)
	r.POST("/create", controller.HandlerCreate)
	r.POST("/additems", controller.HandlerAddItems)

}
