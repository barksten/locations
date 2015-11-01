package locations

import (
	"github.com/gin-gonic/gin"
	"labix.org/v2/mgo"
)

type Server struct {
	*gin.Engine
}

/*
Create a new *martini.ClassicMartini server.
We'll use a JSON renderer and our MongoDB
database handler. We define two routes:
"GET /signatures" and "POST /signatures".
*/
func NewServer(session *DatabaseSession) Server {
	// Create the server and set up middleware.
	router := Server{gin.Default()}
	mongo := router.Group("/", session.Database)

	mongo.GET("/locations", func(c *gin.Context) {
		db := c.MustGet("mongo_session").(*mgo.Database)
		c.JSON(200, fetchAllLocations(db))
	})

	mongo.POST("/locations", func(c *gin.Context) {
		db := c.MustGet("mongo_session").(*mgo.Database)
		var location Location
		if c.BindJSON(&location) == nil && location.valid() {
			err := db.C("locations").Insert(location)
			if err == nil {
				c.JSON(201, location)
			} else {
				c.JSON(400, map[string]string{
					"error": err.Error(),
				})
			}
		} else {
			c.JSON(400, map[string]string{
				"error": "Not a valid location",
			})
		}
	})

	// Return the server. Call Run() on the server to
	// begin listening for HTTP requests.
	return router
}
