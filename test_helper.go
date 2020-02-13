package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func emptyDB() {
  psqlInfo := fmt.Sprintf("host=%s port=%d  "+
    " dbname=%s sslmode=disable",
    host, port, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  sqlStatement_1 := `TRUNCATE users CASCADE`
  sqlStatement_2 := `TRUNCATE posts CASCADE`
  sqlStatement_3 := `TRUNCATE likes CASCADE`

  db.Query(sqlStatement_1)
  db.Query(sqlStatement_2)
  db.Query(sqlStatement_3)
}

func populateDB() {
  psqlInfo := fmt.Sprintf("host=%s port=%d  "+
    " dbname=%s sslmode=disable",
    host, port, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  sqlStatement_1 := `INSERT INTO users (user_id, first_name, last_name, email, password) VALUES (1, 'name_test', 'surname_test', 'test@test.com', '123')`
  sqlStatement_2 := `INSERT INTO posts (post_id, created_by, content) VALUES (10, 1, 'Test post')`
  sqlStatement_3 := `INSERT INTO likes (like_id, created_by, post_liked) VALUES (11, 1, 10)`

  db.Exec(sqlStatement_1)
  db.Exec(sqlStatement_2)
  db.Exec(sqlStatement_3)
}
