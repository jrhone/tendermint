package app

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapValid(t *testing.T) {
	worldMap := strings.Split(`
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
Baz east=Foo west=Qu-ux
Qu-ux north=Foo
Bee`, "\n")

	g := &Game{}
	g.Cities = make(map[string]*City)

	g.loadWorldMap(worldMap)
	assert.Equal(t, g.mapIsValid(), true)
	// assert g.Cities
}

func Test_MapMissingCity(t *testing.T) {
	worldMap := strings.Split(`
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee`, "\n")

	g := &Game{}
	g.Cities = make(map[string]*City)

	g.loadWorldMap(worldMap)
	assert.Equal(t, g.mapIsValid(), false)
	// assert g.Cities and Links
}
