package main

import (
	"github.com/gin-gonic/gin"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "acebook"
)

func setupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Use(setUserStatus())

	router.GET("/", ensureLoggedIn(), showIndexPage)

	router.GET("/create", ensureLoggedIn(), showPostCreatePage)

	router.POST("/post/create", ensureLoggedIn(), createPost)

	router.GET("/login", ensureNotLoggedIn(), showLogInPage)

	router.GET("/logout", ensureLoggedIn(), logOut)

	router.POST("/login", ensureNotLoggedIn(), logIn)

	router.GET("/signup", ensureNotLoggedIn(), showSignUpPage)

	router.POST("/signup", ensureNotLoggedIn(), signUp)

	router.POST("/like/create", ensureLoggedIn(), createLike)

	return router
}

func main() {

	router := setupRouter()
	router.Run()
}
