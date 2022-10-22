package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf746",
		"white": "#ffffff",
	}
	printMap(colors)

	//var rect = map[string]string{"height": "10"}
	//fmt.Println(rect)

	//colors := make(map[int]string)
	//colors[10] = "#fff"
	//colors[11] = "#rrr"
	//fmt.Println(colors)
	//
	//delete(colors, 10)
	//fmt.Println(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
