package world

import (
	"time"

	"math/rand"
)

// World представляет мир, в котором происходит игра.
type World struct {
	Weather   string
	Terrain   string
	Building  string
	Crafting  string
	LocationX int
	LocationY int
}

// WeatherChange меняет погодные условия в игре.
func (w *World) WeatherChange() {
	weatherConditions := []string{"Дождь", "Гроза", "Туман", "Ясная погода"}
	rand.Seed(time.Now().UnixNano())
	w.Weather = weatherConditions[rand.Intn(len(weatherConditions))]
}

// New создает новый мир
func New() *World {

	w := World{}
	w.WeatherChange()
	return &w
}
