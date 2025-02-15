package game

import (
	"fmt"
	"image/color"

	"time"

	"goquest/pkg/player"
	"goquest/pkg/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game представляет игру с миром, погодой и игроком.
type Game struct {
	player *player.Player
	world  *world.World
}

// NewGame создает новый объект игры.
func NewGame(playerName string) *Game {
	return &Game{
		player: player.New(playerName),
		world:  world.New(),
	}
}

// Run запускает игру.
func (g *Game) Run() error {

	ebiten.SetWindowSize(1000, 600)
	ebiten.SetWindowTitle("Open World Game with Crafting and Weather")

	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}

// Update обновляет состояние игры (каждый кадр).
func (g *Game) Update() error {
	// Изменяем погодные условия каждую секунду
	if time.Now().Second()%5 == 0 {
		g.world.WeatherChange()
	}

	// Управление движением персонажа
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.ChangeLocation("Вперед")
		g.player.CurrentLocation = "вперед"
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.ChangeLocation("Назад")
		g.player.CurrentLocation = "назад"
	}

	// Крафт предметов
	if ebiten.IsKeyPressed(ebiten.KeyC) {
		g.player.Craft("Меч")
	}

	// Строительство
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		g.player.Build("Укрытие")
	}

	// Для симуляции динамической погоды
	g.world.WeatherChange()

	return nil
}

// Draw отрисовывает все элементы игры (каждый кадр).
func (g *Game) Draw(screen *ebiten.Image) {
	// Фон
	screen.Fill(color.RGBA{R: 51, G: 77, B: 102, A: 255})

	// Отображаем статус игрока
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Игрок: %s\nЗдоровье: %d\nМагия: %d\nПогода: %s\nМестоположение: %s",
		g.player.Name, g.player.Health, g.player.Magic, g.world.Weather, g.player.CurrentLocation))

	// Выводим информацию о мире
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Мир: %s\nСтроительство: %s\nКрафт: %s", g.world.Terrain, g.world.Building, g.world.Crafting), 10, 100)
}

// Layout задает размер экрана
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
