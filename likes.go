package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func createLike(c *gin.Context) {
  post := c.PostForm("post_id")
	post_id, _ := strconv.Atoi(post)

	token, _ := c.Cookie("token")
	cookie, _ := strconv.Atoi(token)
	err := insertDBLike(cookie, post_id)
	if err != nil {
		fmt.Println(err)
	}

	c.Redirect(
		303,
		"/",
	)
}

func insertDBLike(cookie int, post int) error {

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
	return err

}
