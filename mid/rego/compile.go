// Copyright 2016 The go-ego Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-ego/ego/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

// Package rego renders vgo
package rego

import (
	"strings"
)

func RendersVgo(name string) {
	var (
		rname string
		racss string
		rajs  string
	)

	layout, err := Readfile("public/layout.html")
	if err != nil {
		panic(err)
	}

	if len(name) > 0 {
		rname = "public/" + name + ".vgo"
	} else {
		rname = "public/banner.vgo"
	}

	abanner, err := Readfile(rname)
	if err != nil {
		panic(err)
	}

	notes := TrimNotes(abanner)
	_, retext, _, class, recss, rejs := ImpStr(notes)

	aRecss, aRejs := styScr(notes)

	for i := 0; i < len(aRecss); i++ {
		racss += aRecss[i]
	}

	for i := 0; i < len(aRejs); i++ {
		rajs += aRejs[i]
	}

	for h := 0; h < len(recss); h++ {
		for i := 0; i < len(recss[h]); i++ {
			if strings.Contains(racss, recss[h][i]) {
				recss[h][i] = ""
			}
			racss += recss[h][i]
		}
	}

	for h := 0; h < len(rejs); h++ {
		for i := 0; i < len(rejs[h]); i++ {
			if strings.Contains(rajs, rejs[h][i]) {
				rejs[h][i] = ""
			}
			rajs += rejs[h][i]
		}
	}

	trimtext := TrimIs(retext)

	var wname string
	if strings.Contains(name, "/") {
		sname := strings.Split(name, "/")
		wname = sname[0] + "_" + sname[1]
	} else {
		wname = name
	}

	WHtml(class, layout, racss, trimtext, rajs, wname)

}
