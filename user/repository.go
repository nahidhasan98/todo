package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/nahidhasan98/todo/config"
)

type repoInterface interface {
	getAllUser() (*[]Data, error)
	getAllUserWithTask() (*[]Data, error)
	getSingleUser(id string) (*Data, error)
	getSingleUserWithTask(id string) (*Data, error)
}

type repoStruct struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func (r *repoStruct) getAllUser() (*[]Data, error) {
	var user []User
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{}).All(&user)
	if err != nil {
		return nil, err
	}

	var resp []Data

	for _, val := range user {
		var tData Data
		tData.User = val
		tData.User.Password = ""
		tData.Task = []Todo{}

		resp = append(resp, tData)
	}

	return &resp, nil
}

func (r *repoStruct) getAllUserWithTask() (*[]Data, error) {
	var user []User
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{}).All(&user)
	if err != nil {
		return nil, err
	}

	var resp []Data

	for _, val := range user {
		val.Password = ""

		var todo []Todo
		coll := r.DBSession.DB(r.DBName).C("todo")
		err := coll.Find(bson.M{"author": val.Username}).All(&todo)
		if err != nil {
			return nil, err
		}

		var tData Data
		tData.User = val
		tData.Task = todo

		resp = append(resp, tData)
	}

	return &resp, nil
}

func (r *repoStruct) getSingleUser(id string) (*Data, error) {
	var user User
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"_id": id}).One(&user)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	resp := &Data{
		User: user,
		Task: []Todo{},
	}

	return resp, nil
}

func (r *repoStruct) getSingleUserWithTask(id string) (*Data, error) {
	var user User
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"_id": id}).One(&user)
	if err != nil {
		return nil, err
	}
	user.Password = ""

	var todo []Todo
	coll = r.DBSession.DB(r.DBName).C("todo")
	err = coll.Find(bson.M{"author": user.Username}).All(&todo)
	if err != nil {
		return nil, err
	}

	resp := &Data{
		User: user,
		Task: todo,
	}

	return resp, nil
}

func NewRepository(dbSession *mgo.Session) repoInterface {
	return &repoStruct{
		DBSession: dbSession,
		DBName:    config.DBName,
		DBTable:   config.UserTable,
	}
}
