package main

import (
	"log"
	"testing"
	"time"
)

func TestUserCreation(t *testing.T) {

	u := UserCreate{Name: "test_" + getTestTimestamp(), Email: "test_" + getTestTimestamp() + "@eucj.mx", Verified: true}

	printPrettyStruct(u)
	cu, e, er := CreateUser(&u)
	if e != nil {
		log.Println(e)
		t.Fail()

	}

	if er != nil {
		log.Println("User not created")
		printPrettyStruct(er)
		t.Fail()
	} else {

		printPrettyStruct(cu)

	}

}

func TestUserUpdate(t *testing.T) {

}

func TestUserSearch(t *testing.T) {

	TestUserCreation(t)

	rs, err, er := SearchUser("eucj.mx")

	if err != nil {
		t.Fail()
	}

	if er != nil {
		log.Println("Invalid search")
		printPrettyStruct(er)

	} else {

		printPrettyStruct(rs)

	}

}

func TestRelationshiTypeCreate(t *testing.T) {

	err, er := CreateRelationshipType("test_object", "has_test_object", "test_object")

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
		t.Fail()
	}

	err, er = CreateRelationshipType("test_object", "has_test_objects", []string{"test_object"})

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
		t.Fail()
	}

}

func TestObjectRecordCreatet(t *testing.T) {

	ts := getTestTimestamp()

	m := map[string]interface{}{"id": ts, "name": ts}
	orr, err, er := CreateObjectRecord("test_object", m)

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
		t.Fail()
	}
	printPrettyStruct(orr)

}

func TestRelationshipRecordSet(t *testing.T) {

	ts := getTestTimestamp()
	uc := UserCreate{Name: "test_" + ts, Email: "test_" + ts + "@eucj.mx", Verified: true}
	CreateUser(&uc)
	rs, _, _ := SearchUser("@eucj.mx")

	u := rs[0]
	uid := u.ID

	m := map[string]interface{}{"id": ts, "name": ts}
	orr, err, er := CreateObjectRecord("test_object", m)

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
		t.Fail()
	}

	tid := orr.Data.ID

	SetRelationship(uid, "has_test_objects", tid)

}

func getTestTimestamp() string {

	return time.Now().Format("2006_01_02_03.04.05")

}
