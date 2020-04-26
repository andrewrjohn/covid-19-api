package ui

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView("layout", -1, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	info(g)
	stats(g)
	quitButton(g)
	return nil
}

func quitButton(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("quitButton", maxX-6, maxY-3, maxX-1, maxY-1); err != nil {
		fmt.Fprintln(v, "Quit")
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	if err := g.SetKeybinding("quitButton", gocui.MouseLeft, gocui.ModNone, quit); err != nil {
		// handle error
	}

	return nil
}

func info(g *gocui.Gui) error {
	if v, err := g.SetView("info", 0, 0, 20, 20); err != nil {
		v.Title = "Info"
		fmt.Fprintln(v, "Test")
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}
func stats(g *gocui.Gui) error {
	if v, err := g.SetView("stats", 21, 0, 41, 20); err != nil {
		v.Title = "Stats"
		fmt.Fprintln(v, "test")
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}

// InitUI Starts the graphical user interface to monitor the server
func InitUI() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Mouse = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
