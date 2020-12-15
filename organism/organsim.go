package organism

import (
	"math/rand"
)

// Коэффициент мутации
const MutationRate = 0.3

// Размер популяции
const PopSize = 500

// Структура организма
type Organism struct {
	DNA     []byte
	Fitness float64
}

// Создать организм
func CreateOrganism(target []byte) (organism Organism) {
	ba := make([]byte, len(target))
	for i := 0; i < len(target); i++ {
		ba[i] = byte(rand.Intn(1072) + 33)
	}
	organism = Organism{
		DNA:     ba,
		Fitness: 0,
	}
	organism.CalcFitness(target)
	return
}

// Создает популяцию
func CreatePopulation(target []byte) (population []Organism) {
	population = make([]Organism, PopSize)
	for i := 0; i < PopSize; i++ {
		population[i] = CreateOrganism(target)
	}
	return
}

//Вычисляет текущую присобленность организма
func (d *Organism) CalcFitness(target []byte) {
	score := 0
	for i := 0; i < len(d.DNA); i++ {
		if d.DNA[i] == target[i] {
			score++
		}
	}
	d.Fitness = float64(score) / float64(len(d.DNA))
	return
}

// Механизм селекции, основанный на рулеточной селекции
func CreatePool(population []Organism, target []byte, maxFitness float64) (pool []Organism) {
	pool = make([]Organism, 0)
	// create a pool for next generation
	for i := 0; i < len(population); i++ {
		population[i].CalcFitness(target)
		num := int((population[i].Fitness / maxFitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

// Выполняет срещивание особей, основанное на одноточечном кроссингорвере со случайно точкой разрыва
func NaturalSelection(pool []Organism, population []Organism, target []byte) []Organism {
	next := make([]Organism, len(population))

	for i := 0; i < len(population); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]

		child := Crossover(a, b)
		child.Mutate()
		child.CalcFitness(target)

		next[i] = child
	}
	return next
}

// Кроссинговер
func Crossover(d1 Organism, d2 Organism) Organism {
	child := Organism{
		DNA:     make([]byte, len(d1.DNA)),
		Fitness: 0,
	}
	mid := rand.Intn(len(d1.DNA))
	for i := 0; i < len(d1.DNA); i++ {
		if i > mid {
			child.DNA[i] = d1.DNA[i]
		} else {
			child.DNA[i] = d2.DNA[i]
		}

	}
	return child
}

// Производит мутацию организма
func (d *Organism) Mutate() {
	for i := 0; i < len(d.DNA); i++ {
		if rand.Float64() < MutationRate {
			d.DNA[i] = byte(rand.Intn(1072) + 33)
		}
	}
}

// Возращает наиболее присобленный организм популяции
func GetBest(population []Organism) Organism {
	best := 0.0
	index := 0
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > best {
			index = i
			best = population[i].Fitness
		}
	}
	return population[index]
}
