package mongoService
import (
	"gopkg.in/mgo.v2"
	"time"
)
//my mlab: "ds043982.mlab.com:43982"
func CreateMongoSession() *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
    //mine
    Addrs:    []string{"ds043982.mlab.com:43982"},
		Timeout:  60 * time.Second,
		Database: "dev_skills",
		Username: "disney",
		Password: "disney",
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	return session

}
