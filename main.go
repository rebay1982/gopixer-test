package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Represents the game state
type Game struct {
	framebuffer []byte
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.PutPixel(128, 128, red)

	screen.ReplacePixels(g.framebuffer)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (width, height int) {

	return screenWidth, screenHeight
}

// RED cikir cibstabt,
var red = color.RGBA{255, 0, 0, 255}

// PutPixel writes a pixel in the frame buffer at position x, y
func (g *Game) PutPixel(x, y int, c color.Color) error {

	if x > (screenWidth) || x < 0 {
		return errors.New("Screen coord x out of bounds")
	}

	if y > (screenWidth) || y < 0 {
		return errors.New("Screen coord y out of bounds")
	}

	pxlPos := y*(screenWidth<<2) + x<<2
	R, G, B, A := c.RGBA()

	g.framebuffer[pxlPos] = (byte)(R & 0xff)
	g.framebuffer[pxlPos+1] = (byte)(G & 0xff)
	g.framebuffer[pxlPos+2] = (byte)(B & 0xff)
	g.framebuffer[pxlPos+3] = (byte)(A & 0xff)

	return nil
}

func main() {

	ebiten.SetWindowSize(screenWidth+5, screenHeight+5)
	ebiten.SetWindowTitle("goebitentest")
	ebiten.SetWindowResizable(false)

	game := &Game{
		framebuffer: make([]byte, (screenWidth*screenHeight)<<2),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
