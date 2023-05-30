package controllers

import (
	"google.golang.org/genproto/googleapis/cloud/oslogin/common"
	"gopkg.in/mgo.v2"
)

// Struct for maintaining HTTP Request Context
type Context struct {
	MongoSession *mgo.Session
}

// Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// returns mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
