package main

import (

  "testing"

	_ "github.com/lib/pq"
)

func TestGetAllPosts(t *testing.T) {

  emptyDB()

  populateDB()

  actual := getAllPosts()
  expect := 1

  if len(actual) != expect {
			t.Fail()
  }
}
