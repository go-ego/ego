// Copyright 2016 The go-ego Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-ego/ego/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package rego

import (
	// "fmt"
	"regexp"
	"strings"
)

func TrimNotes(rfile string) string {
	regc := regexp.MustCompile("(?U)\\/\\*[\\s\\S]*\\*\\/|\\/\\/.*(\n|\r)")

	rname := rfile + "\n"
	restr := regc.ReplaceAllString(rname, "")
	return restr
}

func TrimSlot(text string) string {
	slot := strings.Replace(text, "<slot></slot>", "", -1)
	// fmt.Println(slot)
	return slot
}

func TrimQ(tkey string, restr []string) []string {
	var reval []string
	for i := 0; i < len(restr); i++ {
		tval := strings.Trim(restr[i], tkey+"\\=\\")
		bval := strings.Trim(tval, `\"`)

		reval = append(reval, bval)
	}
	return reval
}

func TrimTVal(tkey string, restr []string) []string {
	var reval []string
	for i := 0; i < len(restr); i++ {
		tval := strings.Trim(restr[i], tkey+"\\=\\")
		var bval string

		if strings.HasSuffix(tval, "}") {
			bval = strings.Trim(tval, `\{|\}`)
		} else {
			bval = strings.Trim(tval, `\"`)
		}

		reval = append(reval, bval)
	}
	return reval
}

func TrimVal(fval []string, keyarr, header, tvgostr string) string {
	tkey := TrimB(keyarr)
	tfval := TrimTVal(tkey, fval)

	if len(header) == 0 {
		header = strings.Replace(tvgostr, keyarr, tfval[0], -1)
	} else {
		header = strings.Replace(header, keyarr, tfval[0], -1)
	}

	return header
}

func Quotation(rfile string) []string {
	regc := regexp.MustCompile(`(?U)\".*\"`)
	restr := regc.FindAllString(rfile, -1)
	return restr
}

func TrimQt(str string) string {
	restr := strings.Trim(str, `\"`)
	return restr
}

func TrimB(str string) string {
	restr := strings.Trim(str, `\{|\}`)
	return restr
}

//TrimBlank
func TrimBlank(str string) string {
	regt := regexp.MustCompile(`\t|\r|\n|\"|\s*`)
	restr := regt.ReplaceAllString(str, "")
	return restr
}

func TrimIs(str string) string {
	regt := regexp.MustCompile("(?U)\\<style\\>[\\s\\S]*\\<\\/style\\>|\\<script\\>[\\s\\S]*\\<\\/script\\>")
	restr := regt.ReplaceAllString(str, "")

	regi := regexp.MustCompile(`(?U)import \([\s\S]*\)|import \"[\s\S]*\"`)
	reistr := regi.ReplaceAllString(restr, "")

	return reistr
}

func TrimS(str string) string {
	regt := regexp.MustCompile("(?U)\\<style\\>[\\s\\S]*\\<\\/style\\>|\\<script\\>[\\s\\S]*\\<\\/script\\>")
	restr := regt.ReplaceAllString(str, "")
	return restr
}

func TrimBrace(rfile string) string {
	regc := regexp.MustCompile("(?U)\\{[^\\{|^\\}|^\\.|^\\<|^\\>]*\\}")
	restr := regc.ReplaceAllString(rfile, "")
	return restr
}

func TBrace(rfile string) string {
	restr := strings.Trim(rfile, `\{\}`)
	return restr
}
