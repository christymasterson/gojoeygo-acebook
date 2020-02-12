package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type post struct {
	ID      int
	Name    string
	Surname string
	Content string
	Date    time.Time
}

func showIndexPage(c *gin.Context) {

	psqlInfo := fmt.Sprintf("host=%s port=%d  "+
		" dbname=%s sslmode=disable",
		host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT posts.post_id, users.first_name, users.last_name, posts.content, posts.created_at FROM posts JOIN users ON posts.created_by = users.user_id ORDER BY created_at DESC`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	posts := make([]post, 0)

	for rows.Next() {
		var entry post
		rows.Scan(&entry.ID, &entry.Name, &entry.Surname, &entry.Content, &entry.Date)
		posts = append(posts, entry)

	}

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": posts,
		},
	)
}

func showPostCreatePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"create-post.html",
		gin.H{
			"title": "Create New Post",
		},
	)
}

func createPost(c *gin.Context) {
	post := c.PostForm("content")

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
  INSERT INTO posts (content, created_by)
  VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, post, cookie)
	if err != nil {
		panic(err)
	}

	c.Redirect(
		303,
		"/",
	)
}
