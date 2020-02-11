package main

import	"github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)

	router.GET("/create", showPostCreatePage)

	router.POST("/post/create", createPost)

	return router
}

func main() {

	router := setupRouter()
	router.Run()
}
