package color

import (
	"fmt"
	"strings"
	"strconv"
	"bytes"
)

// Color is a 24 bit sRGB web color.
type Color uint

// CMin is the lowest possible color value.
const CMin Color = 0x0

// CMax is the highest possible color value.
const CMax Color = 0xffffff

func (c Color) String() string {
	return c.Hex().String()
}

// HexColor is a Color represented as a hexadecimal string, with a
// #-sign prefixed.
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
// For example: #fff -> #ffffff
//              #abc -> #aabbcc
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

func validHex(hex string) bool {
	invalidChar := strings.IndexFunc(hex, func(r rune) bool {
		return !('0' <= r && r <= 'f')
	})

	return invalidChar == -1
}

func (h HexColor) String() string {
	return string(h)
}

func (c Color) Hex() HexColor {
	rgb := c.RGB()
	h := fmt.Sprintf("#%02x%02x%02x", rgb.r, rgb.g, rgb.g)

	return HexColor(h)
}

func (c Color) RGB() RGBColor {
	r := (c >> 16) & 255
	g := (c >> 8) & 255
	b := c & 255

	return RGBColor{uint8(r), uint8(g), uint8(b)}
}

// RGBColor is a color represented by three 8 bit channels; red, green, and blue.
type RGBColor struct {
	r uint8
	g uint8
	b uint8
}

func RGB(r, g, b uint8) RGBColor {
	return RGBColor{r, g, b}
}

func (rgb *RGBColor) Color() Color {
	// Cast to full-size uint to allow left shifting the red and green values.
	v := (uint(rgb.r) << 16) | (uint(rgb.g) << 8) | uint(rgb.b)

	return Color(v)
}

func (rgb RGBColor) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rgb.r, rgb.g, rgb.b)
}

// RGBA is an RGBColor color with an additional alpha (transparency) channel.
type RGBA struct {
	RGBColor
	// a is the color alpha channel. It assumes values between 0.0 and 1.0
	// where 0.0 represents complete transparency.
	a float32
}

