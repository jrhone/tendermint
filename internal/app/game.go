package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

const iterations = 10000
const worldMapPath = "configs/world_map.txt"

// Alien in the game
type Alien struct {
	Name     string
	City     *City
	NumMoves int
}

// City in the game
type City struct {
	Name   string
	Links  []*Link
	Aliens []*Alien
}

// Link connecting a city
type Link struct {
	City      *City
	Direction string
}

// Game world
type Game struct {
	numAliens int
	Cities    []*City
	Aliens    []*Alien
}

// Init sets up the Game
func (g *Game) Init(numAliens int) {

	var cities []*City
	g.Cities = cities
	g.numAliens = numAliens
	g.Aliens = make([]*Alien, 0, g.numAliens)

	g.loadWorldMap(g.loadFile())
	g.initializeAliens()
}

func (g *Game) loadFile() []string {
	data, err := ioutil.ReadFile(worldMapPath)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(data), "\n")
	filtered := []string{}

	// Remove blank lines
	for i := range lines {
		line := strings.TrimSpace(lines[i])
		if len(line) > 0 {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func (g *Game) getOrCreateCity(name string, links []*Link) *City {
	for _, c := range g.Cities {
		if c.Name == name {
			if links != nil {
				c.Links = links
			}
			return c
		}
	}

	city := &City{
		Name:  name,
		Links: links,
	}
	g.Cities = append(g.Cities, city)
	return city
}

func (g *Game) loadWorldMap(worldMap []string) {
	// Parse cities and links
	for _, line := range worldMap {
		lineParts := strings.Split(line, " ")
		name := lineParts[0]
		var links []*Link

		for _, part := range lineParts[1:] {
			directions := strings.Split(part, "=")
			city := g.getOrCreateCity(directions[1], nil)
			links = append(links, &Link{
				City:      city,
				Direction: directions[0],
			})
		}

		g.getOrCreateCity(name, links)
	}
}

func (g *Game) initializeAliens() {
	// Aliens start out at random places on the map
	for idx := 0; idx < g.numAliens; idx++ {
		city := g.Cities[rand.Intn(len(g.Cities))]
		alien := &Alien{
			Name: strconv.Itoa(idx),
			City: city,
		}
		g.Aliens = append(g.Aliens, alien)
		city.Aliens = append(city.Aliens, alien)
	}
}

func removeAlien(l []*Alien, item *Alien) []*Alien {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func removeCity(l []*City, item *City) []*City {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func removeLink(l []*Link, item *City) []*Link {
	for i, other := range l {
		if other.City == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

// WanderAround changes the location of the alien
func (alien *Alien) WanderAround() {
	// Aliens wander around randomly following links
	oldCity := alien.City
	links := oldCity.Links
	alien.NumMoves = alien.NumMoves + 1

	if len(links) > 0 {
		alien.City = links[rand.Intn(len(links))].City
		oldCity.Aliens = removeAlien(oldCity.Aliens, alien)
		alien.City.Aliens = append(alien.City.Aliens, alien)
		// log.Printf("%s moved to %s from %s", alien.Name, alien.City.Name, oldCity.Name)
	}
}

func (g *Game) fight(city *City) {
	if len(city.Aliens) < 2 {
		return
	}
	log.Printf("%s has been destroyed by alien %s and alien %s", city.Name, city.Aliens[0].Name, city.Aliens[1].Name)

	// Kill aliens, remove from map, destroy city
	for _, alien := range city.Aliens {
		g.Aliens = removeAlien(g.Aliens, alien)
	}
	for _, otherCity := range g.Cities {
		otherCity.Links = removeLink(otherCity.Links, city)
	}
	g.Cities = removeCity(g.Cities, city)
}

func (g *Game) run() {
	for _, alien := range g.Aliens {
		alien.WanderAround()
	}
	for i := len(g.Cities) - 1; i >= 0; i-- {
		g.fight(g.Cities[i])
	}
}

// Start controls the game loop
func (g *Game) Start() {
	// Run until all the aliens have been destroyed or each alien has moved at least 10,000 times.
	for {
		runAgain := false
		for _, alien := range g.Aliens {
			if alien.NumMoves < iterations {
				runAgain = true
				break
			}
		}
		if !runAgain {
			break
		}
		g.run()
	}
}

// PrintMap prints the current state of the map
func (g *Game) PrintMap() {
	lines := make([]string, len(g.Cities))

	for _, city := range g.Cities {
		line := make([]string, 1+len(city.Links))
		line[0] = city.Name

		for _, link := range city.Links {
			line = append(line, fmt.Sprintf("%s=%s", link.Direction, link.City.Name))
		}
		lines = append(lines, strings.Join(line[:], " "))
	}
	log.Println(strings.Join(lines[:], "\n"))
}
