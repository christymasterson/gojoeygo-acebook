package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


	type person struct {
		UserCount int
		User_id   int
		Email     string
		Password  string
	}

func showSignUpPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"signup.html",
		gin.H{
			"title": "Sign Up",
		},
	)
}
func showLogInPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Log In",
		},
	)
}

func signUp(c *gin.Context) {
	firstname := c.PostForm("first_name")
	lastname := c.PostForm("last_name")
	mail := c.PostForm("email")
	pass := c.PostForm("password")

	user := registerNewUser(firstname, lastname, mail, pass)

	user_id := user.User_id

	token := strconv.Itoa(user_id)

	c.SetCookie("token", token, 3600, "", "", false, true)
	c.Set("is_logged_in", true)

	c.Redirect(
		303,
		"/",
	)
}

func registerNewUser(firstname string, lastname string, mail string, pass string) person {

	psqlInfo := fmt.Sprintf("host=%s port=%d  "+
		" dbname=%s sslmode=disable",
		host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStatement := `
  INSERT INTO users (first_name, last_name, email, password)
  VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, firstname, lastname, mail, pass)
	if err != nil {
		panic(err)
	}

	row := db.QueryRow("SELECT COUNT(user_id) AS userCount, user_id, email, password FROM users WHERE email=$1 GROUP BY user_id", mail)

	var user person

	row.Scan(&user.UserCount, &user.User_id, &user.Email, &user.Password)

	return user
}

func logIn(c *gin.Context) {

	remail := strings.TrimSpace(c.PostForm("email"))
	rpassword := strings.TrimSpace(c.PostForm("password"))

	user_id, password := validateUser(remail, rpassword)

	if password == rpassword {

		token := strconv.Itoa(user_id)

		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		c.Redirect(
			303,
			"/",
		)

	} else {
		c.Redirect(
			303,
			"/signup",
		)
	}

}

func validateUser(remail, rpassword string) (int, string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d  "+
		" dbname=%s sslmode=disable",
		host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var (
		userCount int
		user_id   int
		email     string
		password  string
	)

	db.QueryRow("SELECT COUNT(user_id) AS userCount, user_id, email, password FROM users WHERE email=$1 GROUP BY user_id", remail).Scan(&userCount, &user_id, &email, &password)

	return user_id, password
}

func logOut(c *gin.Context) {
	// Clear the cookie

	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(
		303,
		"/login",
	)
}
