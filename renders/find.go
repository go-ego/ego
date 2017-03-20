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
	// "fmt"
	"regexp"
	"strings"
)

func FindIf(str string) []string {
	regt := regexp.MustCompile("(?U)if [\\s\\S]*\\{")
	restr := regt.FindAllString(str, -1)
	return restr
}

func FindS(str string) []string {
	regt := regexp.MustCompile("(?U)\\<style\\>[\\s\\S]*\\<\\/style\\>|\\<script\\>[\\s\\S]*\\<\\/script\\>")
	restr := regt.FindAllString(str, -1)
	return restr
}

func FindSty(str string) []string {
	regt := regexp.MustCompile("(?U)\\<style\\>[\\s\\S]*\\<\\/style\\>")
	restr := regt.FindAllString(str, -1)
	return restr
}

func FindScr(str string) []string {
	regt := regexp.MustCompile("(?U)\\<script\\>[\\s\\S]*\\<\\/script\\>")
	restr := regt.FindAllString(str, -1)
	return restr
}

func Brace(text string) []string {
	regc := regexp.MustCompile("(?U)\\{[^\\{|^\\}|^\\.|^\\<|^\\>]*\\}")
	restr := regc.FindAllString(text, -1)
	return restr
}

func KeyBrace(key, text string) []string {
	regc := regexp.MustCompile("(?U)\\<" + key + "\\>\\{[^\\{|^\\}|^\\.|^\\<|^\\>]*\\}\\<\\/" + key + "\\>")
	restr := regc.FindAllString(text, -1)
	return restr
}

func ImpBrace(imp, key, text string) []string {
	regc := regexp.MustCompile("(?U)\\<" + imp + "\\>[\\s\\S]*" + key + "[\\s\\S]*\\<\\/" + imp + "\\>")
	restr := regc.FindAllString(text, -1)
	return restr
}

func FindaKey(key []string) []string {
	var rekey []string

	for i := 0; i < len(key); i++ {
		tkey := TrimB(key[i])
		rekey = append(rekey, tkey)
	}

	return rekey
}

func FindKey(key []string) []string {
	var rekey []string
	for h := 0; h < len(key); h++ {
		bkey := Brace(key[h])

		for i := 0; i < len(bkey); i++ {
			tkey := TrimB(bkey[i])
			rekey = append(rekey, tkey)
		}
	}

	return rekey
}

func FindNoVal(text string) []string {
	regc := regexp.MustCompile(`(?U)\"[\s\S]*\"`)
	restr := regc.FindAllString(text, -1)
	return restr
}

func FindBVal(key, rstr string) []string {
	var streg string

	tkey := TrimB(key)

	if strings.Contains(rstr, "{") {
		streg = `(?U)` + tkey + "\\=\\{[\\s\\S]*\\}"
	} else {
		streg = `(?U)` + tkey + `\=\"[\s\S]*\"`
	}

	regc := regexp.MustCompile(streg)
	restr := regc.FindAllString(rstr, -1)

	return restr
}

func FindQVal(key, text string) []string {

	tkey := TrimB(key)

	streg := `(?U)` + tkey + `\=\"[\s\S]*\"`

	regc := regexp.MustCompile(streg)
	restr := regc.FindAllString(text, -1)

	return restr
}

func FindVal(key, text string) []string {
	var val []string
	bval := FindBVal(key, text)
	qval := FindQVal(key, text)
	for h := 0; h < len(bval); h++ {
		val = append(val, bval[h])
	}

	for i := 0; i < len(qval); i++ {
		val = append(val, qval[i])
	}

	return val
}

func FindArrVal(key []string) []string {
	var reval []string
	for h := 0; h < len(key); h++ {
		bval := Brace(key[h])

		for i := 0; i < len(bval); i++ {
			tval := TrimB(bval[i])
			reval = append(reval, tval)
		}
	}

	return reval
}
