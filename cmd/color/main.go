package main

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/kayex/color"
	"os"
	"strconv"
)

func printFormats(formats []colorFormat) {
	for _, format := range formats {
		fmt.Printf(" [%d] %s\t%v\n", format.key, format.name, format.value)
	}
}

type colorFormat struct {
	key   int
	name  string
	value string
}

func createFormats(f color.Format) []colorFormat {
	c := f.Color()
	ri := c.RGBInt()

	formats := []colorFormat{
		{1, "sRGB", strconv.Itoa(int(c))},
		{2, "Hex", c.Hex().String()},
		{3, "RGB", fmt.Sprintf("%v %v %v", ri.R, ri.G, ri.B)},
		{4, "RGB", c.RGBInt().String()},
	}

	return formats
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	args := os.Args[1:]
	var format color.Format
	var err error

	if len(args) >= 1 {
		input := args[0]
		format, err = color.Parse(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("> ")
		for scanner.Scan() {
			input := scanner.Text()

			if input == "" {
				continue
			}

			format, err = color.Parse(input)
			if err != nil {
				fmt.Println(err)
				fmt.Printf("> ")
			} else {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}

	formats := createFormats(format)

	fmt.Println()
	fmt.Printf(" Input (%s)\t%v", getFormatName(format), format.String())
	fmt.Println()
	fmt.Println()
	printFormats(formats)
	fmt.Println()
	fmt.Printf("Copy> ")

	for scanner.Scan() {
		choice := scanner.Text()

		i, err := strconv.Atoi(choice)
		if err != nil || i > len(formats) {
			continue
		}

		// Copy format value
		v := formats[i-1].value
		clipboard.WriteAll(v)
		fmt.Printf("%v value copied to clipboard.\n", formats[i-1].name)
		break
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func getFormatName(f color.Format) string {
	switch f.(type) {
	case color.HexColor:
		return "hex"
	case color.RGBInt, color.RGBFloat:
		return "RGB"
	}

	return ""
}
