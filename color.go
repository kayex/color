package color

import "fmt"

// Hex is a color represented as single numeric value, usually written out
// in its hexadecimal form.
type Hex struct {
	// v is the color value from 0 to 16 777 215.
	v int
}

func (h *Hex) Hex() Hex {
	return *h
}

func (h *Hex) RGB() RGB {
	r := (h.v >> 16) & 255
	g := (h.v >> 8) & 255
	b := h.v & 255

	return RGB{uint8(r), uint8(g), uint8(b)}
}

func (h *Hex) String() string {
	return fmt.Sprintf("%x", h.v)
}

// RGB is a color consisting of three 8 bit channels; red, green, and blue.
type RGB struct {
	r uint8
	g uint8
	b uint8
}

func (rgb *RGB) RGB() RGB {
	return *rgb
}

func (rgb *RGB) Hex() Hex {
	v := rgb.b | (rgb.g << 8) | (rgb.r << 16)

	return Hex{int(v)}
}

// RGBA is an RGB color with an additional alpha (transparency) channel.
type RGBA struct {
	RGB
	// a is the color alpha channel, where
	// 0 is completely transparent
	// 1 is completely opaque
	a float32
}

type Color interface {
	Hex() Hex
	RGB() RGB
}

