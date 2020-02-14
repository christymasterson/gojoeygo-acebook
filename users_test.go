package main

import (

  "testing"

	_ "github.com/lib/pq"
)

func TestRegisterNewUser(t *testing.T) {

  emptyDB()

  populateDB()

  actual := registerNewUser("NameTest", "SurnameTest", "mail@test.com", "12345")

  if (actual.Email != "mail@test.com" || actual.Password != "12345") {
			t.Fail()
  }
}

func TestValidateUser(t *testing.T) {

  emptyDB()

  populateDB()

  _, password := validateUser("test@test.com", "123")

  if password != "123" {
			t.Fail()
  }
}
