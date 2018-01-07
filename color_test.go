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

func TestHex(t *testing.T) {
	cases := []struct{
		c   string
		exp Color
		err bool
		desc string
	}{
		{
			c:   "#ffffff",
			exp: CMax,
		},
		{
			c:   "#000000",
			exp: CMin,
		},
		{
			c: "#000",
			exp: CMin,

		},
		{
			c: "#fff",
			exp: CMax,

		},
		{
			c:   "#101010",
			exp: Color(0x101010),
		},
		{
			c: "#1000000",
			err: true,
			desc: "color value exceeds CMax",
		},
		{
			c: "#aaff",
			err: true,
			desc: "invalid color format",
		},
	}

	for _, c := range cases {
		act, err := Hex(c.c)

		if err != nil && !c.err {
			t.Errorf("Expected Hex(%q) to be %v, error given: %v", c.c, c.exp, err)
			continue
		} else if err == nil && c.err {
			t.Errorf("Expected Hex(%q) to return error: %v", c.c, c.desc)
			continue
		}

		if act != c.exp {
			t.Errorf("Expected %v.Color() to be %v, %v given", c.c, c.exp, act)
		}
	}
}

func TestRGB_Color(t *testing.T) {
	cases := []struct{
		c   RGBColor
		exp Color
	}{
		{
			c:   RGBColor{255, 255, 255},
			exp: CMax,
		},
		{
			c:   RGBColor{0, 0, 0},
			exp: CMin,
		},
		{
			c:   RGBColor{16, 16, 16},
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

