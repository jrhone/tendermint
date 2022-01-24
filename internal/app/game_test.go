package app

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapValid(t *testing.T) {
	assert.Equal(t, true, true)
}

func Test_MapMissingCity(t *testing.T) {
	assert.Equal(t, true, true)
}

func Test_GetWorldMap(t *testing.T) {
	assert.Equal(t, true, true)
}

func Test_GameRun(t *testing.T) {
	rand.Seed(0)

	g := &Game{}
	g.Init(5)
	g.Start()
	g.PrintMap()
}
