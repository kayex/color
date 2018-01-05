package color

import "testing"

func TestColor_String(t *testing.T) {
	cases := []struct{
		c Color
		exp string
	}{
		{
			c: CMax,
			exp: "#ffffff",
		},
		{
			c: CMin,
			exp: "#000000",
		},
		{
			c: Color(0x101010),
			exp: "#101010",
		},
	}

	for _, c := range cases {
		act := c.c.String()

		if act != c.exp {
			t.Errorf("Expected %v.String() to return %q, %q given", c.c, c.exp, act)
		}
	}
}

func TestHex_Color(t *testing.T) {
	cases := []struct{
		c Hex
		exp Color
	}{
		{
			c: Hex{"#ffffff"},
			exp: CMax,
		},
		{
			c: Hex{"#000000"},
			exp: CMin,
		},
		{
			c: Hex{"#101010"},
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

func TestRGB_Color(t *testing.T) {
	cases := []struct{
		c RGB
		exp Color
	}{
		{
			c: RGB{255, 255, 255},
			exp: CMax,
		},
		{
			c: RGB{0, 0, 0},
			exp: CMin,
		},
		{
			c: RGB{16, 16, 16},
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

