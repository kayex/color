package color

import (
	"fmt"
	"strings"
	"strconv"
)

// Color is a HTML color represented as single int value.
type Color uint

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

// CMin is the lowest color value.
const CMin Color = 0x0

// CMax is the highest color value.
const CMax Color = 0xffffff

// Hex is a color represented as a hexadecimal string, prefixed with
// a single # sign.
//
// Hex only represents full hexadecimal color triplets (i.e. not colors on the
// shorthand hexadecimal form).
type Hex struct {
	// v is the hexadecimal color string.
	v string
}

func (h *Hex) Color() Color {
	v := strings.TrimPrefix(h.v, "#")
	c, err := strconv.ParseInt(v, 16, 32)
	if err != nil {
		panic(err)
	}
	return Color(c)
}

func (h Hex) String() string {
	return fmt.Sprintf("#%s", h.v)
}

func (c Color) Hex() Hex {
	rgb := c.RGB()
	h := fmt.Sprintf("%02x%02x%02x", uint8(rgb.r), rgb.g, rgb.g)

	return Hex{h}
}

func (c Color) RGB() RGB {
	r := (c >> 16) & 255
	g := (c >> 8) & 255
	b := c & 255

	return RGB{uint(r), uint(g), uint(b)}
}

// RGB is a color consisting of three 8 bit channels; red, green, and blue.
type RGB struct {
	r uint
	g uint
	b uint
}

func (rgb *RGB) Color() Color {
	v := (rgb.r << 16) | (rgb.g << 8) | rgb.b

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

type Format interface {
	Color() int
}
