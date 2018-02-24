package router

import(
	"github.com/gin-gonic/gin"
	"github.com/cancelv5/couple-server/api/controllers"
	"fmt"
)

func InitRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware)


	router.GET("/", IndexController.Index)


	return router
}

func middleware(c *gin.Context) {
	fmt.Println("middleware is working")
}