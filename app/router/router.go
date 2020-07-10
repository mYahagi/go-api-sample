package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	fishG := r.Group("fish")
	{
		fishG.GET("", func(c *gin.Context) { controller.Catch(c) })
	}

	r.Run()
}
