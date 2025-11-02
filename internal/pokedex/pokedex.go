package pokedex

import (
	"fmt"
	"sync"
)

type Pokedex struct {
	mu       sync.Mutex
	pokemons map[string]Pokemon
}

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  struct {
		Hp             int
		Attack         int
		Defense        int
		SpecialAttack  int
		SpecialDefense int
		Speed          int
	}
	Types []string
}

func (p *Pokedex) Add(key string, pokemon Pokemon) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pokemons[key] = pokemon
}

func (p *Pokedex) Get(key string) (Pokemon, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	val, ok := p.pokemons[key]
	if ok == false {
		return Pokemon{}, false
	}
	return val, ok

}

func (p *Pokedex) List() []string {
	result := make([]string, 0)
	for k := range p.pokemons {
		result = append(result, k)
	}
	return result
}

func (p Pokemon) String() string {
	return fmt.Sprintf("Name: %s\n"+
		"Height: %d\n"+
		"Weight: %d\n"+
		"Stats: \n"+
		"	HP: %d\n"+
		"	Attack: %d\n"+
		"	Defense: %d\n"+
		"	Special-Attack: %d\n"+
		"	Special-Defense: %d\n"+
		"	Speed: %d\n"+
		"Types: %v\n", p.Name, p.Height, p.Weight, p.Stats.Hp, p.Stats.Attack, p.Stats.Defense, p.Stats.SpecialAttack, p.Stats.SpecialDefense, p.Stats.Speed, p.Types)
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemons: make(map[string]Pokemon),
	}
}
