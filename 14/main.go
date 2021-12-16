package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	aoc "github.com/ThomasVictoria/advent_of_code_2021"
	"github.com/davecgh/go-spew/spew"
)

type Polymers struct {
	Map      map[string]int
	Template []Polymer
	Inserted int
}

type Polymer struct {
	Left     string
	Right    string
	ToInsert string
}

type Key string

func (p *Polymers) fillTemplate(raw []string) {
	for _, line := range raw {

		splitted := strings.Split(line, " -> ")

		p.Template = append(p.Template, Polymer{
			Left:     fmt.Sprintf("%c", splitted[0][0]),
			Right:    fmt.Sprintf("%c", splitted[0][1]),
			ToInsert: fmt.Sprintf("%c", splitted[1][0]),
		})
	}
}

func (key *Key) recurse(polymers Polymers, index int, previousInsert string) *Key {

	if len(*key)-1 <= index+polymers.Inserted {
		ok := (*key)[:index+polymers.Inserted] + Key(previousInsert) + (*key)[index+polymers.Inserted:]
		key = &ok

		return key
	}

	char := fmt.Sprintf("%c", (*key)[index+polymers.Inserted])
	nextChar := fmt.Sprintf("%c", (*key)[index+polymers.Inserted+1])

	if previousInsert != "" {
		ok := (*key)[:index+polymers.Inserted] + Key(previousInsert) + (*key)[index+polymers.Inserted:]
		key = &ok
		polymers.Inserted++
	}

	var toInsert string

	for _, t := range polymers.Template {
		if t.Left == char && t.Right == nextChar {
			toInsert = t.ToInsert
			if val, ok := polymers.Map[toInsert]; ok {
				polymers.Map[toInsert] = 1 + val
			} else {
				polymers.Map[toInsert] = 1
			}
		}
	}

	return key.recurse(polymers, index+1, toInsert)
}

func main() {
	rawData := strings.Split(aoc.DecodeInput("14"), "\n")

	tmp := Key(rawData[0])
	template := &tmp
	data := append(rawData[:0], rawData[1:]...)
	data = append(data[:0], data[1:]...)

	polymers := Polymers{
		Map: make(map[string]int),
	}

	polymers.fillTemplate(data)

	start := time.Now()

	i := 0
	for i < 40 {
		template = template.recurse(polymers, 0, "")
		spew.Dump(i)
		i++
	}

	lowest := math.Inf(1)
	highest := 0

	for _, polymer := range polymers.Template {
		res := strings.Count(string(*template), polymer.ToInsert)
		if float64(res) < lowest {
			lowest = float64(res)
		}
		if res > highest {
			highest = res
		}
	}

	elapsed := time.Since(start)

	log.Printf("Binomial took %s", elapsed)
	fmt.Println(highest - int(lowest))
}
