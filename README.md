# Ego
<!--<img align="right" src="https://raw.githubusercontent.com/go-ego/ego/master/logo.jpg">-->
<!--<a href="https://circleci.com/gh/go-ego/ego/tree/dev"><img src="https://img.shields.io/circleci/project/go-ego/ego/dev.svg" alt="Build Status"></a>-->
[![Build Status](https://travis-ci.org/go-ego/ego.svg)](https://travis-ci.org/go-ego/ego)
[![codecov](https://codecov.io/gh/go-ego/ego/branch/master/graph/badge.svg)](https://codecov.io/gh/go-ego/ego)
[![CircleCI Status](https://circleci.com/gh/go-ego/ego.svg?style=shield)](https://circleci.com/gh/go-ego/ego)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ego/ego)](https://goreportcard.com/report/github.com/go-ego/ego)
[![GoDoc](https://godoc.org/github.com/go-ego/ego?status.svg)](https://godoc.org/github.com/go-ego/ego)
[![Release](https://github-release-version.herokuapp.com/github/go-ego/ego/release.svg?style=flat)](https://github.com/go-ego/ego/releases/latest)
[![Join the chat at https://gitter.im/go-ego/ego](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-ego/ego?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
<!--<a href="https://github.com/go-ego/ego/releases"><img src="https://img.shields.io/badge/%20version%20-%206.0.0%20-blue.svg?style=flat-square" alt="Releases"></a>-->
  
  >Ego is a full-stack web framework written in Go, lightweight and efficient front-end component solutions, based on [gin](https://github.com/gin-gonic/gin). The front-end is compiled, does not affect the back-end. 

This is a work in progress.

[简体中文](https://github.com/go-ego/ego/blob/master/README_zh.md)

## Contents
- [Docs](#docs)
- [Requirements](#requirements)
- [Installation](#installation)
- [Update](#update)
- [Build-tools](#build-tools)
- [Examples](#examples)
- [TestRestful](#testrestful)
- [Plans](#plans)
- [Donate](#donate)
- [Contributing](#contributing)
- [License](#license)

## Docs
  - [API Docs](https://github.com/go-ego/ego/blob/master/docs/doc.md) &nbsp;&nbsp;&nbsp;
  - [中文文档](https://github.com/go-ego/ego/blob/master/docs/doc_zh.md)
  - [GoDoc](https://godoc.org/github.com/go-ego/ego)

## Requirements:

Go Version ≥1.7

## Installation:
```
go get github.com/go-ego/ego
```
## Update:
```
go get -u github.com/go-ego/ego  
```
## [Build-tools](https://github.com/go-ego/re)
```
go get -u github.com/go-ego/re 
```
### re new 
To create a new Ego web application

```
$ re new my-webapp
```

### re run

To run the application we just created, you can navigate to the application folder and execute:
```
$ cd my-webapp && re run
```

## [Examples:](https://github.com/go-ego/ego/tree/master/examples)

#### [Router](https://github.com/go-ego/ego/blob/master/examples/ego/main.go)

```Go
package main

import (
	"github.com/go-ego/ego"
)

func main() {

	router := ego.Classic()
	ego.UseRenders()

	router.GlobHTML("views/html/*")

	parArr := [5]int{1, 2, 3, 4, 5}
	router.Ego("/head/", "head/head.html", ego.Map{
		"head":   "Test to load the HTML template",
		"parArr": parArr,
	})

	router.Run(":3100")
}
``` 

#### [icon.vgo](https://github.com/go-ego/ego/tree/master/examples/ego/public/src/icons)

```html
// pkg icon

<div class="icon">
	<i class="iconfont {vclass}" {node}></i>
	<p>{prpo}</p>
</div>

<style>

.header-left{
	float:left;
}

.header-right{
	float:right;
}

.iconfont {
  position: relative;
  font-size:24px
}
</style>

```

#### [head.vgo](https://github.com/go-ego/ego/blob/master/examples/ego/public/head/head.vgo)

```html
import (
	"icons"
	icon "icons/icon.vgo"
	)

<div class="head">
	<div>ego:{{.head}}</div>

	<icon>
		vclass={icon-share-to}
		node={ id="slot1"}
		prpo={node---1}
	</icon>

	<div>
		{{range .parArr}}
	        <p>arr::: {{.}}</p>
		{{end}}
	</div>

</div>

```
### Renderings:

<p align="center">
    <img src="https://github.com/go-ego/ego/blob/master/img/head.png" width="700" hight="500">
</p>

## TestRestful
```Go

package main

import (
	"github.com/go-ego/ego"
)

const httpUrl string = "http://127.0.0.1:3000"

func main() {

  router := ego.Classic()

  router.Static("/js", "./views/js")
  router.Static("/src", "./views/src")
  router.GlobHTML("views/html/*")

  strUrl := httpUrl + "/test/hlist"
  paramMap := ego.Map{
    "lon":  "10.1010101",
    "lat":  "20.202020",
    "type": "1",
  }
  router.TestHtml(strUrl, paramMap) // http url, http parameter, args (optional parameters): The default is "data".

  router.Run(":3100")
}

```
### Renderings:

<p align="center">
    <img src="https://github.com/go-ego/ego/blob/master/img/test.png" width="700" hight="500">
</p>

[More instructions](https://github.com/go-ego/RESTest)

## Plans
- Compression and merge css/js
- CSS Preprocessing
- Try supports MVVM and vdom
- Update web framework

## Donate
- Supporting ego, [buy me a coffee](https://github.com/go-vgo/buy-me-a-coffee).

## Contributing

- To contribute to Ego, please see [Contribution Guidelines](https://github.com/go-ego/ego/blob/master/CONTRIBUTING.md).
Fork -> Patch -> Push -> Test -> Pull Request.

- See [contributors page](https://github.com/go-ego/ego/graphs/contributors) for full list of contributors.

## License

Ego is primarily distributed under the terms of both the MIT license and the Apache License (Version 2.0), with portions covered by various BSD-like licenses.

See [LICENSE-APACHE](http://www.apache.org/licenses/LICENSE-2.0), [LICENSE-MIT](https://github.com/go-ego/ego/blob/master/LICENSE), and COPYRIGHT for details.