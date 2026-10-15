package terminal

import (
	"testing"

	"fyne.io/fyne/v2"
	"github.com/stretchr/testify/assert"
)

func TestEncodeMouse(t *testing.T) {
	term := New()
	assert.Equal(t, "\x1b[M !!", string(term.encodeMouse(1, 0, fyne.NewPos(4, 4))))
	assert.Equal(t, "\x1b[M!$#", string(term.encodeMouse(2, 0, fyne.NewPos(30, 36))))
	assert.Equal(t, "\x1b[M#!!", string(term.encodeMouse(0, 0, fyne.NewPos(4, 4))))
}

func TestEncodeMouse_ScrollBack(t *testing.T) {
	term := New()
	cell := term.guessCellSize()
	term.Resize(fyne.NewSize(10*cell.Width, 5.5*cell.Height))
	term.handleOutput([]byte(numberedLines(10)))
	scroll := term.scrollContainer
	scroll.Refresh()
	scroll.ScrollToBottom()

	screenTop := float32(term.rowOffset())*cell.Height - scroll.Offset.Y
	assert.Equal(t, "\x1b[M !!", string(term.encodeMouse(1, 0, fyne.NewPos(1, screenTop+cell.Height*0.75))))

	scroll.ScrollToTop()
	assert.Equal(t, "\x1b[M !!", string(term.encodeMouse(1, 0, fyne.NewPos(1, 1))))
}

func TestEncodeMouse_Mods(t *testing.T) {
	term := New()
	assert.Equal(t, "\x1b[M$!!", string(term.encodeMouse(1,
		fyne.KeyModifierShift, fyne.NewPos(4, 4))))
	assert.Equal(t, "\x1b[M4!!", string(term.encodeMouse(1,
		fyne.KeyModifierShift|fyne.KeyModifierControl, fyne.NewPos(4, 4))))
	assert.Equal(t, "\x1b[M%!!", string(term.encodeMouse(2,
		fyne.KeyModifierShift, fyne.NewPos(4, 4))))
	assert.Equal(t, "\x1b[M5!!", string(term.encodeMouse(2,
		fyne.KeyModifierShift|fyne.KeyModifierControl, fyne.NewPos(4, 4))))
}
