// Copyright 2017 The go-ego Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-ego/ego/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package main

import (
	// "fmt"
	"github.com/go-ego/ego"
)

// const httpUrl string = "http://127.0.0.1:3000"

func main() {

	router := ego.Classic()
	ego.UseRenders()
	// router := ego.Default()

	router.Static("/js", "./views/js")
	router.Static("/src", "./views/src")
	router.GlobHTML("views/html/*")

	// strUrl := httpUrl + "/test/hlist"
	// paramMap := ego.Map{
	// 	"lon":  "10.1010101",
	// 	"lat":  "20.202020",
	// 	"type": "1",
	// }
	// router.TestHtml(strUrl, paramMap)

	router.Ego("/banner/", "banner.html", ego.Map{
		"head": "Test to load the HTML template",
	})

	parArr := [5]int{1, 2, 3, 4, 5}
	router.Ego("/head/", "head/head.html", ego.Map{
		"head":   "Test to load the HTML template",
		"parArr": parArr,
	})

	// rMap := ego.Map{
	// 	"/b1/": "banner.html",
	// 	"/b2/": "banner.html",
	// 	"/he/":  "head/head.html",
	// }
	// router.EgoGroup(rMap)

	router.Run(":3100")
}
