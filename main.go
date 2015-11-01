package main

import "github.com/barksten/locations/locations"

/*
Create a new MongoDB session, using a database
named "locations". Create a new server using
that session, then begin listening for HTTP requests.
*/
func main() {
	session := locations.NewSession("locations")
	server := locations.NewServer(session)
	server.Run()
}
