package locations

import "labix.org/v2/mgo"

type Location struct {
	Latitude float32 `json:"lat"`
}

func (location *Location) valid() bool {
	return true
}

func fetchAllLocations(db *mgo.Database) []Location {
	locations := []Location{}
	err := db.C("locations").Find(nil).All(&locations)
	if err != nil {
		panic(err)
	}
	return locations
}
