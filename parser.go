package color

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var patterns = []struct {
	pattern *regexp.Regexp
	parser  func(string) (Color, error)
}{
	{
		// Matches six digit hex colors (#FFFFFF)
		pattern: regexp.MustCompile(`^#?([ABCDEFabcdef0-9]{2}){3}$`),
		parser:  parseHex,
	},
	{
		// Matches three digit hex colors (#FFF)
		pattern: regexp.MustCompile(`^#?([ABCDEFabcdef0-9]){3}$`),
		parser:  parseHex,
	},
	{
		// Matches RGB colors with int channel values (rgb(255, 255, 255))
		pattern: regexp.MustCompile(`^(rgb\()?(?:[0-9]{1,3}(?:,|, )?\)?){3}$`),
		parser:  parseRGBint,
	},
	{
		// Matches RGB colors with float channel values (rgb(1.0, 1.0, 1.0))
		pattern: regexp.MustCompile(`^(?:rgb\()?(?:(?:[01])\.[0-9](?:,|, )?){3}\)?$`),
		parser:  parseRGBfloat,
	},
}

func Parse(s string) (Color, error) {
	for _, p := range patterns {
		match := p.pattern.MatchString(s)
		if match {
			c, err := p.parser(s)
			if err != nil {
				return 0, err
			}

			return c, nil
		}
	}

	return 0, fmt.Errorf("unknown color format: %s", s)
}

func parseHex(s string) (Color, error) {
	h := strings.TrimPrefix(s, "#")
	if !validHex(h) {
		return 0, fmt.Errorf("invalid hex value %q", h)
	}

	// Here it is safe to take the byte length of the string, since no characters that are valid hexadecimal values
	// are bigger than 1B.
	hl := len(h)
	switch hl {
	case 6:
	case 3:
		h = extendShortHex(h)
	default:
		return 0, fmt.Errorf("invalid hex format %q, hex colors should be either 3 or 6 characters long", h)
	}

	v, err := strconv.ParseInt(h, 16, 32)
	if err != nil {
		return 0, err
	}
	c := Color(v)

	if c < CMin || c > CMax {
		return 0, fmt.Errorf("invalid color value %q, values should be between %x and %x", h, CMin, CMax)
	}

	return c, nil
}

func parseRGBint(s string) (Color, error) {
	rgbValues, err := rgbValues(s)
	if err != nil {
		return 0, err
	}

	var channels []uint8
	for _, v := range rgbValues {
		c, err := strconv.Atoi(v)
		if err != nil || c > 255 {
			return 0, fmt.Errorf("invalid RGB color channel value: %s", s)
		}

		channels = append(channels, uint8(c))
	}

	return RGB(channels[0], channels[1], channels[2]), nil
}

func parseRGBfloat(s string) (Color, error) {
	rgbValues, err := rgbValues(s)
	if err != nil {
		return 0, err
	}

	var channels []uint8
	for _, v := range rgbValues {
		c, err := strconv.ParseFloat(v, 64)
		if err != nil || c > 1.0 {
			return 0, fmt.Errorf("invalid RGB color channel value: %s", s)
		}

		var ci uint8
		if c == 0.0 {
			ci = 0
		} else {
			ci = uint8(math.Round(255 / c))
		}

		channels = append(channels, ci)
	}

	return RGB(channels[0], channels[1], channels[2]), nil
}

func rgbValues(s string) ([]string, error) {
	// Remove "rgb()"
	if strings.HasPrefix(s, "rgb(") && strings.HasSuffix(s, ")") {
		s = strings.TrimLeft(s, "rgb(")
		s = strings.TrimRight(s, ")")
	}

	// Remove spaces
	s = strings.Replace(s, " ", "", -1)

	channels := strings.Split(s, ",")

	if len(channels) != 3 {
		return nil, fmt.Errorf("invalid RGB color string: %s", s)
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
