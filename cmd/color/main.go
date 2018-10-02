package main

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/kayex/color"
	"log"
	"os"
	"strconv"
)

func printFormats(formats []colorFormat) {
	for _, format := range formats {
		fmt.Printf(" [%d] %s\t%v\n", format.key, format.name, format.value)
	}
	fmt.Println()
}

type colorFormat struct {
	key   int
	name  string
	value string
}

func createFormats(c color.Color) []colorFormat {
	formats := []colorFormat{
		{1, "sRGB", strconv.Itoa(int(c))},
		{2, "Hex", c.Hex().String()},
		{3, "RGB", c.RGBInt().String()},
		{4, "RGB", c.RGBFloat().String()},
	}

	return formats
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: color [value]")
		os.Exit(1)
	}
	input := args[0]

	c, err := color.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	formats := createFormats(c)

	fmt.Println()
	printFormats(formats)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		choice := scanner.Text()

		i, err := strconv.Atoi(choice)
		if err != nil || i > len(formats) {
			continue
		}

		// Copy format value
		v := formats[i-1].value
		clipboard.WriteAll(v)
		fmt.Printf("%v copied to clipboard.\n", formats[i-1].name)
		break
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
