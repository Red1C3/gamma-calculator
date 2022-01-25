package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
)

const (
	winWidth  = 600
	winHeight = 300

	R = 1
	G = 2
	B = 4
)

var (
	mainBuffer screen.Buffer
	channel    uint8 = 1
	currentmed uint8 = 128
)

func fillMainBuffer(med uint8, chn uint8) {
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
					mainBuffer.RGBA().SetRGBA(i, j, color.RGBA{255 * (chn & 1), 255 * ((chn & 2) >> 1), 255 * ((chn & 4) >> 2), 255})
				}
			} else {
				mainBuffer.RGBA().SetRGBA(i, j, color.RGBA{med * (chn & 1), med * ((chn & 2) >> 1), med * ((chn & 4) >> 2), 255})
			}
		}
		startsWhite = !startsWhite
	}
}
func printGamma() {
	log.Printf("gamma = %f\n", math.Log(0.5)/math.Log(float64(currentmed)))
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
		fillMainBuffer(currentmed, channel)
		for {
			switch e := window.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case key.Event:
				if e.Direction == key.DirRelease {
					if e.Code == key.CodeLeftArrow {
						currentmed--
					}
					if e.Code == key.CodeRightArrow {
						currentmed++
					}
					if e.Code == key.CodeR {
						channel = R
					}
					if e.Code == key.CodeB {
						channel = B
					}
					if e.Code == key.CodeG {
						channel = G
					}
					fillMainBuffer(currentmed, channel)
					printGamma()
				}
			}
			window.Upload(image.Pt(0, 0), mainBuffer, mainBuffer.Bounds())
			window.Publish()
		}
	})
}
