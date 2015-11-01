package locations

import (
	"github.com/gin-gonic/gin"
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

// http://blog.mongodb.org/post/80579086742/running-mongodb-queries-concurrently-with-go
// Request a socket connection from the session to process our query.
// Close the session when the goroutine exits and put the connection
// back
// into the pool.
func (session *DatabaseSession) Database(c *gin.Context) {
	s := session.Clone()
	defer s.Close()

	c.Set("mongo_session", s.DB(session.databaseName))
	c.Next()
}
