package color

import (
	"fmt"
	"strings"
	"strconv"
)

// Color is a 24 bit sRGB web color.
type Color uint

// CMin is the lowest color value.
const CMin Color = 0x0

// CMax is the highest color value.
const CMax Color = 0xffffff

func (c Color) String() string {
	return c.Hex().String()
}

// HexColor is a color represented as a hexadecimal string with a #-sign prefixed,
// for example #ff0023.
//
// HexColor only represents full hexadecimal color triplets. This means colors on the
// shorthand hexadecimal form need to be converted to the full format before
// being used as HexColor.
type HexColor string

func Hex(s string) (Color, error) {
	v := strings.TrimPrefix(s, "#")
	c, err := strconv.ParseInt(v, 16, 32)
	if err != nil {
		return 0, err
	}

	return Color(c), nil
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

