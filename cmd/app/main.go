package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/jrhone/tendermint/internal/app"
)

func main() {
	rand.Seed(time.Now().Unix())

	if len(os.Args) == 1 {
		log.Fatalln("Number of aliens is a required argument")
	}

	numAliens, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("Number of aliens should be an int")
	}

	g := &app.Game{}
	g.Init(numAliens)
	g.Start()
	g.PrintMap()
}
