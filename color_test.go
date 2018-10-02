package color

import "testing"

func TestColor_Hex(t *testing.T) {
	cases := []struct {
		c   Color
		exp HexColor
	}{
		{
			c:   CMax,
			exp: HexColor("ffffff"),
		},
		{
			c:   CMin,
			exp: HexColor("000000"),
		},
		{
			c:   Color(0x808080),
			exp: HexColor("808080"),
		},
		{
			c:   Color(0xddeeff),
			exp: HexColor("ddeeff"),
		},
	}

	for _, c := range cases {
		act := c.c.Hex()

		if act != c.exp {
			t.Errorf("Expected %v.Hex() to be %s, %s given", c.c, c.exp, act)
		}
	}
}

func TestColor_RGBFloat(t *testing.T) {
	cases := []struct {
		c   Color
		exp RGBFloat
	}{
		{
			c:   CMax,
			exp: RGBFloat{1.0, 1.0, 1.0},
		},
		{
			c:   CMin,
			exp: RGBFloat{0.0, 0.0, 0.0},
		},
		{
			c:   Color(0x808080),
			exp: RGBFloat{0.5, 0.5, 0.5},
		},
		{
			c:   Color(0xaaaaaa),
			exp: RGBFloat{0.67, 0.67, 0.67},
		},
	}

	for _, c := range cases {
		act := c.c.RGBFloat()

		if !act.Equals(&c.exp) {
			t.Errorf("Expected %v.Color() to be %s, %s given", c.c, c.exp, act)
		}
	}
}

func TestRGBInt_Color(t *testing.T) {
	cases := []struct {
		c   RGBInt
		exp Color
	}{
		{
			c:   RGBInt{255, 255, 255},
			exp: CMax,
		},
		{
			c:   RGBInt{0, 0, 0},
			exp: CMin,
		},
		{
			c:   RGBInt{16, 16, 16},
			exp: Color(0x101010),
		},
	}

	for _, c := range cases {
		act := c.c.Color()

		if act != c.exp {
			t.Errorf("Expected %v.Color() to be %v, %v given", c.c, c.exp, act)
		}
	}
}

func TestRGBFloat_String(t *testing.T) {
	cases := []struct {
		c   RGBFloat
		exp string
	}{
		{
			c:   RGBFloat{1.0, 1.0, 1.0},
			exp: "rgb(1.0, 1.0, 1.0)",
		},
		{
			c:   RGBFloat{0.0, 0.0, 0.0},
			exp: "rgb(0.0, 0.0, 0.0)",
		},
		{
			c:   RGBFloat{0.67, 0.67, 0.67},
			exp: "rgb(0.67, 0.67, 0.67)",
		},
	}

	for _, c := range cases {
		act := c.c.String()

		if act != c.exp {
			t.Errorf("Expected %v.String() to be %s, %s given", c.c, c.exp, act)
		}
	}
}
