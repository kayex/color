package color

import (
	"fmt"
	"strings"
	"strconv"
)

// Color is a HTML color represented as single int value.
type Color uint

// CMin is the lowest color value.
const CMin Color = 0x0

// CMax is the highest color value.
const CMax Color = 0xffffff

func (c Color) String() string {
	return c.Hex().String()
}

func C(v int) (Color, error) {
	if v > int(CMax) {
		return 0, fmt.Errorf("%x exceeds color max value (%x)", v, CMax)
	}

	if v < int(CMin) {
		return 0, fmt.Errorf("invalid color value %x, color value must be positive", v)
	}

	return Color(v), nil
}

// Hex is a color represented as a hexadecimal string, prefixed with
// a single # sign.
//
// Hex only represents full hexadecimal color triplets (i.e. not colors on the
// shorthand hexadecimal form).
type Hex string

func (h *Hex) Color() Color {
	v := strings.TrimPrefix(h.String(), "#")
	c, err := strconv.ParseInt(v, 16, 32)
	if err != nil {
		panic(err)
	}

	return Color(c)
}

func (h Hex) String() string {
	return string(h)
}

func (c Color) Hex() Hex {
	rgb := c.RGB()
	h := fmt.Sprintf("#%02x%02x%02x", rgb.r, rgb.g, rgb.g)

	return Hex(h)
}

func (c Color) RGB() RGB {
	r := (c >> 16) & 255
	g := (c >> 8) & 255
	b := c & 255

	return RGB{uint8(r), uint8(g), uint8(b)}
}

// RGB is a color consisting of three 8 bit channels; red, green, and blue.
type RGB struct {
	r uint8
	g uint8
	b uint8
}

func (rgb *RGB) Color() Color {
	// Cast to uint to allow left shifting the red and green values.
	v := (uint(rgb.r) << 16) | (uint(rgb.g) << 8) | uint(rgb.b)

	return Color(v)
}

func (rgb RGB) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rgb.r, rgb.g, rgb.b)
}

// RGBA is an RGB color with an additional alpha (transparency) channel.
type RGBA struct {
	RGB
	// a is the color alpha channel between 0.0 and 1.0, where
	// 0.0 is completely transparent
	// 1.0 is completely opaque
	a float32
}

