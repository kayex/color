package color

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
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

func (c Color) String() string {
	return c.Hex().String()
}

// HexColor is a Color represented as a hexadecimal string.
type HexColor string

func Hex(s string) (Color, error) {
	h := strings.TrimPrefix(s, "#")
	if !validHex(h) {
		return 0, fmt.Errorf("invalid hex value %q", h)
	}

	// Here, we can safely assume that v is in the ASCII range (since it
	// passed validHex()) and index byte-wise.
	hl := len(h)
	switch hl {
	case 6:
	case 3:
		h = convertShortHex(h)
	default:
		return 0, fmt.Errorf("invalid hex format %q, hex colors should be either 3 or 6 characters long", h)
	}

	v, err := strconv.ParseInt(h, 16, 32)
	if err != nil {
		return 0, err
	}
	c := Color(v)

	if c < CMin || c > CMax {
		return 0, fmt.Errorf("invalid color value %q, values should be between %x and %x", h, CMin, CMax)
	}

	return c, nil
}

// convertShortHex converts color values on the shorthand hexadecimal format
// to the full 6 character format.
//
// For example: fff -> ffffff
//              abc -> aabbcc
func convertShortHex(hex string) string {
	var b bytes.Buffer

	b.WriteByte(hex[0])
	b.WriteByte(hex[0])
	b.WriteByte(hex[1])
	b.WriteByte(hex[1])
	b.WriteByte(hex[2])
	b.WriteByte(hex[2])

	return b.String()
}

func (h HexColor) String() string {
	return string("#" + h)
}

func (c Color) Hex() HexColor {
	rgb := c.RGBInt()
	h := fmt.Sprintf("%02x%02x%02x", rgb.r, rgb.g, rgb.g)

	return HexColor(h)
}

// RGBInt is a color represented by three 8 bit channels; red, green, and blue.
type RGBInt struct {
	r uint8
	g uint8
	b uint8
}

func (c Color) RGBInt() RGBInt {
	r := (c >> OffsetR) & 0xff
	g := (c >> OffsetG) & 0xff
	b := c & 0xff

	return RGBInt{uint8(r), uint8(g), uint8(b)}
}

func (rgb *RGBInt) Color() Color {
	c := consolidate(rgb.r, rgb.g, rgb.b)

	return c
}

func (rgb RGBInt) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rgb.r, rgb.g, rgb.b)
}

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
	c := consolidate(uint8(rgb.r), uint8(rgb.g), uint8(rgb.b))

	return c
}

func (rgb *RGBFloat) Equals(o *RGBFloat) bool {
	return colorChanEqual(rgb.r, o.r) &&
		colorChanEqual(rgb.g, o.g) &&
		colorChanEqual(rgb.b, o.b)
}

func (rgb RGBFloat) String() string {
	return fmt.Sprintf("rgb(%0.2f, %0.2f, %0.2f)", rgb.r, rgb.g, rgb.b)
}

func consolidate(r, g, b uint8) Color {
	v := (uint(r) << OffsetR) | (uint(g) << OffsetG) | uint(b)<<OffsetB

	return Color(v)
}

func validHex(hex string) bool {
	invalidChar := strings.IndexFunc(hex, func(r rune) bool {
		return !('0' <= r && r <= 'f')
	})

	return invalidChar == -1
}

func colorChanEqual(a, b float32) bool {
	// p is the size of each distinct color value in 24 bit sRGB.
	const p = 1.0 / 0xff

	eq := (a-b) < p && (b-a) < p

	return eq
}
