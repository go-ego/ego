// Copyright 2016 The go-ego Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-ego/ego/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package ego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ListFile(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

// Get http get
func Get(apiUrl string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiUrl)
	if err != nil {
		log.Printf("analytic url error:\r\n%v", err)
		return nil, err
	}

	// URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		log.Println("http get error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Post http post, params is url.Values type
func Post(apiUrl string, params url.Values, args ...int) (rs []byte, err error) {
	out := 1000
	if len(args) > 0 {
		out = args[0]
	}

	timeOut := time.Duration(out) * time.Millisecond

	c := &http.Client{
		Timeout: timeOut,
	}

	resp, err := c.PostForm(apiUrl, params)
	if err != nil {
		return nil, err
	}

	// fmt.Println("http:", resp)
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// API http api
func API(httpUrl string, paramMap Map, method ...string) (rs []byte, err error) {
	param := url.Values{}
	for k, v := range paramMap {
		param.Set(k, v.(string))
	}

	apiMethod := "post"
	if len(method) > 0 {
		apiMethod = method[0]
	}

	if apiMethod == "get" {
		rs, err = Get(httpUrl, param)
		return
	}

	rs, err = Post(httpUrl, param)
	return
}

// PostFile post file
func PostFile(filename, targetUrl, upParam string) (string, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// uploadfile
	fileWriter, err := bodyWriter.CreateFormFile(upParam, filename)
	if err != nil {
		log.Println("error writing to buffer")
		return "", err
	}

	// openfile
	fh, err := os.Open(filename)
	if err != nil {
		log.Println("error opening file")
		return "", err
	}

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return "", err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println(resp.Status)
	// fmt.Println(string(respBody))

	return string(respBody), nil
}

var (
	Url  url.Values = url.Values{}
	ajax int64
)

// TestRest test restful and return json
func (router *Engine) TestRest(httpUrl string, param url.Values) {

	listUrl := strings.Split(httpUrl, "/")
	lastUrl := listUrl[len(listUrl)-1]

	url := "/t_" + lastUrl

	router.GET(url, func(c *Context) {
		data, err := Post(httpUrl, param)

		if err != nil {
			log.Printf("Request failed, error message:\r\n%v", err)
		} else {
			var netReturn map[string]interface{}

			json.Unmarshal(data, &netReturn)

			reContent := netReturn["content"]

			c.JSON(200, reContent)
		}
	})
}

// TestJson test restful and return json
func (router *Engine) TestJson(httpUrl string, param url.Values, args ...string) {
	var content string

	if len(args) > 0 {
		content = args[0]
	} else {
		// content = "content"
		content = "data"
	}

	listUrl := strings.Split(httpUrl, "/")
	lastUrl := listUrl[len(listUrl)-1]

	url := "/t/" + lastUrl + "json"
	router.GET(url, func(c *Context) {
		data, err := Post(httpUrl, param)
		if err != nil {
			log.Printf("Request failed, error message:\r\n%v", err)
		} else {
			var netReturn map[string]interface{}
			// ffjson.Unmarshal(data, &netReturn)
			json.Unmarshal(data, &netReturn)
			reContent := netReturn[content]

			c.JSON(200, reContent)
		}
	})
}

// TestHtml test restful and show pretty in the browser
func (router *Engine) TestHtml(httpUrl string, paramMap Map, args ...string) {
	if ajax != 1 {
		router.StaticFile("/t/ajax", "./views/js/ajax.js")
	}
	ajax = 1

	param := url.Values{}
	for k, v := range paramMap {
		param.Set(k, v.(string))
	}
	listUrl := strings.Split(httpUrl, "/")
	lastUrl := listUrl[len(listUrl)-1]

	url := "/t/" + lastUrl

	if len(args) > 0 {
		router.TestJson(httpUrl, param, args[0])
	} else {
		router.TestJson(httpUrl, param)
	}

	router.GET(url, func(c *Context) {
		c.HTML(200, "json.html", Map{
			"test": httpUrl,
		})
	})
}

// TestFile test restful and show pretty in the browser
func (router *Engine) TestFile(httpUrl string, paramMap Map, filename, upParam string) {
	if ajax != 1 {
		router.StaticFile("/t/ajax", "./views/js/ajax.js")
	}
	ajax = 1

	var (
		url string
		i   int64
	)

	for k, v := range paramMap {
		i++
		if i == 1 {
			url += k + "=" + v.(string)
		} else {
			url += "&" + k + "=" + v.(string)
		}
	}

	confUrl := httpUrl + "?" + url

	fmt.Println("confUrl-------", confUrl)
	// confUrl := url.Values{}

	listUrl := strings.Split(httpUrl, "/")
	lastUrl := listUrl[len(listUrl)-1]

	htmlurl := "/t/" + lastUrl
	jsonurl := htmlurl + "json"

	router.GET(jsonurl, func(c *Context) {
		resp, err := PostFile(filename, confUrl, upParam)
		if err != nil {
			fmt.Println("err--------", err)
		}
		fmt.Println("resp---------", resp)

		c.JSON(200, resp)
	})

	router.GET(htmlurl, func(c *Context) {
		c.HTML(200, "json.html", Map{
			"test": httpUrl,
		})
	})

}
