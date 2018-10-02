package color

import "testing"

func TestParse(t *testing.T) {
	var cases = []struct {
		color string
		exp   Color
		err   bool
	}{
		{"rgb(0,0,0)", RGB(0, 0, 0), false},
		{"rgb(255,255,255)", RGB(255, 255, 255), false},
		{"rgb(255, 255, 255)", RGB(255, 255, 255), false},
		{"rgb(255, 255, 255, 0.5)", RGB(255, 255, 255), false},
		{"rgb(255,255,255,0.5)", RGB(255, 255, 255), false},
		{"rgb(1.0,1.0,1.0)", RGB(255, 255, 255), false},
		{"rgb(1.0, 1.0, 1.0)", RGB(255, 255, 255), false},
		{"rgba(1.0, 1.0, 1.0, 0.5)", RGB(255, 255, 255), false},
		{"rgba(0.0, 0.0, 0.0, 0.0)", RGB(0, 0, 0), false},
		{"rgba(1.0, 1.0, 1.0, 1.0)", RGB(255, 255, 255), false},
		{"rgba(1.0,1.0,1.0,0.5)", RGB(255, 255, 255), false},
		{"rgba(1.0,1.0,1.0)", RGB(255, 255, 255), false},
		{"0,0,0", RGB(0, 0, 0), false},
		{"174,235,255", RGB(174, 235, 255), false},
		{"0, 0, 0", RGB(0, 0, 0), false},
		{"0 0 0", RGB(0, 0, 0), false},
		{"255 255 255", RGB(255, 255, 255), false},
		{"255 255 255 0.5", RGB(255, 255, 255), false},

		{"#000", Color(0x000), false},
		{"#FFF", Color(0xffffff), false},
		{"#fff", Color(0xffffff), false},
		{"#ABC", Color(0xaabbcc), false},
		{"#AABBCC", Color(0xaabbcc), false},
		{"#aabbcc", Color(0xaabbcc), false},
		{"#abc", Color(0xaabbcc), false},
		{"#def", Color(0xddeeff), false},
		{"def", Color(0xddeeff), false},
		{"#1000000", 0, true},
		{"#aaff", 0, true},
	}

	for _, c := range cases {
		act, err := Parse(c.color)

		if err != nil {
			if !c.err {
				t.Errorf("Expected Parse(%q) to be %v, got error: %v", c.color, c.exp, err)
				continue
			}
		} else {
			if c.err {
				t.Errorf("Expected Parse(%q) to return error", c.color)
				continue
			}

			if act == nil {
				t.Errorf("Expected Parse(%q) to be %v, got nil", c.color, c.exp)
				continue
			}

			if act.Color() != c.exp {
				t.Errorf("Expected Parse(%q) to return %v, got %v", c.color, c.exp, act)
			}
		}
	}
}
