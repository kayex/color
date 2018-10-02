package color

import (
	"fmt"
	"math"
	"strconv"
)

// Color is a 24 bit sRGB web color.
type Color uint

// CMin is the lowest possible color value.
const CMin Color = 0x0

// CMax is the highest possible color value.
const CMax Color = 0xffffff

const (
	OffsetR = 0x10
	OffsetG = 0x08
	OffsetB = 0x00
)

func (c Color) Channels() (uint8, uint8, uint8) {
	r := uint8(c>>OffsetR) & 0xff
	g := uint8(c>>OffsetG) & 0xff
	b := uint8(c>>OffsetB) & 0xff

	return r, g, b
}

// HexColor is a Color represented as a hexadecimal string.
type HexColor string

// extendShortHex converts color values on the shorthand hexadecimal format
// to the full 6 character format.
//
// For example: fff -> ffffff
//              abc -> aabbcc
func extendShortHex(hex string) string {
	var l [6]byte

	for i := 0; i < 6; i++ {
		l[i] = hex[i/2]
	}

	return string(l[:])
}

func (h HexColor) String() string {
	return string("#" + h)
}

func (c Color) Hex() HexColor {
	rgb := c.RGBInt()
	h := fmt.Sprintf("%02x%02x%02x", rgb.r, rgb.g, rgb.b)

	return HexColor(h)
}

// RGBInt is a color represented by three 8 bit channel values.
type RGBInt struct {
	r uint8
	g uint8
	b uint8
}

func RGB(r, g, b uint8) Color {
	v := (uint(r) << OffsetR) | (uint(g) << OffsetG) | uint(b)<<OffsetB

	return Color(v)
}

func (c Color) RGBInt() RGBInt {
	r, g, b := c.Channels()

	return RGBInt{r, g, b}
}

func (rgb *RGBInt) Color() Color {
	c := RGB(rgb.r, rgb.g, rgb.b)

	return c
}

func (rgb RGBInt) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rgb.r, rgb.g, rgb.b)
}

// RGBFloat is a color represented by three float channel values between 0.0 and 1.0.
type RGBFloat struct {
	r float32
	g float32
	b float32
}

func (c Color) RGBFloat() RGBFloat {
	rgb := c.RGBInt()
	r := float32(rgb.r) / 0xff
	g := float32(rgb.g) / 0xff
	b := float32(rgb.b) / 0xff

	return RGBFloat{r, g, b}
}

func (rgb *RGBFloat) Color() Color {
	r := toIntChannelVal(rgb.r)
	g := toIntChannelVal(rgb.g)
	b := toIntChannelVal(rgb.b)

	return RGB(r, g, b)
}

func (rgb *RGBFloat) Equals(o *RGBFloat) bool {
	return colorEqual(rgb.r, o.r) &&
		colorEqual(rgb.g, o.g) &&
		colorEqual(rgb.b, o.b)
}

func (rgb RGBFloat) String() string {
	r := formatRGBChannelFloat(float64(rgb.r))
	g := formatRGBChannelFloat(float64(rgb.g))
	b := formatRGBChannelFloat(float64(rgb.b))

	return fmt.Sprintf("rgb(%v, %v, %v)", r, g, b)
}

func formatRGBChannelFloat(i float64) string {
	// First see if we have 2 or fewer significant decimal places,
	// and if so, return the number with up to 2 trailing 0s.
	if i*10 == math.Floor(i*10) {
		return strconv.FormatFloat(i, 'f', 1, 32)
	}
	// Otherwise, just format normally, using the minimum number of
	// necessary digits.
	return strconv.FormatFloat(i, 'f', 2, 32)
}

// colorEqual compares two sRGB float color values.
func colorEqual(a, b float32) bool {
	// p is the size of each distinct color value in 24 bit sRGB.
	const p = 1.0 / 0xff
	eq := (a-b) < p && (b-a) < p

	return eq
}

func toIntChannelVal(f float32) uint8 {
	var v uint8

	if v == 0.0 {
		v = 0
	} else {
		v = uint8(math.Round(255 / float64(f)))
	}

	return v
}
