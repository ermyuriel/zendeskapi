package zendeskapi

import (
	"log"
	"testing"
	"time"
)

func TestUserCreation(t *testing.T) {

	u := User{Name: "test_" + getTestTimestamp(), Email: "test_" + getTestTimestamp() + "@luuna.mx", Verified: true}

	printPrettyStruct(u)
	cu, e := createUser(&u)
	if e != nil {
		log.Println(e)
		t.Fail()

	}

	printPrettyStruct(cu)

}

func TestUserUpdate(t *testing.T) {

	u := User{Name: "test_" + getTestTimestamp(), Email: "test_" + getTestTimestamp() + "@testing.mx", Verified: true}

	printPrettyStruct(u)

	cu, err := createUser(&u)

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	ur, err := getUser(cu.User.ID)

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	printPrettyStruct(ur)

}

func TestUserSearch(t *testing.T) {

	TestUserCreation(t)

	rs, err := searchUser("luuna.mx")

	if err != nil {
		t.Fail()
	}

	printPrettyStruct(rs)

}

func getTestTimestamp() string {

	return time.Now().Format("2006_01_02_03.04.05")

}
