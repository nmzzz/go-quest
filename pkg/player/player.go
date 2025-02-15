package player

import (
	"fmt"
	"time"
)

// Player представляет игрока в игре.
type Player struct {
	Name            string
	Health          int
	Magic           int
	Inventory       []string
	CurrentLocation string
}

// New создает нового игрока с заданным именем.
func New(name string) *Player {
	return &Player{
		Name:            name,
		Health:          100,
		Magic:           50,
		Inventory:       []string{},
		CurrentLocation: "123154",
	}
}

// AddToInventory добавляет предмет в инвентарь игрока.
func (p *Player) AddToInventory(item string) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s добавил %s в инвентарь.\n", p.Name, item)
}

// UseMagic использует магический артефакт.
func (p *Player) UseMagic(artifact string) {
	fmt.Printf("%s использует магический артефакт: %s.\n", p.Name, artifact)
	p.Magic -= 10
}

// Craft создает новый предмет с использованием материалов.
func (p *Player) Craft(tool string) {
	fmt.Printf("%s создает инструмент: %s.\n", p.Name, tool)
	p.AddToInventory(tool)
}

// Build строит сооружение (например, укрытие).
func (p *Player) Build(structure string) {
	fmt.Printf("%s строит: %s.\n", p.Name, structure)
}

// ChangeLocation изменяет местоположение игрока.
func (p *Player) ChangeLocation(newLocation string) {
	p.CurrentLocation = newLocation
	fmt.Printf("%s перемещается в: %s.\n", p.Name, newLocation)
}

// WeatherChange симулирует изменение погоды.
func (p *Player) WeatherChange() {
	weatherConditions := []string{"Дождь", "Гроза", "Туман", "Ясная погода"}
	randIndex := time.Now().Unix() % int64(len(weatherConditions))
	weather := weatherConditions[randIndex]
	fmt.Printf("Погода изменилась: %s.\n", weather)
}

// PrintStatus выводит текущий статус игрока.
func (p *Player) PrintStatus() {
	fmt.Printf("Игрок: %s\nЗдоровье: %d\nМагия: %d\nМестоположение: %s\nИнвентарь: %v\n",
		p.Name, p.Health, p.Magic, p.CurrentLocation, p.Inventory)
}
