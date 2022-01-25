package app

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapFull(t *testing.T) {
	worldMap := strings.Split(`Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
Baz east=Foo west=Qu-ux
Qu-ux north=Foo
Bee`, "\n")

	g := &Game{}
	var cities []*City
	g.Cities = cities

	g.loadWorldMap(worldMap)
	assert.Equal(t, len(g.Cities), 5)
}

func Test_MapMissingCity(t *testing.T) {
	worldMap := strings.Split(`Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee`, "\n")

	g := &Game{}
	var cities []*City
	g.Cities = cities

	g.loadWorldMap(worldMap)
	assert.Equal(t, len(g.Cities), 5)
}
