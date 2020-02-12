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

	return router
}

func main() {

	router := setupRouter()
	router.Run()
}

func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.Redirect(
				303,
				"/login",
			)
		}
	}

}

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.Redirect(
				303,
				"/",
			)
		}
	}
}

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
