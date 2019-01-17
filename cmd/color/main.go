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

type colorFormat struct {
	key   int
	name  string
	value string
}

func (f *colorFormat) String() string {
	return fmt.Sprintf("[%d] %s\t%v", f.key, f.name, f.value)
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

	converted := representations(format)

	fmt.Println()
	fmt.Printf(" Input (%s)\t%v", name(format), format.String())
	fmt.Println()
	fmt.Println()
	for _, f := range converted {
		_, err = fmt.Fprintf(writer, " %s\n", f.String())
		if err != nil {
			fatalErr(errWriter, err)
		}
	}
	fmt.Println()

	_ = clipboardPrompt(reader, writer, converted)
}

func representations(f color.Format) []colorFormat {
	c := f.Color()
	rgb := c.RGBInt()
	formats := []colorFormat{
		{1, "sRGB", strconv.Itoa(int(c))},
		{2, "Hex", c.Hex().String()},
		{3, "RGB", fmt.Sprintf("%v %v %v", rgb.R, rgb.G, rgb.B)},
		{4, "RGB", c.RGBInt().String()},
	}

	return formats
}

func clipboardPrompt(in io.Reader, out io.Writer, formats []colorFormat) error {
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

func fatalErr(w io.Writer, e error) {
	_, _ = fmt.Fprintln(w, "error:", e)
	os.Exit(1)
}

func prompt(w io.Writer) {
	_, err := fmt.Fprintf(w, "> ")
	if err != nil {
		fatalErr(w, err)
	}
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
