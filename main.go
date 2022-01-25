package main

import (
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
)

const (
	winWidth  = 600
	winHeight = 300
)

// TODO create a buffer and draw it
func main() {
	driver.Main(func(s screen.Screen) {
		window, err := s.NewWindow(&screen.NewWindowOptions{Width: winWidth, Height: winHeight, Title: "gamma-calculator"})
		if err != nil {
			log.Fatal(err)
		}
		defer window.Release()
		for {
			switch e := window.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			}
		}
	})
}
