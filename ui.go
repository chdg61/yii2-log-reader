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
	VIEW_GROUP        = "group"
	VIEW_LEFT         = "left"
	VIEW_MAIN         = "main"
	VIEW_FOOTER_HELP  = "footer_help"
	VIEW_FOOTER_GROUP = "footer_group"
)

const (
	GROUP_IP   = "ip"
	GROUP_TIME = "time"
	GROUP_TYPE = "type"
)

type UI struct {
	gui         *gocui.Gui
	collection  *Collection
	selectGroup string
}

func NewUI() *UI {

	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		log.Panicln(err)
	}

	ui := UI{
		gui: g,
	}

	ui.gui.Cursor = true

	ui.gui.SetManager(&ui)

	ui.gui.InputEsc = true

	ui.gui.Update(func(gc *gocui.Gui) error {
		ui.selectGroup = GROUP_IP
		return ui.build()
	})

	if err := ui.keyBindings(); err != nil {
		log.Panicln(err)
	}

	return &ui
}

func (u *UI) Start() {

	if err := u.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (u *UI) Destroy() {
	u.gui.Close()
}

func (u *UI) AddCollection(collection *Collection) {
	u.collection = collection
}

func (u *UI) build() error {

	viewLeft, err := u.gui.View(VIEW_LEFT)

	if err != nil {
		return err
	}
	var viewMain *gocui.View
	viewMain, err = u.gui.View(VIEW_MAIN)

	if err != nil {
		return err
	}
	firstView := false;

	var group GroupCollection

	switch u.selectGroup {
		case GROUP_IP:
			group = u.collection.ip
		case GROUP_TIME:
			group = u.collection.time
		case GROUP_TYPE:
			group = u.collection.chunkType
	}

	viewLeft.Clear()
	viewMain.Clear()

	group.EachCollection(func(key fmt.Stringer, chunkList *[]Chunk) bool {
		fmt.Fprintln(viewLeft, key)
		if !firstView {
			for _, chunk := range *chunkList {
				fmt.Fprintln(viewMain, chunk.text)
			}

			firstView = true
		}

		return true
	})
	//for key, chunkList := range u.collection.ip {
	//	fmt.Fprintln(viewLeft, key)
	//	if !firstView {
	//		for _, chunk := range chunkList {
	//			fmt.Fprintln(viewMain, chunk.text)
	//		}
	//
	//		firstView = true
	//	}
	//}

	return nil
}

func (u *UI) keyBindings() error {
	if err := u.gui.SetKeybinding(VIEW_LEFT, gocui.KeyCtrlSpace, gocui.ModNone, guiNextView); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding(VIEW_MAIN, gocui.KeyCtrlSpace, gocui.ModNone, guiNextView); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding(VIEW_LEFT, gocui.KeyArrowDown, gocui.ModNone, guiCursorDown); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding(VIEW_LEFT, gocui.KeyArrowUp, gocui.ModNone, guiCursorUp); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, guiQuit); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding(VIEW_LEFT, gocui.KeyEnter, gocui.ModNone, guiGetLine); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, guiDeleteMessage); err != nil {
		return err
	}

	if err := u.gui.SetKeybinding(VIEW_MAIN, gocui.KeyCtrlS, gocui.ModNone, guiSaveMain); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding(VIEW_MAIN, gocui.KeyCtrlW, gocui.ModNone, guiSaveVisualMain); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding("", gocui.KeyF2, gocui.ModNone, u.changeGroup); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, u.closeGroup); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding(VIEW_GROUP, gocui.KeyArrowDown, gocui.ModNone, u.groupCursorDown); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding(VIEW_GROUP, gocui.KeyArrowUp, gocui.ModNone, u.groupCursorUp); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding(VIEW_GROUP, gocui.KeyEnter, gocui.ModNone, u.groupEnter); err != nil {
		return err
	}

	return nil
}

func (u *UI) Layout(gui *gocui.Gui) error {

	maxX, maxY := gui.Size()
	if v, err := gui.SetView(VIEW_LEFT, -1, -1, 30, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
	}

	if v, err := gui.SetView(VIEW_MAIN, 30, -1, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//b, err := ioutil.ReadFile("Mark.Twain-Tom.Sawyer.txt")
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Fprintf(v, "%s", b)
		v.Editable = true
		v.Wrap = true
		if _, err := gui.SetCurrentView(VIEW_MAIN); err != nil {
			return err
		}
	}

	if v, err := gui.SetView(VIEW_FOOTER_HELP, -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		fmt.Fprint(v, "F1: Help")
	}
	if v, err := gui.SetView(VIEW_FOOTER_GROUP, maxX/2, maxY-2, maxX, maxY); err != nil {
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
	if v == nil || v.Name() == VIEW_LEFT {
		_, err := g.SetCurrentView(VIEW_MAIN)
		return err
	}
	_, err := g.SetCurrentView(VIEW_LEFT)
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
	if _, err := g.SetCurrentView(VIEW_LEFT); err != nil {
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

func (u *UI) changeGroup(g *gocui.Gui, v *gocui.View) error {

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

		cursorIndex := 0

		for i := 0; ; i++ {
			l, _ := v.Line(i)
			if l == "" {
				break
			}
			switch {
			case u.selectGroup == GROUP_IP && l == "IP":
				cursorIndex = i
				break
			case u.selectGroup == GROUP_TIME && l == "Time":
				cursorIndex = i
				break
			case u.selectGroup == GROUP_TYPE && l == "Type":
				cursorIndex = i
				break
			}
		}

		if err := v.SetCursor(0, cursorIndex); err != nil {
			return err
		}

		if _, err := g.SetCurrentView(VIEW_GROUP); err != nil {
			return err
		}
	}
	return nil
}

func (u *UI) closeGroup(g *gocui.Gui, v *gocui.View) error {

	v, err := g.View(VIEW_GROUP)

	if err == nil {
		if _, err := g.SetCurrentView(VIEW_MAIN); err != nil {
			return err
		}
		return g.DeleteView(VIEW_GROUP)
	} else if err != gocui.ErrUnknownView {
		return err
	}

	return nil
}

func (u *UI) groupCursorDown(g *gocui.Gui, v *gocui.View) error {
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

func (u *UI) groupCursorUp(g *gocui.Gui, v *gocui.View) error {
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

func (u *UI) groupEnter(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, cy := v.Cursor()

		l, err := v.Line(cy)

		if l == "" {
			return nil
		}

		if err != nil {
			return err
		}

		switch l {
		case "IP":
			u.selectGroup = GROUP_IP
		case "Time":
			u.selectGroup = GROUP_TIME
		case "Type":
			u.selectGroup = GROUP_TYPE
		}

		u.build()

		if _, err := g.SetCurrentView(VIEW_MAIN); err != nil {
			return err
		}
		g.DeleteView(VIEW_GROUP)

	}
	return nil
}
