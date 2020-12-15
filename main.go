package main

import (
	"bytes"
	"fmt"
	"gen_alg/organism"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	target := []byte("МИР")
	population := organism.CreatePopulation(target)

	found := false
	generation := 0
	for !found {
		generation++
		bestOrganism := organism.GetBest(population)
		fmt.Printf("\r generation: %d | %s | fitness: %2f", generation, string(bestOrganism.DNA), bestOrganism.Fitness)

		if bytes.Compare(bestOrganism.DNA, target) == 0 {
			found = true
		} else {
			maxFitness := bestOrganism.Fitness
			pool := organism.CreatePool(population, target, maxFitness)
			population = organism.NaturalSelection(pool, population, target)
		}

	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime taken: %s\n", elapsed)
}
