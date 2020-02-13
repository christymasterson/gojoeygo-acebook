package main

import (

  "testing"
  _ "github.com/lib/pq"
  )

func TestCreateLikes(t *testing.T) {
  emptyDB()

  populateDBLikes()

  post_id := 10
  cookie := 1

  actual := insertDBLike(cookie, post_id)

  if actual != nil {
    t.Fail()
  }

}
