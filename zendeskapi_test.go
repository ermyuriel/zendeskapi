package zendeskapi

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestUserCreation(t *testing.T) {

	cu, e, er := CreateUser("test_"+getTestTimestamp(), "test_"+getTestTimestamp()+"@eucj.mx")
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
		log.Println(err)
		t.Fail()
	}

	if er != nil {
		log.Println("Invalid search")
		printPrettyStruct(er)

	} else {

		printPrettyStruct(rs)

	}

}

func TestObjectTypeCreate(t *testing.T) {

	ts := getTestTimestamp()
	s := TestType{1, ts}

	err, er := CreateObjectType(s)

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Object not created")
		printPrettyStruct(er)

	}

}
func TestRelationshipTypeCreate(t *testing.T) {

	err, er := CreateRelationshipType("zen:user", "user_has_test_object", "test_object")

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)

	}

	err, er = CreateRelationshipType("zen:user", "user_has_test_objects", []string{"test_object"})

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)

	}

	err, er = CreateRelationshipType("test_object", "test_object_has_test_objects", []string{"test_object"})

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)

	}

	err, er = CreateRelationshipType("test_object", "test_object_has_test_object", "test_object")

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)

	}
}

func TestObjectRecordCreatet(t *testing.T) {

	ts := getTestTimestamp()

	m := TestType{1, ts}

	orr, err, er := CreateObjectRecord("test_object", m)

	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Object not created")
		printPrettyStruct(er)
	}
	printPrettyStruct(orr)

}

func TestRelationshipRecordSet(t *testing.T) {

	ts := getTestTimestamp()

	CreateUser("test_"+ts, "test_"+ts+"@eucj.mx")
	us, _, _ := SearchUser("test_" + ts + "@eucj.mx")
	uid := us[0].ID
	m := TestType{1, ts}
	oss, _, _ := CreateObjectRecord("test_object", m)
	sid := oss.Data.ID
	t1, _, _ := CreateObjectRecord("test_object", m)
	t2, _, _ := CreateObjectRecord("test_object", m)
	t1ID := t1.Data.ID
	t2ID := t2.Data.ID

	err, er := CreateRelationshipRecord(fmt.Sprintf("zen:user:%v", uid), "user_has_test_object", t1ID)
	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
	}

	err, er = CreateRelationshipRecord(fmt.Sprintf("zen:user:%v", uid), "user_has_test_objects", t1ID)
	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
	}

	err, er = CreateRelationshipRecord(fmt.Sprintf("zen:user:%v", uid), "user_has_test_objects", t2ID)
	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
	}

	err, er = CreateRelationshipRecord(sid, "test_object_has_test_object", t1ID)
	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
	}

	err, er = CreateRelationshipRecord(sid, "test_object_has_test_objects", t1ID)
	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
	}

	err, er = CreateRelationshipRecord(sid, "test_object_has_test_objects", t2ID)
	if err != nil {
		log.Println(err)
		t.Fail()

	}

	if er != nil {
		log.Println("Relationship not created")
		printPrettyStruct(er)
	}

	rs, _, _ := ListRelationships("zen:user:"+fmt.Sprintf("%v", uid), "user_has_test_objects")

	printPrettyStruct(rs)
}

func TestRelationshipRecordDelete(t *testing.T) {

	ts := getTestTimestamp()

	CreateUser("test_"+ts, "test_"+ts+"@eucj.mx")
	us, _, _ := SearchUser("test_" + ts + "@eucj.mx")
	uid := us[0].ID

	m := TestType{1, ts}

	t1, _, _ := CreateObjectRecord("test_object", m)

	t1ID := t1.Data.ID

	CreateRelationshipRecord(fmt.Sprintf("zen:user:%v", uid), "user_has_test_object", t1ID)

	rs, _, _ := ListRelationships("zen:user:"+fmt.Sprintf("%v", uid), "user_has_test_object")
	log.Println(len(rs))
	for _, r := range rs {

		DeleteRelationshipRecord(r.ID)

	}

	rs, _, _ = ListRelationships("zen:user:"+fmt.Sprintf("%v", uid), "user_has_test_object")

	if len(rs) != 0 {
		t.Fail()
	}

}
func TestStructToSchema(t *testing.T) {
	s := TestType{1, "test"}
	if StructToSchema(s) != "{\"data\": {\"key\":\"testtype\",\"schema\": {\"properties\": {\"x\":{\"type\":\"number\",\"description\": \"autogenerated\"},\"y\":{\"type\":\"string\",\"description\": \"autogenerated\"}}}}}" {

		t.Fail()

	}
}

func getTestTimestamp() string {

	return time.Now().Format("2006_01_02_03-04-05")

}

type TestType struct {
	X int    `json:"x"`
	Y string `json:"y"`
}
