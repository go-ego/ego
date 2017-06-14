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
	"regexp"
	"strings"
)

func findKTag(text, key string) []string {
	regc := regexp.MustCompile("(?U)\\<" + key + "\\>[\\s\\S]*\\<\\/" + key + "\\>")
	restr := regc.FindAllString(text, -1)
	return restr
}

func findLabel(text, key string) []string {
	regstr := "(?U)\\<" + key + "\\>[\\s\\S]*\\<\\/" + key + "\\>"
	regc := regexp.MustCompile(regstr)
	str := regc.FindAllString(text, -1)

	return str
}

func findSlot(text string) []string {
	slot := findLabel(text, "slot")

	return slot
}

func reoSlot(text string) string {
	reslot := strings.Trim(text, "<slot>|</slot>")

	return reslot
}

func reSlot(text string, slotOld, slotNew []string) string {
	var rslotNew []string

	for i := 0; i < len(slotNew); i++ {
		rslot := reoSlot(slotNew[i])

		rslotNew = append(rslotNew, rslot)

	}
	// fmt.Println("slotold:", slotOld, "soltnew:", rslotNew)

	slot := ReArrOne(text, slotOld, rslotNew)

	return slot
}

func findKeyp(text, key string) []string {
	regt := regexp.MustCompile("(?U)\\<" + key + "[^\\>][\\s\\S]*\\/\\>")
	str := regt.FindAllString(text, -1)

	return str
}

func reKeypArr(text []string, key string) []string {
	var tkeyarr []string
	for i := 0; i < len(text); i++ {
		tbkey := strings.Trim(text[i], "<"+key)
		tkey := strings.Trim(tbkey, "/>")
		tkeyarr = append(tkeyarr, tkey)
	}

	return tkeyarr
}

func reKeyp(text, key string) string {
	tbkey := strings.Trim(text, "<"+key)
	tkey := strings.Trim(tbkey, "/>")
	// keytag := "<" + key + ">" + tkey + "</" + key + ">"

	return tkey
}

func reKeyTag(text, key string) string {
	tbkey := strings.Trim(text, "<"+key)
	tkey := strings.Trim(tbkey, "/>")

	keytag := "<" + key + ">" + tkey + "</" + key + ">"

	// return tkey
	return keytag
}

func reSubKey(text, tagstr, trimName string) string {
	var keyarr []string

	keyp := findKeyp(tagstr, trimName)

	for i := 0; i < len(keyp); i++ {
		rekeyp := reKeyTag(keyp[i], trimName)
		// fmt.Println("rekeyp---", rekeyp)
		// keytag := "<" + key + ">" + rekeyp + "</" + key + ">"
		keyarr = append(keyarr, rekeyp)
	}

	fkeyp := findKTag(text, trimName)

	rekey := ReArrOne(text, fkeyp, keyarr)

	return rekey
}

func FindDiv(text string) []string {
	regc := regexp.MustCompile("(?U)\\<div[\\s\\S]*\\<\\/div\\>")
	restr := regc.FindAllString(text, -1)
	return restr
}

func FindLabel(text string) []string {
	regc := regexp.MustCompile("(?U)\\<div[\\s\\S]*\\<\\/div\\>")
	restr := regc.FindAllString(text, -1)
	return restr
}

func FindTag(rmap Map, name []string, text string) ([]string, string, [][]string, []string, [][]string, [][]string) {
	var (
		reCss        [][]string
		reJs         [][]string
		arrText      []string
		restr        [][]string
		openNotesarr []string
		retext       string
	)

	for i := 0; i < len(name); i++ {
		k := name[i]
		v := rmap[k]

		var tagstr []string
		if len(retext) > 0 {
			tagstr = findKTag(retext, k)
		} else {
			tagstr = findKTag(text, k)
		}

		if len(tagstr) > 0 {
			restr = append(restr, tagstr)
			var path string
			if strings.Contains(v.(string), "/") {
				path = "public/src/" + TrimQt(v.(string))
			} else {
				path = "public/" + TrimQt(v.(string))
			}

			readstr, err := Readfile(path)
			PrintErr(err)

			notes := TrimNotes(readstr)

			arecss, arejs := styScr(notes)

			if len(arecss) > 0 {
				reCss = append(reCss, arecss)
			}
			if len(arejs) > 0 {
				reJs = append(reJs, arejs)
			}
			////////////

			openNotes := TrimS(notes)

			openNotesarr = append(openNotesarr, openNotes)

			keyarr := Brace(openNotes)

			_, trimName, _ := ImpMap(openNotes)

			var headerArr []string
			var renotes string
			for t := 0; t < len(tagstr); t++ {
				var header string
				var soltNotes string

				slotOld := findSlot(openNotes)
				slotNew := findSlot(tagstr[t])

				if len(slotOld) == len(slotNew) {
					soltNotes = reSlot(openNotes, slotOld, slotNew)
				} else {
					soltNotes = TrimSlot(openNotes)
				}

				if len(trimName) > 0 {
					if len(soltNotes) > 0 {
						renotes = reSubKey(soltNotes, tagstr[t], trimName[0])
					} else {
						renotes = reSubKey(openNotes, tagstr[t], trimName[0])
					}
				}

				for i := 0; i < len(keyarr); i++ {
					fval := FindVal(keyarr[i], tagstr[t])
					if len(fval) > 0 {
						if len(trimName) > 0 {
							if strings.Contains(IsImp(trimName, keyarr[i], openNotes), "true") {
								if len(header) == 0 {
									header = strings.Replace(renotes, keyarr[i], fval[0], 1)
								} else {
									header = strings.Replace(header, keyarr[i], fval[0], 1)
								}
							} else {
								tkey := TrimB(keyarr[i])
								tfval := TrimTVal(tkey, fval)

								if len(header) == 0 {
									header = strings.Replace(renotes, keyarr[i], tfval[0], 1)
								} else {
									header = strings.Replace(header, keyarr[i], tfval[0], 1)
								}
							}

						} else {
							tkey := TrimB(keyarr[i])
							tfval := TrimTVal(tkey, fval)

							if len(header) == 0 {
								// header = strings.Replace(openNotes, keyarr[i], tfval[0], -1)
								header = strings.Replace(openNotes, keyarr[i], tfval[0], 1)
							} else {
								header = strings.Replace(header, keyarr[i], tfval[0], 1)
							}

						}

					}
				}
				//////
				amap, headName, _ := ImpMap(header)
				if len(headName) > 0 {
					arrt, _, retagstr, _, _, _ := FindTag(amap, headName, header)

					ret := ReArrOne(header, retagstr[0], arrt)
					headerArr = append(headerArr, ret)
				} else {
					headerArr = append(headerArr, header)
					arrText = append(arrText, header)
				}

			}
			if len(retext) > 0 {
				retext = ReArrOne(retext, tagstr, headerArr)
			} else {
				retext = ReArrOne(text, tagstr, headerArr)
			}

		}
	}
	return arrText, retext, restr, openNotesarr, reCss, reJs
}
