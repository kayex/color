package color

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var patterns = []struct {
	pattern *regexp.Regexp
	parser  func(string) (Format, error)
}{
	{
		// Matches six and three digit hex colors (#FFFFFF, #FFF)
		pattern: regexp.MustCompile(`#([ABCDEFabcdef0-9]{6}|[ABCDEFabcdef0-9]{3})`),
		parser:  parseHex,
	},
	{
		// Matches RGB colors with int channel values (rgb(255, 255, 255)) including
		// RGBA strings with a float alpha-channel value (rgba(255, 255, 255, 0.5).
		pattern: regexp.MustCompile(`(?:rgba?\()?(?:[0-9]{1,3}(?:, |,| )?){3}(?:[01]\.[0-9]+)?\)?`),
		parser:  parseRGBint,
	},
	{
		// Matches RGB colors with float channel values (rgb(1.0, 1.0, 1.0)) including
		// RGBA strings with a an alpha-channel value (rgba(1.0, 1.0, 1.0, 0.5).
		pattern: regexp.MustCompile(`(?:rgba?\()?(?:(?:[01])\.[0-9]+(?:, |,| )?){3,4}\)?`),
		parser:  parseRGBfloat,
	},
}

func Parse(s string) (Format, error) {
	for _, p := range patterns {
		match := p.pattern.FindString(s)
		if match != "" {
			c, err := p.parser(match)
			if err != nil {
				return nil, err
			}

			return c, nil
		}
	}

	return nil, fmt.Errorf("unknown color format: %s", s)
}

func parseHex(s string) (Format, error) {
	h := strings.TrimPrefix(s, "#")
	if !validHex(h) {
		return nil, fmt.Errorf("invalid hex value %q", h)
	}

	// Here it is safe to take the byte length of the string, since no characters that are valid hexadecimal values
	// are bigger than 1B.
	hl := len(h)
	switch hl {
	case 6:
	case 3:
		h = extendShortHex(h)
	default:
		return nil, fmt.Errorf("invalid hex format %q, hex colors should be either 3 or 6 characters long", h)
	}

	v, err := strconv.ParseInt(h, 16, 32)
	if err != nil {
		return nil, err
	}
	c := Color(v)

	if c < CMin || c > CMax {
		return nil, fmt.Errorf("invalid color value %q, values should be between %x and %x", h, CMin, CMax)
	}

	return HexColor(h), nil
}

func parseRGBint(s string) (Format, error) {
	rgbValues, err := rgbValues(s)
	if err != nil {
		return nil, err
	}

	var c []uint8
	for i := 0; i < 3; i++ {
		v, err := strconv.Atoi(rgbValues[i])
		if err != nil {
			return nil, fmt.Errorf("could not decode RGB int channel value %q in color string: %s", rgbValues[i], s)
		} else if v > 0xff {
			return nil, fmt.Errorf("invalid RGB channel value %q in color string: %s", strconv.Itoa(v), s)
		}

		c = append(c, uint8(v))
	}

	return RGBInt{c[0], c[1], c[2]}, nil
}

func parseRGBfloat(s string) (Format, error) {
	rgbValues, err := rgbValues(s)
	if err != nil {
		return nil, err
	}

	var c []float32
	for i := 0; i < 3; i++ {
		v, err := strconv.ParseFloat(rgbValues[i], 64)
		if err != nil || v > 1.0 {
			return nil, fmt.Errorf("invalid RGB channel value %q in color string %s", FormatRGBChannelFloat(v), s)
		}

		c = append(c, float32(v))
	}

	return RGBFloat{c[0], c[1], c[2]}, nil
}

func rgbValues(s string) ([]string, error) {
	original := copyString(s)
	// Remove "rgb()"
	if (strings.HasPrefix(s, "rgb(") || strings.HasPrefix(s, "rgba(")) && strings.HasSuffix(s, ")") {
		s = strings.TrimLeft(s, "rgba(")
		s = strings.TrimRight(s, ")")
	}

	var channels []string

	if strings.Index(s, ",") != -1 {
		s = strings.Replace(s, " ", "", -1)
		channels = strings.Split(s, ",")
	} else {
		channels = strings.Split(s, " ")
	}

	if len(channels) < 3 || len(channels) > 4 {
		return nil, fmt.Errorf("invalid RGB color string: %s", original)
	}

	return channels, nil
}

func validHex(hex string) bool {
	invalidChar := strings.IndexFunc(hex, func(r rune) bool {
		return !('0' <= r && r <= 'f')
	})
	valid := invalidChar == -1

	return valid
}

func copyString(a string) string {
	return (a + " ")[:len(a)]
}
