package locations_test

import (
	. "github.com/barksten/locations/locations"

	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

/*
Convert JSON data into a slice.
*/
func sliceFromJSON(data []byte) []interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.([]interface{})
}

/*
Convert JSON data into a map.
*/
func mapFromJSON(data []byte) map[string]interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.(map[string]interface{})
}

/*
Server unit tests.
*/
var _ = Describe("Server", func() {
	var dbName string
	var session *DatabaseSession
	var server Server
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	BeforeEach(func() {
		// Set up a new server, connected to a test database,
		// before each test.
		dbName = "locations_test"
		session = NewSession(dbName)
		server = NewServer(session)

		// Record HTTP responses.
		recorder = httptest.NewRecorder()
	})

	AfterEach(func() {
		// Clear the database after each
		// test.
		session.DB(dbName).DropDatabase()
	})

	Describe("GET /locations", func() {

		// Set up a new GET request before every test
		// in this describe block.
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/locations", nil)
		})

		Context("when no locations exist", func() {
			It("returns a status code of 200", func() {
				server.Martini.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))
			})

			It("returns a empty body", func() {
				server.Martini.ServeHTTP(recorder, request)
				Expect(recorder.Body.String()).To(Equal("[]"))
			})
		})
	})
})
