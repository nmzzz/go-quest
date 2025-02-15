package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/exp/rand"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	playerX, playerY float64 = screenWidth / 2, screenHeight / 2
	playerSpeed      float64 = 2
	playerImage      *ebiten.Image
)

type Game struct{}

func (g *Game) Update() error {
	// Управление игроком
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		playerY -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		playerY += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		playerX -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		playerX += playerSpeed
	}

	// Здесь можно добавить логику для сбора ресурсов, крафта, создания ловушек, артефактов и т.д.

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Очищаем экран
	screen.Fill(color.RGBA{0, 0, 0, 255}) // Черный цвет

	// Отображение игрока
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(playerX, playerY)
	screen.DrawImage(playerImage, op)

	// Рисуем текст на экране
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player X: %.2f, Y: %.2f", playerX, playerY))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func loadPlayerImage() (*ebiten.Image, error) {
	// Создаем изображение для игрока (красный квадрат)
	img := ebiten.NewImage(32, 32)
	img.Fill(color.RGBA{255, 0, 0, 255}) // Красный цвет
	return img, nil
}

func main() {
	// Инициализация случайного числа (для генерации случайных событий)
	rand.Seed(uint64(time.Now().UnixNano()))

	var err error
	playerImage, err = loadPlayerImage()
	if err != nil {
		log.Fatal(err)
	}

	// Создание нового игрового объекта
	game := &Game{}

	// Запуск игры
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
