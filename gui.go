package main

import (
	"github.com/jroimartin/gocui"
	"fmt"
	"io/ioutil"
	"io"
	"strings"
	"log"
)

const (
	VIEW_GROUP = "group"
)

type GUI struct {
	gui *gocui.Gui
	collection *Collection
}

func NewGui() *GUI {

	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		log.Panicln(err)
	}

	gui := GUI{
		gui: g,
	}

	gui.gui.Cursor = true

	gui.gui.SetManager(&gui)

	gui.gui.InputEsc = true

	gui.gui.Update(func(gc *gocui.Gui) error {
		return gui.build()
	})

	if err := gui.keyBindings(); err != nil {
		log.Panicln(err)
	}

	return &gui
}



func (g *GUI) Start() {

	if err := g.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (g *GUI) Destroy()  {
	g.gui.Close()
}

func (g *GUI) AddCollection(collection *Collection)  {
	g.collection = collection
}

func (g *GUI ) build() error {
	view, err := g.gui.View("side")

	if err != nil {
		return err
	}

	for key, _ := range g.collection.ip {
		fmt.Fprintln(view, key)
	}

	return nil
}

func (gui *GUI) keyBindings() error {
	if err := gui.gui.SetKeybinding("side", gocui.KeyCtrlSpace, gocui.ModNone, guiNextView); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding("main", gocui.KeyCtrlSpace, gocui.ModNone, guiNextView); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding("side", gocui.KeyArrowDown, gocui.ModNone, guiCursorDown); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding("side", gocui.KeyArrowUp, gocui.ModNone, guiCursorUp); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, guiQuit); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding("side", gocui.KeyEnter, gocui.ModNone, guiGetLine); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, guiDeleteMessage); err != nil {
		return err
	}

	if err := gui.gui.SetKeybinding("main", gocui.KeyCtrlS, gocui.ModNone, guiSaveMain); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding("main", gocui.KeyCtrlW, gocui.ModNone, guiSaveVisualMain); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding("", gocui.KeyF2, gocui.ModNone, gui.changeGroup); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, gui.closeGroup); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding(VIEW_GROUP, gocui.KeyArrowDown, gocui.ModNone, gui.groupCursorDown); err != nil {
		return err
	}
	if err := gui.gui.SetKeybinding(VIEW_GROUP, gocui.KeyArrowUp, gocui.ModNone, gui.groupCursorUp); err != nil {
		return err
	}

	return nil
}

func (g *GUI) Layout(gui *gocui.Gui) error {

	maxX, maxY := gui.Size()
	if v, err := gui.SetView("side", -1, -1, 30, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Item 1")
		fmt.Fprintln(v, "Item 2")
		fmt.Fprintln(v, "Item 3")
		fmt.Fprint(v, "\rWill be")
		fmt.Fprint(v, "deleted\rItem 4\nItem 5")
	}

	if v, err := gui.SetView("main", 30, -1, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		b, err := ioutil.ReadFile("Mark.Twain-Tom.Sawyer.txt")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(v, "%s", b)
		v.Editable = true
		v.Wrap = true
		if _, err := gui.SetCurrentView("main"); err != nil {
			return err
		}
	}

	if v, err := gui.SetView("footer_help", -1, maxY - 2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		fmt.Fprint(v, "F1: Help")
	}
	if v, err := gui.SetView("footer_help", -1, maxY - 2, maxX/2, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		fmt.Fprint(v, "F1: Help")
	}
	if v, err := gui.SetView("footer_groups", maxX/2, maxY - 2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		fmt.Fprint(v, "F2: Groups: ")
	}

	return nil
}

func guiQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}


func guiNextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "side" {
		_, err := g.SetCurrentView("main")
		return err
	}
	_, err := g.SetCurrentView("side")
	return err
}

func guiCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func guiCursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}


func guiGetLine(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
		if _, err := g.SetCurrentView("msg"); err != nil {
			return err
		}
	}
	return nil
}

func guiDeleteMessage(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("msg"); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("side"); err != nil {
		return err
	}
	return nil
}

func guiSaveMain(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	p := make([]byte, 5)
	v.Rewind()
	for {
		n, err := v.Read(p)
		if n > 0 {
			if _, err := f.Write(p[:n]); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func guiSaveVisualMain(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	vb := v.ViewBuffer()
	if _, err := io.Copy(f, strings.NewReader(vb)); err != nil {
		return err
	}
	return nil
}

func (gui *GUI) changeGroup(g *gocui.Gui, v *gocui.View) error {

	maxX, maxY := g.Size()
	if v, err := g.SetView(VIEW_GROUP, maxX/2-20, maxY/2-10, maxX/2+20, maxY/2+10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Groups"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.BgColor = gocui.ColorCyan
		v.FgColor = gocui.ColorWhite
		fmt.Fprintln(v, "Time")
		fmt.Fprintln(v, "IP")
		fmt.Fprintln(v, "Token")
		fmt.Fprintln(v, "Type")
		fmt.Fprintln(v, "Application")

		if _, err := g.SetCurrentView(VIEW_GROUP); err != nil {
			return err
		}
	}
	return nil
}

func (gui *GUI) closeGroup(g *gocui.Gui, v *gocui.View) error {

	v, err := g.View(VIEW_GROUP)

	if err == nil {
		return g.DeleteView(VIEW_GROUP)
	} else if err != gocui.ErrUnknownView {
		return err
	}

	return nil
}

func (gui *GUI) groupCursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()

		l, err := v.Line(cy + 1)

		if l == "" {
			return nil
		}

		if err != nil {
			return err
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func (gui *GUI) groupCursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()

		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}