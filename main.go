package main

import	"github.com/gin-gonic/gin"

const (
	host   = "localhost"
	port   = 5432
	dbname = "acebook"
)

func setupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)

	router.GET("/create", showPostCreatePage)

	router.POST("/post/create", createPost)

	router.GET("/login", showLogInPage)

	router.POST("/login", logIn)

	router.GET("/signup", showSignUpPage)

	router.POST("/signup", signUp)

	return router
}

func main() {

	router := setupRouter()
	router.Run()
}
