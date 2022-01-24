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
	Links  map[string]string
	Aliens []*Alien
}

// Game world
type Game struct {
	numAliens int
	Cities    map[string]*City
	Aliens    []*Alien
}

// Init sets up the Game
func (g *Game) Init(numAliens int) {
	g.numAliens = numAliens
	g.Cities = make(map[string]*City)
	g.Aliens = make([]*Alien, 0, g.numAliens)

	g.loadWorldMap(g.loadFile())
	if !g.mapIsValid() {
		log.Fatalln("Invalid map: add a line for each City")
	}
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

func (g *Game) loadWorldMap(worldMap []string) {
	// Parse cities and links
	for _, line := range worldMap {
		lineParts := strings.Split(line, " ")
		name := lineParts[0]
		links := make(map[string]string)

		for _, part := range lineParts[1:] {
			directions := strings.Split(part, "=")
			links[directions[1]] = directions[0]
		}

		g.Cities[name] = &City{
			Name:  name,
			Links: links,
		}
	}
}

func (g *Game) mapIsValid() bool {
	// A line is required for each city
	for _, city := range g.Cities {
		for name := range city.Links {
			if _, ok := g.Cities[name]; !ok {
				return false
			}
		}
	}
	return true
}

func (g *Game) initializeAliens() {
	// Aliens start out at random places on the map
	for idx := 0; idx < g.numAliens; idx++ {
		cities := make([]*City, 0, len(g.Cities))
		for _, v := range g.Cities {
			cities = append(cities, v)
		}

		city := cities[rand.Intn(len(cities))]
		alien := &Alien{
			Name: strconv.Itoa(idx),
			City: city,
		}
		g.Aliens = append(g.Aliens, alien)
		city.Aliens = append(city.Aliens, alien)
	}
}

func remove(l []*Alien, item *Alien) []*Alien {
	// Remove an alien from a slice
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func (g *Game) wanderAround(alien *Alien) {
	// Aliens wander around randomly following links
	oldCity := alien.City
	links := oldCity.Links
	alien.NumMoves = alien.NumMoves + 1

	if len(links) > 0 {
		names := make([]string, 0, len(links))
		for k := range links {
			names = append(names, k)
		}

		name := names[rand.Intn(len(names))]
		if city, ok := g.Cities[name]; ok {
			alien.City = city
		} else {
			log.Fatalln("Invalid map: add a line for", name)
		}

		// log.Printf("%s moved to %s from %s", alien.Name, alien.City.Name, oldCity.Name)
		oldCity.Aliens = remove(oldCity.Aliens, alien)
		alien.City.Aliens = append(alien.City.Aliens, alien)
	}
}

func (g *Game) fight(city *City) {
	if len(city.Aliens) > 1 {
		log.Printf("%s has been destroyed by alien %s and alien %s", city.Name, city.Aliens[0].Name, city.Aliens[1].Name)

		// Kill aliens, remove from map, destroy city
		for _, alien := range city.Aliens {
			g.Aliens = remove(g.Aliens, alien)
		}
		for _, otherCity := range g.Cities {
			delete(otherCity.Links, city.Name)
		}
		delete(g.Cities, city.Name)
	}
}

func (g *Game) run() {
	for _, alien := range g.Aliens {
		g.wanderAround(alien)
	}
	for _, city := range g.Cities {
		g.fight(city)
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
			}
		}
		if runAgain {
			g.run()
		} else {
			break
		}
	}
}

// PrintMap prints the current state of the map
func (g *Game) PrintMap() {
	lines := make([]string, len(g.Cities))

	for _, city := range g.Cities {
		line := make([]string, 1+len(city.Links))
		line[0] = city.Name

		for cityName, direction := range city.Links {
			line = append(line, fmt.Sprintf("%s=%s", direction, cityName))
		}
		lines = append(lines, strings.Join(line[:], " "))
	}
	log.Println(strings.Join(lines[:], "\n"))
}
