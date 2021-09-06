package todo

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/nahidhasan98/todo/config"
)

type repoInterface interface {
	createTodo(todo *Todo, author string) error
	getAllTodo(author string) (*[]Todo, error)
	getSingleTodo(author, todoID string) (*Todo, error)
	deleteAllTodo(author string) error
	deleteSingleTodo(author, todoID string) error
	updateTodo(todo *Todo, author, todoID string) error
	getID() string
}

type repoStruct struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func (r *repoStruct) createTodo(todo *Todo, author string) error {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Insert(&todo)

	return err
}

func (r *repoStruct) getAllTodo(author string) (*[]Todo, error) {
	var todo []Todo
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"author": author}).All(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *repoStruct) getSingleTodo(author, todoID string) (*Todo, error) {
	var todo Todo
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"author": author, "_id": todoID}).One(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *repoStruct) deleteAllTodo(author string) error {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	_, err := coll.RemoveAll(bson.M{"author": author})
	if err != nil {
		return err
	}

	return nil
}

func (r *repoStruct) deleteSingleTodo(author, todoID string) error {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Remove(bson.M{"author": author, "_id": todoID})
	if err != nil {
		return err
	}

	return nil
}

func (r *repoStruct) updateTodo(todo *Todo, author, todoID string) error {
	todo.ID = todoID
	todo.Author = author
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	selector := bson.M{"_id": todoID, "author": author}
	err := coll.Update(selector, bson.M{"$set": todo})
	if err != nil {
		return err
	}

	return nil
}

func (r *repoStruct) getID() string {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	count, err := coll.Find(bson.M{}).Count()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("t%v", count+1)
}

func NewRepository(dbSession *mgo.Session) repoInterface {
	return &repoStruct{
		DBSession: dbSession,
		DBName:    config.DBName,
		DBTable:   config.TodoTable,
	}
}
