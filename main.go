package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "4fbe6e",
		"white": "ffffff",
	}

	printMap(colors)
	/*---------------------------------------
		colors := map[string]string{
			"red":   "#ff0000",
			"green": "4fbe6ed",
		}
	-----------------------------------------*/

	// another way:
	//var colors map[string]string
	/*-----------------------------------------
		// another way:
		colors := make(map[int]string)
		colors[10] = "fffffff"
		colors[8] = "efefeee"

		delete(colors, 8) // deleting from map

		fmt.Println(colors)
	------------------------------------------*/
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("HEX code of color ", color, " is ", hex)
	}
}
