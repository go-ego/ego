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
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func FindImp(text string) ([]string, []string) {
	regc := regexp.MustCompile("(?U)import \\([\\s\\S]*\\)")
	reimp := regc.FindAllString(text, -1)
	regq := regexp.MustCompile(`(?U)import \"[\s\S]*\"`)
	reimpq := regq.FindAllString(text, -1)
	return reimp, reimpq
}

func IsImp(trimName []string, keyarr, tnotes string) string {
	var abool string
	for h := 0; h < len(trimName); h++ {
		if len(ImpBrace(trimName[h], keyarr, tnotes)) > 0 {
			abool += "true"
		} else {
			abool += "false"
		}
	}

	return abool
}

type Tag struct {
	Name   string
	PosArr []int
}

var tagArr []Tag

func findAllTag(text, key string) []string {
	// var intArr []int
	var nameArr []string

	tag := Tag{}

	regc := regexp.MustCompile("(?U)\\<.*\\>")
	retag := regc.FindAllString(text, -1)

	for i := 0; i < len(retag); i++ {
		if retag[i] == "<"+key+">" || retag[i] == "</"+key+">" {
			tag.Name = key
			tag.PosArr = append(tag.PosArr, i)
		}
	}

	if tag.Name != "" {
		tagArr = append(tagArr, tag)
	}

	if len(tagArr) > 0 {
		if len(tagArr) > 1 {
			for i := 0; i < len(tagArr); i++ {
				nameArr = append(nameArr, tagArr[i].Name)
			}

			var (
				fben int
				fend int
				tben int
				tend int
			)
			// var status int
			for i := 0; i < len(tagArr); i++ {
				var status int

				for iplus := 1; iplus < len(tagArr); iplus++ {

					for h := 0; h < len(tagArr[iplus-1].PosArr); h++ {
						var hplus int
						hlen := len(tagArr[i].PosArr)
						if h+1 < hlen {
							hplus = h + 1
						}

						if !isOdd(hplus) {
							fben = tagArr[iplus-1].PosArr[hplus-1] //t
							fend = tagArr[iplus-1].PosArr[hplus]   //tplus
						}

						for t := 0; t < len(tagArr[iplus].PosArr); t++ {
							var tplus int
							tlen := len(tagArr[i].PosArr)

							if t+1 < tlen {
								tplus = t + 1
							}

							if !isOdd(tplus) {
								tben = tagArr[iplus].PosArr[tplus-1]
								tend = tagArr[iplus].PosArr[tplus]
							}

							if fben < tben && fend > tend {
								status++
								if status == 1 {
									tagArr[iplus], tagArr[iplus-1] = tagArr[iplus-1], tagArr[iplus]
									nameArr[iplus], nameArr[iplus-1] = nameArr[iplus-1], nameArr[iplus]
								} else {
									break
								}
							}
						}

					}
				}

			}

		} else {
			nameArr = append(nameArr, tagArr[0].Name)
		}
	}

	return nameArr
	// return retag
}

func findImpTag(text, key string) []string {
	regc := regexp.MustCompile("(?U)[\\<" + key + "\\>|\\<\\/" + key + "\\>]")
	retag := regc.FindAllString(text, -1)

	return retag
}

func ImpMap(notes string) (Map, []string, []string) {
	var (
		atag    []Tag
		tagName []string
		class   []string
	)

	amap := Map{}
	tagArr = atag

	imp, impq := FindImp(notes)

	if len(imp) == 0 && len(impq) == 0 {
		return amap, class, class
	}

	if len(imp) > 0 {
		class = Quotation(imp[0])

		name, nclass := ImpName(imp[0])
		for i := 0; i < len(name); i++ {
			fname := TrimBlank(name[i])

			if fname != "" {
				amap[fname] = nclass[i]
			}

			tagName = findAllTag(notes, fname)
		}
	}

	if len(impq) > 0 {

		nameq, nclassq := ImpName(impq[0])

		for i := 0; i < len(nameq); i++ {
			fname := TrimBlank(nameq[i])

			if fname != "" {
				amap[fname] = nclassq[i]
			}

			tagName = findAllTag(notes, fname)
		}
	}

	return amap, tagName, class
}

var (
	astatus int
	chtml   []string
	acss    string
	ajs     string
)

func isClass(file, random string) []string {
	var (
		src  string
		dst  string
		root string
	)

	if strings.HasSuffix(file, ".vgo") {
		fmt.Println("--------------------------", file)
	} else if strings.HasSuffix(file, ".ttf") {
		if astatus == 1 {
			src = file
			repub := strings.Trim(file, "./public")
			dst = "views/" + repub
		} else {
			acss = "views/css/"
			src = "./public/css/" + file
			dst = acss + file
		}

		CopyFile(src, dst)
	} else if strings.HasSuffix(file, ".css") {

		if astatus == 1 {
			src = file
			repub := strings.Trim(file, "./public")
			dst = "views/" + repub
			root = "{{.root}}/" + repub
		} else {
			acss = "views/css/"
			src = "./public/css/" + file
			dst = acss + file
			root = "{{.root}}/" + file
		}

		chtml = append(chtml, `<link rel="stylesheet" type="text/css" href="`+root+random+`"/>`)

		CopyFile(src, dst)
	} else if strings.HasSuffix(file, ".js") {

		if astatus == 1 {

			src = file
			repub := strings.Trim(file, "./public")
			dst = "views/" + repub
			root = "{{.root}}/" + repub
		} else {
			ajs = "views/js/"
			src = "./public/js/" + file
			dst = ajs + file
			root = "{{.root}}/" + file
		}

		chtml = append(chtml, `<script type="text/javascript" src="`+root+random+`"></script>`)
		CopyFile(src, dst)
	} else {
		astatus = 1

		dsrc := "views/src/"
		src := "./public/src/" + file
		afile, err := ListFile(src, "")
		checkErr(err)

		ImpClass(afile)
		dst := dsrc + file
		CopyFile(src, dst)
	}

	return chtml
}

var (
	status int
	// ahtml  []string
)

func UnuseRand() {
	status = 1
}

//ImpClass
func ImpClass(class []string, userandom ...int) []string {
	var ahtml []string
	// chtml = ahtml

	for _, v := range class {
		var random string

		if len(userandom) == 0 && status != 1 {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			random = "?" + strconv.Itoa(r.Intn(10000000))
		}

		file := strings.Trim(v, `\"`)
		ahtml = isClass(file, random)
	}

	return ahtml
}

func ImpName(str string) ([]string, []string) {

	regc := regexp.MustCompile(`(?U)[\n|\t].*\"`)
	restr := regc.FindAllString(str, -1)
	restrq := Quotation(str)
	return restr, restrq
}

func TrimName(name []string) []string {
	var restr []string
	for i := 0; i < len(name); i++ {
		fname := TrimBlank(name[i])
		if fname != "" {
			restr = append(restr, fname)
		}
	}

	return restr
}

func ImpStr(notes string) ([]string, string, []string, []string, [][]string, [][]string) {
	amap, name, class := ImpMap(notes)
	text, retext, _, val, recss, rejs := FindTag(amap, name, notes)

	return text, retext, val, class, recss, rejs
}

var (
	vhtml  string
	thtml  string
	header string
	bodyer string
	footer string
)

func WHtml(class []string, layout string, args ...string) {
	var (
		rhtml []string
		wjs   string
		wcss  string
		aname string
		abody string
	)

	chtml = rhtml

	impclass := ImpClass(class)

	for i := 0; i < len(impclass); i++ {
		if strings.Contains(impclass[i], ".js") {
			wjs += impclass[i]
		} else if strings.Contains(impclass[i], ".css") {
			wcss += impclass[i]
		}
	}

	if len(args) > 0 {
		wcss += args[0]
		abody = args[1]
		wjs += args[2]

		aname += "views/html/" + args[3] + ".html"
	}

	header = strings.Replace(layout, "</head>", wcss+"</head>", -1)
	bodyer = strings.Replace(header, "<div id=app>", "<div id=app>"+abody, -1)
	thtml = strings.Replace(bodyer, "</body>", "</body>"+wjs, -1)

	minhtml, err := minStr(thtml, "text/html")
	if err == nil {
		Writefile(minhtml, aname)
	} else {
		Writefile(thtml, aname)
	}
}

func ImpHtml(class []string, layout string) {

	impclass := ImpClass(class)

	head := strings.Split(layout, "</head>")
	body := strings.Split(head[1], "</body>")
	foot := strings.Split(layout, "</body>")

	header = head[0] + impclass[1] + "</head>"
	bodyer = body[0] + "</body>"
	footer = impclass[0] + foot[1]
	vhtml = header + bodyer + footer

	Writefile(vhtml, "views/t_vgo.html")
}
