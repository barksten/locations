package locations

import (
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

type DatabaseSession struct {
	*mgo.Session
	databaseName string
}

/*
Connect to the local MongoDB and set up the database.
*/
func NewSession(name string) *DatabaseSession {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return &DatabaseSession{session, name}
}

/*
	Martini lets you inject parameters for routing handlers
	by using `context.Map()`. I'll pass each route handler
	a instance of an *mgo.Database, so they can
	get and insert signatures.

	For more information, check out:
	http://blog.gopheracademy.com/day-11-martini
*/
func (session *DatabaseSession) Database() martini.Handler {
	return func(context martini.Context) {
		s := session.Clone()
		context.Map(s.DB(session.databaseName))
		defer s.Close()
		context.Next()
	}
}
