// Copyright 2016 The go-ego Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-ego/ego/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package renders

import (
	"fmt"
)

func Colorize(text string, status string) string {
	out := ""
	switch status {
	case "succ":
		out = "\033[32;1m" // Blue
	case "fail":
		out = "\033[31;1m" // Red
	case "warn":
		out = "\033[33;1m" // Yellow
	case "note":
		out = "\033[34;1m" // Green
	case "blue":
		out = "\033[44;1m" // blue
	default:
		out = "\033[0m" // Default
	}
	return out + text + "\033[0m"
}

type ColorCode string

const (
	cblack    ColorCode = "0;30m"
	dred                = "41;1m"
	cred                = "0;31m"
	cgreen              = "0;32m"
	btextn              = "0;33m"
	navy                = "0;34m"
	purple              = "0;35m"
	ccyan               = "0;36m"
	gray                = "0;37m"
	dim                 = "1;30m"
	orange              = "1;31m"
	lime                = "1;32m"
	cyellow             = "1;33m"
	cblue               = "1;34m"
	pink                = "1;35m"
	aqua                = "1;36m"
	cwhite              = "1;37m"
	underline           = "4m"
	display             = "7m"
)

func colorize(text interface{}, color ColorCode) string {
	return fmt.Sprintf("\033[%s%v\033[0m", color, text)
}

func Black(text interface{}) string {
	return colorize(text, cblack)
}

func RedArr(text interface{}) []string {
	var redarr []string
	for _, v := range text.([]string) {
		redarr = append(redarr, colorize(v, cred))
	}
	return redarr
}

func Red(text interface{}) string {
	return colorize(text, cred)
}

func Dred(text interface{}) string {
	return colorize(text, dred)
}

func Green(text interface{}) string {
	return colorize(text, cgreen)
}

func Btextn(text interface{}) string {
	return colorize(text, btextn)
}

func Navy(text interface{}) string {
	return colorize(text, navy)
}

func Purple(text interface{}) string {
	return colorize(text, purple)
}

func Cyan(text interface{}) string {
	return colorize(text, ccyan)
}

func Gray(text interface{}) string {
	return colorize(text, gray)
}

func Dim(text interface{}) string {
	return colorize(text, dim)
}

func Orange(text interface{}) string {
	return colorize(text, orange)
}

func Lime(text interface{}) string {
	return colorize(text, lime)
}

func Yellow(text interface{}) string {
	return colorize(text, cyellow)
}

func Blue(text interface{}) string {
	return colorize(text, cblue)
}

func Pink(text interface{}) string {
	return colorize(text, pink)
}

func Aqua(text interface{}) string {
	return colorize(text, aqua)
}

func Lblue(text interface{}) string {
	return colorize(text, aqua)
}

func White(text interface{}) string {
	return colorize(text, cwhite)
}

func bold(message string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[21m", message)
}

func Uline(text interface{}) string {
	return colorize(text, underline)
}

func Display(text interface{}) string {
	return colorize(text, display)
}

// Black returns a black string
// func Black(message string) string {
// 	return fmt.Sprintf("\x1b[30m%s\x1b[0m", message)
// }

// // White returns a white string
// func White(message string) string {
// 	return fmt.Sprintf("\x1b[37m%s\x1b[0m", message)
// }

// // Cyan returns a cyan string
// func Cyan(message string) string {
// 	return fmt.Sprintf("\x1b[36m%s\x1b[0m", message)
// }

// // Blue returns a blue string
// func Blue(message string) string {
// 	return fmt.Sprintf("\x1b[34m%s\x1b[0m", message)
// }

// // Red returns a red string
// func Red(message string) string {
// 	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
// }

// // Green returns a green string
// func Green(message string) string {
// 	return fmt.Sprintf("\x1b[32m%s\x1b[0m", message)
// }

// // Yellow returns a yellow string
// func Yellow(message string) string {
// 	return fmt.Sprintf("\x1b[33m%s\x1b[0m", message)
// }

// // Gray returns a gray string
// func Gray(message string) string {
// 	return fmt.Sprintf("\x1b[37m%s\x1b[0m", message)
// }

// Magenta returns a magenta string
func Magenta(message string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", message)
}

// BlackBold returns a black bold string
func BlackBold(message string) string {
	return fmt.Sprintf("\x1b[30m%s\x1b[0m", bold(message))
}

// WhiteBold returns a white bold string
func WhiteBold(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", bold(message))
}

// CyanBold returns a cyan bold string
func CyanBold(message string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", bold(message))
}

// BlueBold returns a blue bold string
func BlueBold(message string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", bold(message))
}

// RedBold returns a red bold string
func RedBold(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", bold(message))
}

// GreenBold returns a green bold string
func GreenBold(message string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", bold(message))
}

// YellowBold returns a yellow bold string
func YellowBold(message string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", bold(message))
}

// GrayBold returns a gray bold string
func GrayBold(message string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", bold(message))
}

// MagentaBold returns a magenta bold string
func MagentaBold(message string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", bold(message))
}
