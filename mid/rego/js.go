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
	"strings"
)

func reJs(str string) string {
	amap := Map{
		//fn
		"func ": "function ",
		"fn ":   "function ",
		"log(":  "console.log(",
		"doc.":  "document.",
		"win.":  "window.",
		"nav.":  "navigator.",
		//doc
		"doc(#":     "document.getElementById(",
		"docid(":    "document.getElementById(",
		"doctag(":   "document.getElementsByTagName(",
		"doclass(":  "document.getElementsByClassName(",
		"docname(":  "document.getElementsByName(",
		"docquery(": "document.querySelector(",
		"docall(":   "document.querySelectorAll(",
		//???
		"docId(":    "document.getElementById(",
		"docTag(":   "document.getElementsByTagName(",
		"docClass(": "document.getElementsByClassName(",
		"docName(":  "document.getElementsByName(",
		"docQuery(": "document.querySelector(",
		"docAll(":   "document.querySelectorAll(",
		"tag(":      "getElementsByTagName(",
		"class(":    "getElementsByClassName(",
		"name(":     "getElementsByName(",
		"query(":    "querySelector(",
		"all(":      "querySelectorAll(",
		//val
		".html": ".innerHTML",
		".val":  ".value",
		".text": ".innerText",
		//style
		".color": ".style.color",
		"fonts":  ".style.fontSize",
		"fontf":  "style.fontFamily",
		//
		".addEvent(": ".addEventListener(",
	}

	rejs := re(str, amap)
	return rejs
}

func reIf(str string) string {
	iftext := strings.Replace(str, "if ", "if (", -1)
	// fmt.Println("str--------", iftext)
	strtext := strings.Replace(iftext, " {", "){", -1)
	return strtext
}

func reWhile(str string) string {
	fortext := strings.Replace(str, "while ", "while (", -1)
	strtext := strings.Replace(fortext, " {", "){", -1)
	return strtext
}

func reFor(str string) string {
	fortext := strings.Replace(str, "for ", "for (", -1)
	strtext := strings.Replace(fortext, " {", "){", -1)
	return strtext
}

func reVar(str string) string {
	vartext := strings.Replace(str, ":=", "var ", -1)
	strtext := strings.Replace(vartext, " {", "){", -1)
	return strtext
}

func reLen(str string) string {
	lentext := strings.Replace(str, "len(", " ", -1)
	strtext := strings.Replace(lentext, ")", ".length", -1)
	return strtext
}
