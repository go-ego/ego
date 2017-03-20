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

// "fmt"

func css(str string) string {
	amap := Map{
		//text
		"lh:": "line-height:",
		"ta:": "text-align:",
		"va:": "vertical-align:",
		"td:": "text-decoration:",
		"tt:": "text-transform:",
		"ti:": "text-indent:",
		"ts:": "text-shadow:",
		"ws:": "white-space:",
		"wb:": "word-break:",
		"to:": "text-overflow:",
		//font
		// "f:": "font",
		"fs:":  "font-size:",
		"ff:":  "font-family:",
		"fst:": "font-style:",
		"fw:":  "font-weight:",
		"fv:":  "font-variant:",
		//bg
		"bg:": "background:",
		// "bgc:": "background-color:",
		"bc:":  "background-color:",
		"bi:":  "background-image:",
		"bp:":  "background-position:",
		"bgr:": "background-repeat:",
		"ba:":  "background-attachment:",
		"bgs":  "background-size:",
		"bo:":  "background-origin:",
		"bgc:": "background-clip:",
		//margin
		"mg:": "margin:",
		"ml:": "margin-left:",
		"mr:": "margin-right:",
		"mt:": "margin-top:",
		"mb:": "margin-bottom:",
		//
		"mh:": "max-height:",
		"mw":  "max-width:",
		"mih": "min-height:",
		"miw": "min-width:",
		"vb:": "visibility:",
		"dp:": "display:",
		"pt:": "position:",
		//padding
		"pd:":  "padding:",
		"pdt:": "padding-top:",
		"pb:":  "padding-bottom:",
		"pr:":  "padding-right:",
		"pl:":  "padding-left:",
		//list
		"ls:":  "list-style:",
		"lst:": "list-style-type:",
		"lsi:": "list-style-image:",
		"lsp:": "list-style-position:",
		//border
		"bd:":  "border:",
		"bw:":  "border-width:",
		"bt:":  "border-top:",
		"bb:":  "border-bottom:",
		"bl:":  "border-left:",
		"bri:": "border-right:",
		"bdi:": "border-image:",
		// border-style
		"bs:":  "border-style:",
		"bts:": "border-top-style:",
		"brs:": "border-right-style:",
		"bbs:": "border-bottom-style:",
		"bls:": "border-left-style:",
		"bdc:": "border-color:",
		"br:":  "border-radius:",
		"bca:": "border-collapse:",
		// "bdca:": "border-collapse:",
		"boxs:": "box-shadow",
		//outline
		"oc:": "outline-color:",
		"os:": "outline-style:",
		"ow:": "outline-width:",
		//
		"tf:": "transform",
	}

	recss := re(str, amap)
	return recss
}
