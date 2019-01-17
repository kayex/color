package color

import (
	"fmt"
	"math"
	"strconv"
)

// Color is a 24 bit sRGB web color.
type Color uint

func (c Color) Color() Color {
	return c
}

func (c Color) String() string {
	return fmt.Sprintf("sRGB(%d)", c)
}

// AlphaColor is a color with a 32 bit alpha channel.
type AlphaColor struct {
	Color
	Alpha float32
}

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

type Format interface {
	Color() Color
	String() string
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
	h := fmt.Sprintf("%02x%02x%02x", rgb.R, rgb.G, rgb.B)

	return HexColor(h)
}

func (h HexColor) Color() Color {
	v, err := strconv.ParseInt(string(h), 16, 32)
	if err != nil {
		panic(err)
	}

	c := Color(v)
	return c
}

// RGBInt is a color represented by three 8 bit channel values.
type RGBInt struct {
	R uint8
	G uint8
	B uint8
}

func RGB(r, g, b uint8) Color {
	v := (uint(r) << OffsetR) | (uint(g) << OffsetG) | uint(b)<<OffsetB

	return Color(v)
}

func (c Color) RGBInt() RGBInt {
	r, g, b := c.Channels()

	return RGBInt{r, g, b}
}

func (rgb RGBInt) Color() Color {
	c := RGB(rgb.R, rgb.G, rgb.B)

	return c
}

func (rgb RGBInt) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", rgb.R, rgb.G, rgb.B)
}

type RGBAInt struct {
	RGBInt
	A float32
}

func RGBA(r, g, b uint8, a float32) AlphaColor {
	v := (uint(r) << OffsetR) | (uint(g) << OffsetG) | uint(b)<<OffsetB

	return AlphaColor{Color(v), a}
}

func (rgba RGBAInt) AlphaColor() AlphaColor {
	c := AlphaColor{rgba.Color(), rgba.A}

	return c
}

func (rgba RGBAInt) String() string {
	return fmt.Sprintf("rgba(%d, %d, %d, %s)", rgba.R, rgba.G, rgba.B, FormatRGBChannelFloat(float64(rgba.A)))
}

// RGBFloat is a color represented by three float channel values between 0.0 and 1.0.
type RGBFloat struct {
	R float32
	G float32
	B float32
}

func (c Color) RGBFloat() RGBFloat {
	rgb := c.RGBInt()
	r := float32(rgb.R) / 0xff
	g := float32(rgb.G) / 0xff
	b := float32(rgb.B) / 0xff

	return RGBFloat{r, g, b}
}

func (rgb RGBFloat) Color() Color {
	r := IntChannelValue(float64(rgb.R))
	g := IntChannelValue(float64(rgb.G))
	b := IntChannelValue(float64(rgb.B))

	return RGB(r, g, b)
}

func (rgb RGBFloat) Equals(o *RGBFloat) bool {
	return channelEqual(rgb.R, o.R) &&
		channelEqual(rgb.G, o.G) &&
		channelEqual(rgb.B, o.B)
}

func (rgb RGBFloat) String() string {
	r, g, b := rgb.Formatted()

	return fmt.Sprintf("rgb(%s, %s, %s)", r, g, b)
}

func (rgb RGBFloat) Formatted() (string, string, string) {
	r := FormatRGBChannelFloat(float64(rgb.R))
	g := FormatRGBChannelFloat(float64(rgb.G))
	b := FormatRGBChannelFloat(float64(rgb.B))

	return r, g, b
}

type RGBAFloat struct {
	RGBFloat
	A float32
}

func (rgba RGBAFloat) AlphaColor() AlphaColor {
	c := AlphaColor{rgba.Color(), rgba.A}

	return c
}

func (rgba RGBAFloat) String() string {
	r, g, b, a := rgba.Formatted()

	return fmt.Sprintf("rgba(%s, %s, %s, %s)", r, g, b, a)
}

func (rgba RGBAFloat) Formatted() (string, string, string, string) {
	r, g, b := rgba.RGBFloat.Formatted()
	a := FormatRGBChannelFloat(float64(rgba.A))

	return r, g, b, a
}

// FormatRGBChannelFloat formats a single float channel value for display purposes.
func FormatRGBChannelFloat(i float64) string {
	// If decimals are not significant, truncate to one decimal point.
	if i*10 == math.Floor(i*10) {
		return strconv.FormatFloat(i, 'f', 1, 32)
	}

	// Otherwise, use at most two decimal points.
	return strconv.FormatFloat(i, 'f', 2, 32)
}

// channelEqual compares two sRGB float color channels.
func channelEqual(a, b float32) bool {
	// p is the size of each distinct channel value in 24 bit sRGB.
	const p = 1.0 / 0xff
	eq := (a-b) < p && (b-a) < p

	return eq
}

func IntChannelValue(f float64) uint8 {
	var v uint8

	if f == 0.0 {
		v = 0
	} else {
		v = uint8(math.Round(0xff / f))
	}

	return v
}
