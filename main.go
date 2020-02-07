package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "acebook"
)

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.GET("/create", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"create-post.html",
			gin.H{
				"title": "Create New Post",
			},
		)
	})

	router.POST("/post/create", func(c *gin.Context){
		post := c.PostForm("content")

		psqlInfo := fmt.Sprintf("host=%s port=%d  "+
			" dbname=%s sslmode=disable",
			host, port, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		sqlStatement := `
		INSERT INTO posts (content)
		VALUES ($1)`
		_, err = db.Exec(sqlStatement, post)
		if err != nil {
			panic(err)
		}

			c.HTML(
				http.StatusOK,
				"index.html",
				gin.H{
					"title": "Home Page",
				},
			)
	})

	return router
}

func main() {

	router := setupRouter()
	router.Run(":8080")

}
