// Copyright 2017 The gooid Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

var gFontRanges []uint16

func GetFontGlyphRanges(str string) *uint16 {
	ret := []uint16{0x0020, 0x00FF} // Basic Latin + Latin Supplement
	for _, r := range []rune(str) {
		ret = append(ret, uint16(r), uint16(r))
	}
	ret = append(ret, uint16(0), uint16(0))
	gFontRanges = ret
	return &gFontRanges[0]
}
