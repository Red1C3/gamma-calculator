package main

import (
	"image"
	"image/color"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
)

const (
	winWidth  = 600
	winHeight = 300
)

var (
	mainBuffer screen.Buffer
)

func fillMainBuffer(medGray uint8) {
	startsWhite := true
	for i := 0; i < mainBuffer.Size().X; i++ {
		var div int
		if startsWhite {
			div = 1
		} else {
			div = 0
		}
		for j := 0; j < mainBuffer.Size().Y; j++ {
			if i < mainBuffer.Size().X/2 {
				if j%2 == div {
					mainBuffer.RGBA().SetRGBA(i, j, color.RGBA{0, 0, 0, 255})
				} else {
					mainBuffer.RGBA().SetRGBA(i, j, color.RGBA{255, 255, 255, 255})
				}
			} else {
				mainBuffer.RGBA().SetRGBA(i, j, color.RGBA{medGray, medGray, medGray, 255})
			}
		}
		startsWhite = !startsWhite
	}
}
func main() {
	driver.Main(func(s screen.Screen) {
		window, err := s.NewWindow(&screen.NewWindowOptions{Width: winWidth, Height: winHeight, Title: "gamma-calculator"})
		if err != nil {
			log.Fatal(err)
		}
		defer window.Release()
		mainBuffer, err = s.NewBuffer(image.Point{winWidth, winHeight})
		if err != nil {
			log.Fatal(err)
		}
		fillMainBuffer(128)
		for {
			switch e := window.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			}
			window.Upload(image.Pt(0, 0), mainBuffer, mainBuffer.Bounds())
			window.Publish()
		}
	})
}
