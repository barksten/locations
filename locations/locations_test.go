package locations_test

import (
	. "github.com/barksten/locations/locations"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/modocache/gory"

	"testing"
)

func TestLocations(t *testing.T) {
	defineFactories()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Locations Suite")
}

func defineFactories() {
	gory.Define("location", Location{},
		func(factory gory.Factory) {
			factory["Latitude"] = "?"
		})
}
