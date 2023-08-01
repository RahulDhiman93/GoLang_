package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#ds0000",
		"white": "#ffffff",
	}

	delete(colors, "blue")
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println("Color is", key, "with hex", val)
	}
}
