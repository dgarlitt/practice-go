package models
import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
)

type Person struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Firstname string
	Lastname string
	Companyname string
	Address string
	City string
	State string
	Zip string
	Phone string
	Email string
	Web string

}

func AddPerson(session *mgo.Session, newPerson Person) {
	c := session.DB("dev_skills").C("persons_lf")
	fmt.Println(c)
	fmt.Println("Done");

	err := c.Insert(&newPerson)
	if err != nil {
		fmt.Println("ERROR INSERTING")
		log.Fatal(err)
	}
}

func GetPersons(session *mgo.Session) []Person {
	var results []Person
	err := session.DB("dev_skills").C("persons_lf").Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	return results
}


func GetPerson(session *mgo.Session, name string) Person{
	result := Person{}
	err := session.DB("dev_skills").C("persons_lf").Find(bson.M{"name": name}).Select(bson.M{"name": 0}).One(&result)
	if err != nil {
    fmt.Println("Did not find this person., ", name)

	}
	return result
}

func RemoveAllPersons(session *mgo.Session){
	_, err := session.DB("dev_skills").C("persons_lf").RemoveAll(bson.M{})
	if err != nil{
		panic(err)
	}

}
