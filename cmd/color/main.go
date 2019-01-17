package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/kayex/color"
	"io"
	"os"
	"strconv"
)

// A conversionOption is used to show the user a format available for conversion to.
type conversionOption struct {
	c color.Format
	// name is the conversionOption name (color format name) that should be displayed to the user.
	name string
	// value is the converted value in a human-readable format.
	value string
}

func (f *conversionOption) String() string {
	return fmt.Sprintf("%s\t%v", f.name, f.value)
}

func main() {
	reader := os.Stdin
	writer := os.Stdout
	errWriter := os.Stderr

	args := os.Args[1:]
	var format color.Format
	var err error

	if len(args) >= 1 {
		format, err = color.Parse(args[0])
		if err != nil {
			fatalErr(writer, err)
		}
	} else {
		format, err = interactiveMode(reader, writer)
		if err != nil {
			fatalErr(errWriter, err)
		}
	}

	fmt.Println()
	fmt.Printf(" Input (%s)\t%v", name(format), format.String())
	fmt.Println()
	fmt.Println()
	ops := options(format)
	for i, f := range ops {
		_, err = fmt.Fprintf(writer, " [%d] %s\n", i+1, f.String())
		if err != nil {
			fatalErr(errWriter, err)
		}
	}
	fmt.Println()

	_ = clipboardPrompt(reader, writer, ops)
}

func options(f color.Format) []conversionOption {
	c := f.Color()
	hex := c.Hex()
	rgb := c.RGBInt()

	options := []conversionOption{
		{c, "sRGB", strconv.Itoa(int(c))},
		{hex, "Hex", hex.String()},
		{rgb, "RGB", fmt.Sprintf("%v %v %v", rgb.R, rgb.G, rgb.B)},
		{rgb, "RGB", rgb.String()},
	}

	return options
}

func interactiveMode(in io.Reader, out io.Writer) (color.Format, error) {
	var format color.Format
	scanner := bufio.NewScanner(in)

Prompt:
	prompt(out)
	if scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			goto Prompt
		}

		var err error
		format, err = color.Parse(input)
		if err != nil {
			fmt.Println(err)
			goto Prompt
		}
	} else {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		return nil, errors.New("EOF")
	}

	return format, nil
}

func clipboardPrompt(in io.Reader, out io.Writer, formats []conversionOption) error {
	prompt(out)
	scanner := bufio.NewScanner(in)
	if scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())

		if err == nil && i <= len(formats) {
			format := formats[i-1]
			v := format.value

			err = clipboard.WriteAll(v)
			if err != nil {
				fmt.Printf("Error copying to clipboard: %v\n", err)
			} else {
				fmt.Printf("%s value copied to clipboard.\n", format.name)
			}
		}
	} else if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func prompt(w io.Writer) {
	_, err := fmt.Fprintf(w, "> ")
	if err != nil {
		fatalErr(w, err)
	}
}

func fatalErr(w io.Writer, e error) {
	_, _ = fmt.Fprintln(w, "error:", e)
	os.Exit(1)
}

func name(f color.Format) string {
	switch f.(type) {
	case color.HexColor:
		return "hex"
	case color.RGBInt, color.RGBFloat:
		return "RGB"
	case color.RGBAInt, color.RGBAFloat:
		return "RGBA"
	}

	return ""
}
