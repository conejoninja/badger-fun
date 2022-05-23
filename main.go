package main

import (
	"image/color"
	"machine"
	"math/rand"
	"time"

	"github.com/conejoninja/badger-fun/tetris"

	"github.com/conejoninja/badger-fun/fonts"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

var display uc8151.Device
var led machine.Pin

var black = color.RGBA{1, 1, 1, 255}
var white = color.RGBA{0, 0, 0, 255}
var w int16 = 296
var h int16 = 128

func main() {
	led = machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display = uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation: uc8151.ROTATION_270,
		Speed:    uc8151.TURBO,
		Blocking: true,
	})

	for {
		myNameIs("@_conejo")
		blinky("TECHNOLOGIST", "for hire")
		myNameIs("@_conejo")
		talkDate("Sunday, 9:30", "Gingerbread room")
		sunrays()
		loading()
		loadingInverted()
		tetrisfx()
		dvd("TinyGo")
	}
}

func myNameIs(name string) {
	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	var r int16 = 8

	// round corners
	tinydraw.FilledCircle(&display, r, r, r, black)
	tinydraw.FilledCircle(&display, w-r-1, r, r, black)
	tinydraw.FilledCircle(&display, r, h-r-1, r, black)
	tinydraw.FilledCircle(&display, w-r-1, h-r-1, r, black)

	// top band
	tinydraw.FilledRectangle(&display, r, 0, w-2*r-1, r, black)
	tinydraw.FilledRectangle(&display, 0, r, w, 26, black)

	// bottom band
	tinydraw.FilledRectangle(&display, r, h-r-1, w-2*r-1, r+1, black)
	tinydraw.FilledRectangle(&display, 0, h-2*r-1, w, r, black)

	// top text : my NAME is
	w32, _ := tinyfont.LineWidth(&fonts.Regular12pt7b, "my NAME is")
	tinyfont.WriteLine(&display, &fonts.Regular12pt7b, (w-int16(w32))/2, 24, "my NAME is", white)

	// middle text
	w32, _ = tinyfont.LineWidth(&fonts.Bold12pt7b, name)
	tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (w-int16(w32))/2, 74, name, black)

	tinyfont.WriteLine(&display, &fonts.Regular58pt, (w-int16(w32))/2-44, 86, "B", black)
	tinyfont.WriteLine(&display, &fonts.Regular58pt, w-40, 100, "F", black)

	display.Display()
	display.WaitUntilIdle()
	time.Sleep(50 * time.Second)
}

func blinky(topline, bottomline string) {
	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&fonts.Bold24pt7b, topline)
	w32bottom, _ := tinyfont.LineWidth(&fonts.Bold24pt7b, bottomline)
	for i := int16(0); i < 10; i++ {
		// fill the screen with with
		tinydraw.FilledRectangle(&display, 0, 0, w, h, white)
		// show black text
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32top))/2, 50, topline, black)
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32bottom))/2, 100, bottomline, black)

		// display
		display.Display()
		display.WaitUntilIdle()

		// repeat the other way around
		tinydraw.FilledRectangle(&display, 0, 0, w, h, black)
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32top))/2, 50, topline, white)
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32bottom))/2, 100, bottomline, white)

		display.Display()
		display.WaitUntilIdle()
	}
}

func talkDate(dateString, roomString string) {
	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	// top text : Come see my talk on
	_, w32 := tinyfont.LineWidth(&fonts.Bold12pt7b, "HackLab Workshop")
	tinyfont.WriteLine(&display, &fonts.Regular12pt7b, (w-int16(w32))/2, 28, "HackLab Workshop", black)

	// middle text
	w32, _ = tinyfont.LineWidth(&fonts.Bold12pt7b, dateString)
	tinyfont.WriteLine(&display, &fonts.Bold12pt7b, (w-int16(w32))/2, 70, dateString, black)

	if roomString != "" {
		// bottom text : at room XYZ
		w32, _ = tinyfont.LineWidth(&fonts.Regular12pt7b, roomString)
		tinyfont.WriteLine(&display, &fonts.Regular12pt7b, (w-int16(w32))/2, 110, roomString, black)
	}

	display.Display()
	display.WaitUntilIdle()
	time.Sleep(5 * time.Second)
}

func sunrays() {
	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	colors := [7][]color.RGBA{
		{white, white, white, white, white, white},
		{black, white, white, white, white, white},
		{black, white, white, black, white, white},
		{black, white, white, black, black, white},
		{black, black, white, black, black, white},
		{black, black, white, black, black, black},
		{black, black, black, black, black, black},
	}

	for i := 0; i < 21; i++ {
		if i%2 == 0 {
			tinydraw.FilledRectangle(&display, 0, 0, w, h, white)
			rays(black)
		} else {
			tinydraw.FilledRectangle(&display, 0, 0, w, h, black)
			rays(white)
		}
		tinydraw.FilledRectangle(&display, 20, 20, w-40, h-20, white)
		w32, _ := tinyfont.LineWidth(&fonts.Regular12pt7b, "Badge coded with")
		tinyfont.WriteLine(&display, &fonts.Regular12pt7b, (w-int16(w32))/2, 50, "Badge coded with", black)

		w32, _ = tinyfont.LineWidth(&fonts.Bold24pt7b, "TinyGo")
		tinyfont.WriteLineColors(&display, &fonts.Bold24pt7b, (w-int16(w32))/2, 100, "TinyGo", colors[i%7])

		display.Display()
		display.WaitUntilIdle()
	}

}

func rays(color color.RGBA) {
	// center point at the bottom
	var cx int16 = 124
	var cy int16 = 122

	// left side rays
	tinydraw.FilledTriangle(&display, cx, cy, 0, 0, 0, 21, color)
	tinydraw.FilledTriangle(&display, cx, cy, 0, 42, 0, 63, color)
	tinydraw.FilledTriangle(&display, cx, cy, 0, 84, 0, 105, color)

	// top rays
	tinydraw.FilledTriangle(&display, cx, cy, 23, 0, 47, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 71, 0, 94, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 118, 0, 142, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 166, 0, 189, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 213, 0, 236, 0, color)
	tinydraw.FilledTriangle(&display, cx, cy, 260, 0, 284, 0, color)

	// right rays
	tinydraw.FilledTriangle(&display, cx, cy, 296, 10, 296, 31, color)
	tinydraw.FilledTriangle(&display, cx, cy, 296, 52, 296, 73, color)
	tinydraw.FilledTriangle(&display, cx, cy, 296, 94, 296, 115, color)
}

func loading() {
	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	for i := int16(0); i < 37; i++ {
		// draw a rectangle bigger each time
		tinydraw.FilledRectangle(&display, i*8, 0, 10, h, black)

		// draw text again since a part of it was behind the rectangle
		w32, _ := tinyfont.LineWidth(&fonts.Bold24pt7b, "TinyGo")
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32))/2, 70, "TinyGo", white)

		display.Display()
		display.WaitUntilIdle()
	}
}

func loadingInverted() {
	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	for i := int16(0); i < 37; i++ {
		// this is the opposite, we draw the text and draw a rectangle of the same color as the background
		// to make the ilusion the text is revealing
		w32, _ := tinyfont.LineWidth(&fonts.Bold24pt7b, "TinyGo")
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32))/2, 70, "TinyGo", black)

		tinydraw.FilledRectangle(&display, i*8, 0, 296-i*8, h, white)

		display.Display()
		display.WaitUntilIdle()
	}
}

// tetrisfx
// will create a new ramdom piece ramdomly rotated each time the previous one stopped
func tetrisfx() {
	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	tetris.NewBoard()

	tetris.NewPiece()
	failed := 0
	k := 0
	for {
		display.ClearBuffer()
		if tetris.MovePiece() {
			failed = 0
		} else {
			failed++
			tetris.NewPiece()
		}
		tetris.DrawBoard(&display)
		tetris.DrawPiece(&display)

		w32, _ := tinyfont.LineWidth(&fonts.Bold24pt7b, "TinyGo")
		// add a white broder around the text
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32))/2-2, 70-2, "TinyGo", white)
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32))/2-2, 70+2, "TinyGo", white)
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32))/2+2, 70-2, "TinyGo", white)
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32))/2+2, 70+2, "TinyGo", white)
		// add the text
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, (w-int16(w32))/2, 70, "TinyGo", black)

		display.Display()
		display.WaitUntilIdle()
		// stop after 5 pieces in a row that can not move (screen is kind of filled)
		if failed >= 5 {
			return
		}
		k++
	}
}

func dvd(text string) {
	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	w32, _ := tinyfont.LineWidth(&fonts.Bold24pt7b, text)
	maxW := w - int16(w32)
	maxH := h - 36 //assume line height is 36

	// random start point
	x := int16(rand.Int31n(int32(maxW)))
	y := int16(rand.Int31n(int32(maxH)))
	d := int16(4)
	dx := d
	dy := d

	for i := 0; i < 80; i++ { //duration 80 frames
		// paint white the previous text to "erase" it
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, x, y+36, text, white)

		// move text
		x += dx
		y += dy

		// paint and show text
		tinyfont.WriteLine(&display, &fonts.Bold24pt7b, x, y+36, text, black)
		display.Display()
		display.WaitUntilIdle()

		// change direction if needed
		if x >= maxW {
			dx = -d
		}
		if x <= 0 {
			dx = d
		}
		if y >= maxH {
			dy = -d
		}
		if y <= 0 {
			dy = d
		}
	}
}
