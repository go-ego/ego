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
	"log"
	"regexp"
	"strings"
)

type Map map[string]interface{}

func re(str string, amap Map) string {
	var text string
	var status int
	for k, v := range amap {
		status++
		if status == 1 {
			text = strings.Replace(str, k, v.(string), -1)
		} else {
			text = strings.Replace(text, k, v.(string), -1)
		}
	}

	return text
}

func styScr(text string) ([]string, []string) {
	var recss []string
	var rejs []string

	sty := FindSty(text)
	scr := FindScr(text)

	for i := 0; i < len(sty); i++ {
		recss = append(recss, css(sty[i]))
	}

	for i := 0; i < len(scr); i++ {
		rejs = append(rejs, js(scr[i]))
	}

	return recss, rejs
}

func isOdd(i int) bool {

	return i%2 == 0
	// return (i & 1) != 0
}

func findstr(text, reg string) []string {
	regt := regexp.MustCompile(reg)
	// str := regt.FindAllString(text, 1)
	str := regt.FindAllString(text, -1)

	return str
}

func strReo(str, k, v string) string {
	text := strings.Replace(str, k, v, 1)
	return text
}

func strRe(str, k, v string) string {
	text := strings.Replace(str, k, v, -1)
	return text
}

func ReRegOne(str, reg string, val []string) string {
	var text string

	find := findstr(str, reg)

	for i := 0; i < len(find); i++ {
		if i == 0 {
			text = strings.Replace(str, find[i], val[i], 1)
		} else {
			text = strings.Replace(text, find[i], val[i], 1)
		}
	}

	return text
}

func ReRegNil(str, reg string) string {
	var text string

	find := findstr(str, reg)

	fmt.Println("find---------", find)

	for i := 0; i < len(find); i++ {
		if i == 0 {
			text = strings.Replace(str, find[i], "", -1)
		} else {
			text = strings.Replace(text, find[i], "", -1)
		}
	}

	return text
}

func ReReg(str, reg string, val []string) string {
	var text string

	find := findstr(str, reg)

	for i := 0; i < len(find); i++ {
		if i == 0 {
			text = strings.Replace(str, find[i], val[i], -1)
		} else {
			text = strings.Replace(text, find[i], val[i], -1)
		}
	}

	return text
}

func ReArro(str string, find, val []string) string {
	var text string

	for i := 0; i < len(find); i++ {
		if i == 0 {
			text = strings.Replace(str, find[i], val[i], 1)
		} else {
			text = strings.Replace(text, find[i], val[i], 1)
		}
	}

	return text
}

func ReArrOne(str string, oldFind, newVal []string) string {
	var text string

	for i := 0; i < len(oldFind); i++ {
		if i == 0 {
			text = strings.Replace(str, oldFind[i], newVal[i], 1)
		} else {
			text = strings.Replace(text, oldFind[i], newVal[i], 1)
		}
	}

	return text
}

func ReArr(str string, find, val []string) string {
	var text string

	for i := 0; i < len(find); i++ {
		if i == 0 {
			text = strings.Replace(str, find[i], val[i], -1)
		} else {
			text = strings.Replace(text, find[i], val[i], -1)
		}
	}

	return text
}

func reAllL(str, reg, rep string) string {
	regexp := regexp.MustCompile(reg)
	restr := regexp.ReplaceAllLiteralString(str, rep)

	return restr
}

func reAll(str, reg, rep string) string {
	regt := regexp.MustCompile(reg)
	restr := regt.ReplaceAllString(str, rep)

	return restr
}

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
		// return false
	}
	// return true
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func useUn(obj interface{}) interface{} {
	return obj
}
