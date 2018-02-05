package DataAccess

import (
	"gopkg.in/mgo.v2"
	"TestProject/Models"
	"gopkg.in/mgo.v2/bson"
	"TestProject/Settings"
	"log"
)

var session *mgo.Session
var c *mgo.Collection

func GetTeams()([]Models.Team, error){
	c,err := getSession()
	if err != nil {
		return nil,err
	}

	var teams []Models.Team
	err = c.Find(bson.M{}).All(&teams)
	if err != nil {
		return nil,err
	}

	return teams,err
}

func GetTeam(id string)(*Models.Team, error){
	c,err := getSession()
	if err != nil {
		return nil,err
	}

	var team Models.Team
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).
		One(&team)
	if err != nil {
		return nil,err
	}

	return &team,err
}

func CreateTeam(team Models.Team)(error){
	c,err := getSession()
	if err != nil {
		return err
	}
	err = c.Insert(team)
	return err
}

func UpdateTeam(id string, team Models.Team)(error){
	c,err := getSession()
	if err != nil {
		return err
	}
	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, team)
	return err
}

func DeleteTeam(id string)(error){
	c,err := getSession()
	if err != nil {
		return err
	}

	err = c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func CloseSession(){
	session.Close()
}

func getSession() (*mgo.Collection, error){
	var err error
	if session == nil {
		session, err= mgo.Dial(Settings.Host)
		if err != nil {
			log.Fatalf("CreateSession: %s\n", err)
			return nil,err
		}
	}

	if c == nil{
		c = session.DB(Settings.DbName).C(Settings.CollectionName)
	}
	return c,err
}