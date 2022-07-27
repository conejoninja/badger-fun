/**
** The original 3x5 font is licensed under the 3-clause BSD license:
**
** Copyright 1999 Brian J. Swetland
** Copyright 1999 Vassilii Khachaturov
** Portions (of vt100.c/vt100.h) copyright Dan Marks
**
** All rights reserved.
**
** Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions
** are met:
** 1. Redistributions of source code must retain the above copyright
**    notice, this list of conditions, and the following disclaimer.
** 2. Redistributions in binary form must reproduce the above copyright
**    notice, this list of conditions, and the following disclaimer in the
**    documentation and/or other materials provided with the distribution.
** 3. The name of the authors may not be used to endorse or promote products
**    derived from this software without specific prior written permission.
**
** THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
** IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
** OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
** IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
** INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
** NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
** DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
** THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
** THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
**
** Modifications to Tom Thumb for improved readability are from Robey Pointer,
** see:
** http://robey.lag.net/2010/01/23/tiny-monospace-font.html
**
** The original author does not have any objection to relicensing of Robey
** Pointer's modifications (in this file) in a more permissive license.  See
** the discussion at the above blog, and also here:
** http://opengameart.org/forumtopic/how-to-submit-art-using-the-3-clause-bsd-license
**
** Feb 21, 2016: Conversion from Linux BDF --> Adafruit GFX font,
** with the help of this Python script:
** https://gist.github.com/skelliam/322d421f028545f16f6d
** William Skellenger (williamj@skellenger.net)
** Twitter: @skelliam
**
 */

package fonts

import "tinygo.org/x/tinyfont"

var TomThumb = tinyfont.Font{
	BBox: [4]int8{8, 6, 0, -5},
	Glyphs: []tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x0}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x80, 0x0, 0x80}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0xa0, 0xe0, 0xa0}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xc0, 0x40}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x20, 0x40, 0x80, 0x20}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0xe0, 0xa0, 0x60}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x8, Height: 0x2, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x80, 0x80, 0x40}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x40, 0x40, 0x80}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x40, 0xa0}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x40, 0xe0, 0x40}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x8, Height: 0x2, XAdvance: 0x3, XOffset: 0, YOffset: -2, Bitmaps: []uint8{0x40, 0x80}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0xe0}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x80}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x20, 0x40, 0x80, 0x80}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xa0, 0xa0, 0xc0}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xc0, 0x40, 0x40, 0x40}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x20, 0x40, 0x80, 0xe0}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x20, 0x40, 0x20, 0xc0}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0x20, 0x20}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0xc0, 0x20, 0xc0}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0xe0, 0xa0, 0xe0}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x40, 0x80, 0x80}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0xa0, 0xe0, 0xa0, 0xe0}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0xa0, 0xe0, 0x20, 0xc0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x8, Height: 0x3, XAdvance: 0x2, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x80, 0x0, 0x80}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x8, Height: 0x4, XAdvance: 0x3, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x40, 0x0, 0x40, 0x80}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0x80, 0x40, 0x20}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0x0, 0xe0}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x20, 0x40, 0x80}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x40, 0x0, 0x40}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xe0, 0x80, 0x60}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xe0, 0xa0, 0xa0}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xc0, 0xa0, 0xc0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x80, 0x60}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xa0, 0xa0, 0xc0}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0xe0, 0x80, 0xe0}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0xe0, 0x80, 0x80}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0xe0, 0xa0, 0x60}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0xa0, 0xa0}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x40, 0x40, 0x40, 0xe0}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x20, 0x20, 0xa0, 0x40}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xc0, 0xa0, 0xa0}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x80, 0x80, 0xe0}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0xe0, 0xa0, 0xa0}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0xe0, 0xe0, 0xa0}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xa0, 0xa0, 0x40}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xc0, 0x80, 0x80}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0xa0, 0xe0, 0x60}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xe0, 0xc0, 0xa0}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x40, 0x20, 0xc0}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x40, 0x40, 0x40, 0x40}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0xa0, 0x60}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0x40, 0x40}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0xe0, 0xa0}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0x40, 0xa0, 0xa0}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0x40, 0x40, 0x40}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x40, 0x80, 0xe0}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x80, 0x80, 0x80, 0xe0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x80, 0x40, 0x20}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x20, 0x20, 0x20, 0xe0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0xe0}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x8, Height: 0x2, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xc0, 0x60, 0xa0, 0xe0}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xc0, 0xa0, 0xa0, 0xc0}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x60}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x60, 0xa0, 0xa0, 0x60}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xa0, 0xc0, 0x60}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xe0, 0x40, 0x40}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xa0, 0xe0, 0x20, 0x40}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xc0, 0xa0, 0xa0, 0xa0}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x0, 0x80, 0x80, 0x80}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x0, 0x20, 0x20, 0xa0, 0x40}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xa0, 0xc0, 0xc0, 0xa0}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x40, 0x40, 0x40, 0xe0}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0xe0, 0xe0, 0xa0}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xc0, 0xa0, 0xa0, 0xa0}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x40, 0xa0, 0xa0, 0x40}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xc0, 0xa0, 0xa0, 0xc0, 0x80}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xa0, 0xa0, 0x60, 0x20}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x80}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xc0}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xe0, 0x40, 0x40, 0x60}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0x60}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xa0, 0xe0, 0x40}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xe0, 0xe0, 0xe0}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0x40, 0x40, 0xa0}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0xa0, 0x60, 0x20, 0x40}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0x60, 0xc0, 0xe0}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x40, 0x80, 0x40, 0x60}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0x80, 0x80}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x40, 0x20, 0x40, 0xc0}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0}},
		/*  */ tinyfont.Glyph{Rune: 127, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x0, 0x80, 0x80, 0x80}},
		/*  */ tinyfont.Glyph{Rune: 128, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xe0, 0x80, 0xe0, 0x40}},
		/*  */ tinyfont.Glyph{Rune: 129, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x40, 0xe0, 0x40, 0xe0}},
		/*  */ tinyfont.Glyph{Rune: 130, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x40, 0xe0, 0x40, 0xa0}},
		/*  */ tinyfont.Glyph{Rune: 131, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0x40, 0xe0, 0x40}},
		/*  */ tinyfont.Glyph{Rune: 132, Width: 0x8, Height: 0x5, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0x80, 0x80}},
		/*  */ tinyfont.Glyph{Rune: 133, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x40, 0xa0, 0x40, 0xc0}},
		/*  */ tinyfont.Glyph{Rune: 134, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0}},
		/*  */ tinyfont.Glyph{Rune: 135, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x60}},
		/*  */ tinyfont.Glyph{Rune: 136, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xe0, 0x0, 0xe0}},
		/*  */ tinyfont.Glyph{Rune: 137, Width: 0x8, Height: 0x3, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40}},
		/*  */ tinyfont.Glyph{Rune: 138, Width: 0x8, Height: 0x2, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0x20}},
		/*  */ tinyfont.Glyph{Rune: 139, Width: 0x8, Height: 0x1, XAdvance: 0x3, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0xc0}},
		/*  */ tinyfont.Glyph{Rune: 140, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0xa0}},
		/*  */ tinyfont.Glyph{Rune: 141, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0}},
		/*  */ tinyfont.Glyph{Rune: 142, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0x40}},
		/*  */ tinyfont.Glyph{Rune: 143, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xe0, 0x40, 0x0, 0xe0}},
		/*  */ tinyfont.Glyph{Rune: 144, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x40, 0x60}},
		/*  */ tinyfont.Glyph{Rune: 145, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x60, 0xe0}},
		/*  */ tinyfont.Glyph{Rune: 146, Width: 0x8, Height: 0x2, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80}},
		/*  */ tinyfont.Glyph{Rune: 147, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xa0, 0xa0, 0xc0, 0x80}},
		/*  */ tinyfont.Glyph{Rune: 148, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0x60, 0x60, 0x60}},
		/*  */ tinyfont.Glyph{Rune: 149, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xe0, 0xe0, 0xe0}},
		/*  */ tinyfont.Glyph{Rune: 150, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0x40, 0x20, 0xc0}},
		/*  */ tinyfont.Glyph{Rune: 151, Width: 0x8, Height: 0x3, XAdvance: 0x2, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x80}},
		/*  */ tinyfont.Glyph{Rune: 152, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0xa0, 0x40, 0x0, 0xe0}},
		/*  */ tinyfont.Glyph{Rune: 153, Width: 0x8, Height: 0x3, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x80}},
		/*  */ tinyfont.Glyph{Rune: 154, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0x60, 0x20}},
		/*  */ tinyfont.Glyph{Rune: 155, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x80, 0x0, 0xc0, 0x60}},
		/*  */ tinyfont.Glyph{Rune: 156, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0x0, 0x60, 0x20}},
		/*  */ tinyfont.Glyph{Rune: 157, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x0, 0x40, 0x80, 0xe0}},
		/*  */ tinyfont.Glyph{Rune: 158, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x40, 0xe0, 0xa0}},
		/*  */ tinyfont.Glyph{Rune: 159, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40, 0xe0, 0xa0}},
		/*   */ tinyfont.Glyph{Rune: 160, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x40, 0xe0, 0xa0}},
		/* ¡ */ tinyfont.Glyph{Rune: 161, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x40, 0xe0, 0xa0}},
		/* ¢ */ tinyfont.Glyph{Rune: 162, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x40, 0xa0, 0xe0, 0xa0}},
		/* £ */ tinyfont.Glyph{Rune: 163, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xc0, 0xa0, 0xe0, 0xa0}},
		/* ¤ */ tinyfont.Glyph{Rune: 164, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0xe0, 0xc0, 0xe0}},
		/* ¥ */ tinyfont.Glyph{Rune: 165, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x80, 0x80, 0x60, 0x20, 0x40}},
		/* ¦ */ tinyfont.Glyph{Rune: 166, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0xe0, 0xc0, 0xe0}},
		/* § */ tinyfont.Glyph{Rune: 167, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0xe0, 0xc0, 0xe0}},
		/* ¨ */ tinyfont.Glyph{Rune: 168, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xe0, 0xc0, 0xe0}},
		/* © */ tinyfont.Glyph{Rune: 169, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xe0, 0xc0, 0xe0}},
		/* ª */ tinyfont.Glyph{Rune: 170, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0xe0, 0x40, 0xe0}},
		/* « */ tinyfont.Glyph{Rune: 171, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0xe0, 0x40, 0xe0}},
		/* ¬ */ tinyfont.Glyph{Rune: 172, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xe0, 0x40, 0xe0}},
		/* ­ */ tinyfont.Glyph{Rune: 173, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xe0, 0x40, 0xe0}},
		/* ® */ tinyfont.Glyph{Rune: 174, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0xa0, 0xe0, 0xa0, 0xc0}},
		/* ¯ */ tinyfont.Glyph{Rune: 175, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0xa0, 0xe0, 0xa0}},
		/* ° */ tinyfont.Glyph{Rune: 176, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0xe0, 0xa0, 0xe0}},
		/* ± */ tinyfont.Glyph{Rune: 177, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0xe0, 0xa0, 0xe0}},
		/* ² */ tinyfont.Glyph{Rune: 178, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xe0, 0xa0, 0xe0}},
		/* ³ */ tinyfont.Glyph{Rune: 179, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0xe0, 0xa0, 0xe0}},
		/* ´ */ tinyfont.Glyph{Rune: 180, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xe0, 0xa0, 0xe0}},
		/* µ */ tinyfont.Glyph{Rune: 181, Width: 0x8, Height: 0x3, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xa0, 0x40, 0xa0}},
		/* ¶ */ tinyfont.Glyph{Rune: 182, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xe0, 0xa0, 0xc0}},
		/* · */ tinyfont.Glyph{Rune: 183, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0xa0, 0xa0, 0xe0}},
		/* ¸ */ tinyfont.Glyph{Rune: 184, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0xa0, 0xe0}},
		/* ¹ */ tinyfont.Glyph{Rune: 185, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xa0, 0xa0, 0xe0}},
		/* º */ tinyfont.Glyph{Rune: 186, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0xa0, 0xe0}},
		/* » */ tinyfont.Glyph{Rune: 187, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0xe0, 0x40}},
		/* ¼ */ tinyfont.Glyph{Rune: 188, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0xe0, 0xa0, 0xe0, 0x80}},
		/* ½ */ tinyfont.Glyph{Rune: 189, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xa0, 0xc0, 0xa0, 0xc0, 0x80}},
		/* ¾ */ tinyfont.Glyph{Rune: 190, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x60, 0xa0, 0xe0}},
		/* ¿ */ tinyfont.Glyph{Rune: 191, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x60, 0xa0, 0xe0}},
		/* À */ tinyfont.Glyph{Rune: 192, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x60, 0xa0, 0xe0}},
		/* Á */ tinyfont.Glyph{Rune: 193, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xa0, 0xe0}},
		/* Â */ tinyfont.Glyph{Rune: 194, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x60, 0xa0, 0xe0}},
		/* Ã */ tinyfont.Glyph{Rune: 195, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0x60, 0x60, 0xa0, 0xe0}},
		/* Ä */ tinyfont.Glyph{Rune: 196, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xe0, 0xe0, 0xc0}},
		/* Å */ tinyfont.Glyph{Rune: 197, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0x80, 0x60, 0x20, 0x40}},
		/* Æ */ tinyfont.Glyph{Rune: 198, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x60, 0xe0, 0x60}},
		/* Ç */ tinyfont.Glyph{Rune: 199, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x60, 0xe0, 0x60}},
		/* È */ tinyfont.Glyph{Rune: 200, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x60, 0xe0, 0x60}},
		/* É */ tinyfont.Glyph{Rune: 201, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x60, 0xe0, 0x60}},
		/* Ê */ tinyfont.Glyph{Rune: 202, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0x80, 0x80, 0x80}},
		/* Ë */ tinyfont.Glyph{Rune: 203, Width: 0x8, Height: 0x5, XAdvance: 0x3, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40, 0x40, 0x40}},
		/* Ì */ tinyfont.Glyph{Rune: 204, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x40, 0x40, 0x40}},
		/* Í */ tinyfont.Glyph{Rune: 205, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x40, 0x40, 0x40}},
		/* Î */ tinyfont.Glyph{Rune: 206, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0x60, 0xa0, 0x60}},
		/* Ï */ tinyfont.Glyph{Rune: 207, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0xc0, 0xa0, 0xa0}},
		/* Ð */ tinyfont.Glyph{Rune: 208, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x20, 0x40, 0xa0, 0x40}},
		/* Ñ */ tinyfont.Glyph{Rune: 209, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x80, 0x40, 0xa0, 0x40}},
		/* Ò */ tinyfont.Glyph{Rune: 210, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0x40, 0xa0, 0x40}},
		/* Ó */ tinyfont.Glyph{Rune: 211, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xc0, 0x60, 0x40, 0xa0, 0x40}},
		/* Ô */ tinyfont.Glyph{Rune: 212, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0x40, 0xa0, 0x40}},
		/* Õ */ tinyfont.Glyph{Rune: 213, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x40, 0x0, 0xe0, 0x0, 0x40}},
		/* Ö */ tinyfont.Glyph{Rune: 214, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xe0, 0xa0, 0xc0}},
		/* × */ tinyfont.Glyph{Rune: 215, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x80, 0x40, 0xa0, 0xa0, 0x60}},
		/* Ø */ tinyfont.Glyph{Rune: 216, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0xa0, 0x60}},
		/* Ù */ tinyfont.Glyph{Rune: 217, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0x0, 0xa0, 0xa0, 0x60}},
		/* Ú */ tinyfont.Glyph{Rune: 218, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0xa0, 0x60}},
		/* Û */ tinyfont.Glyph{Rune: 219, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x20, 0x40, 0xa0, 0x60, 0x20, 0x40}},
		/* Ü */ tinyfont.Glyph{Rune: 220, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x80, 0xc0, 0xa0, 0xc0, 0x80}},
		/* Ý */ tinyfont.Glyph{Rune: 221, Width: 0x8, Height: 0x6, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0x60, 0x20, 0x40}},
		/* Þ */ tinyfont.Glyph{Rune: 222, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x0}},
		/* ß */ tinyfont.Glyph{Rune: 223, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xc0, 0xe0, 0xc0, 0x60}},
		/* à */ tinyfont.Glyph{Rune: 224, Width: 0x8, Height: 0x4, XAdvance: 0x4, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x60, 0xe0, 0xc0, 0xe0}},
		/* á */ tinyfont.Glyph{Rune: 225, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x60, 0xc0, 0x60, 0xc0}},
		/* â */ tinyfont.Glyph{Rune: 226, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x60, 0xc0, 0x60, 0xc0}},
		/* ã */ tinyfont.Glyph{Rune: 227, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0x0, 0xa0, 0x40, 0x40}},
		/* ä */ tinyfont.Glyph{Rune: 228, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0x60, 0xc0, 0xe0}},
		/* å */ tinyfont.Glyph{Rune: 229, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xa0, 0xe0, 0x60, 0xc0, 0xe0}},
		/* æ */ tinyfont.Glyph{Rune: 230, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x0}},
		/* ç */ tinyfont.Glyph{Rune: 231, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0x0}},
		/* è */ tinyfont.Glyph{Rune: 232, Width: 0x8, Height: 0x1, XAdvance: 0x2, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0x80}},
		/* é */ tinyfont.Glyph{Rune: 233, Width: 0x8, Height: 0x1, XAdvance: 0x4, XOffset: 0, YOffset: -1, Bitmaps: []uint8{0xa0}},
		/* ê */ tinyfont.Glyph{Rune: 234, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x60, 0xe0, 0xe0, 0xc0, 0x60}},
		/* ë */ tinyfont.Glyph{Rune: 235, Width: 0x8, Height: 0x5, XAdvance: 0x4, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0xe0, 0xa0, 0xa0, 0xa0, 0xe0}},
	},

	YAdvance: 0x6,
}
