package auth

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/nahidhasan98/todo/config"
)

type repoInterface interface {
	getUserByUsername(username string) (*User, error)
}

type repoStruct struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func (r *repoStruct) getUserByUsername(username string) (*User, error) {
	var user User
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewRepository(dbSession *mgo.Session) *repoStruct {
	return &repoStruct{
		DBSession: dbSession,
		DBName:    config.DBName,
		DBTable:   config.AuthTable,
	}
}
