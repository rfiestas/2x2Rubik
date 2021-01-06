package main

import "fmt"

func createColorsHelper() *colorsHelper {
	return &colorsHelper{
		0: {
			xterm:   15, // White	(SYSTEM)		#ffffff		rgb(255,255,255)
			inverse: 5,  // White -> Yellow
		},
		1: {
			xterm:   40, // Green	Green3			#00d700		rgb(0,215,0)
			inverse: 3,  // Green -> Blue
		},
		2: {
			xterm:   196, // Red	Red1			#ff0000		rgb(255,0,0)
			inverse: 4,   // Red -> Orange
		},
		3: {
			xterm:   33, // Blue	DodgerBlue1		#0087ff		rgb(0,135,255)
			inverse: 1,  // Blue -> Green
		},
		4: {
			xterm:   208, // Orange	DarkOrange		#ff8700		rgb(255,135,0)
			inverse: 2,   // Orange -> Red
		},
		5: {
			xterm:   11, // Yellow 	(SYSTEM)		#ffff00		rgb(255,255,0)
			inverse: 0,  // Yellow -> White
		},
	}
}

type colorProperties struct {
	xterm   byte //https://jonasjacek.github.io/colors/
	inverse byte //Inverse color map
}
type colorsHelper []colorProperties

func (c colorsHelper) getInverseColor(color byte) byte {
	return c[color].inverse
}

func (c colorsHelper) xtermPrint(color byte, text string) string {
	return fmt.Sprintf("\033[1;38;5;%dm%s\033[0m", c[color].xterm, text)
}

func (c *colorsHelper) printable(colors []byte, text string) (printable []interface{}) {
	for _, color := range colors {
		printable = append(printable, c.xtermPrint(color, text))
	}
	return printable
}
