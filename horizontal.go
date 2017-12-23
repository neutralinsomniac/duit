package duit

import (
	"image"

	"9fans.net/go/draw"
)

type Horizontal struct {
	Kids  []*Kid
	Split func(r image.Rectangle) (widths []int)

	size   image.Point
	widths []int
}

func (ui *Horizontal) Layout(env *Env, r image.Rectangle, cur image.Point) image.Point {
	r.Min = image.Pt(0, cur.Y)
	ui.widths = ui.Split(r)
	if len(ui.widths) != len(ui.Kids) {
		panic("bad number of widths from split")
	}
	ui.size = image.ZP
	for i, k := range ui.Kids {
		size := k.UI.Layout(env, image.Rectangle{image.ZP, image.Pt(ui.widths[i], r.Dy())}, image.ZP)
		p := image.Pt(ui.size.X, 0)
		k.r = image.Rectangle{p, p.Add(size)}
		ui.size.X += ui.widths[i]
		if k.r.Dy() > ui.size.Y {
			ui.size.Y = k.r.Dy()
		}
	}
	return ui.size
}

func (ui *Horizontal) Draw(env *Env, img *draw.Image, orig image.Point, m draw.Mouse) {
	kidsDraw(env, ui.Kids, ui.size, img, orig, m)
}

func (ui *Horizontal) Mouse(env *Env, m draw.Mouse) (result Result) {
	return kidsMouse(env, ui.Kids, m)
}

func (ui *Horizontal) Key(env *Env, orig image.Point, m draw.Mouse, k rune) (result Result) {
	return kidsKey(env, ui, ui.Kids, orig, m, k)
}

func (ui *Horizontal) FirstFocus(env *Env) *image.Point {
	return kidsFirstFocus(env, ui.Kids)
}

func (ui *Horizontal) Focus(env *Env, o UI) *image.Point {
	return kidsFocus(env, ui.Kids, o)
}

func (ui *Horizontal) Print(indent int, r image.Rectangle) {
	uiPrint("Horizontal", indent, r)
	kidsPrint(ui.Kids, indent+1)
}
