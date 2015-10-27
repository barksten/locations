package locations

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"
)

/*
Wrap the Martini server struct.
*/
type Server *martini.ClassicMartini

/*
Create a new *martini.ClassicMartini server.
We'll use a JSON renderer and our MongoDB
database handler. We define two routes:
"GET /signatures" and "POST /signatures".
*/
func NewServer(session *DatabaseSession) Server {
	// Create the server and set up middleware.
	m := Server(martini.Classic())
	m.Martini.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))
	m.Martini.Use(session.Database())

	m.Get("/locations", func(r render.Render, db *mgo.Database) {
		r.JSON(200, fetchAllLocations(db))
	})

	m.Post("/locations", binding.Json(Location{}),
		func(location Location,
			r render.Render,
			db *mgo.Database) {

			if location.valid() {
				err := db.C("locations").Insert(location)
				if err == nil {
					r.JSON(201, location)
				} else {
					r.JSON(400, map[string]string{
						"error": err.Error(),
					})
				}
			} else {
				r.JSON(400, map[string]string{
					"error": "Not a valid signature",
				})
			}
		})

	// Return the server. Call Run() on the server to
	// begin listening for HTTP requests.
	return m
}
