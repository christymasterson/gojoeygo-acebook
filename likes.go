package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func createLike(c *gin.Context) {
  post := c.PostForm(".ID")

	cookie, _ := c.Cookie("token")

	psqlInfo := fmt.Sprintf("host=%s port=%d  "+
		" dbname=%s sslmode=disable",
		host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
  INSERT INTO likes (created_by, post_liked)
  VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, cookie, post)
	if err != nil {
		panic(err)
	}

	c.Redirect(
		303,
		"/",
	)
}
