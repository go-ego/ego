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
  
  >Ego 是一个基于 Gin 用 Go 编写的全栈 Web 框架，轻量级和高效的前端组件解决方案. 前端编译执行，不影响后端效率.

这是一项正在完善的工作.

QQ 群: 120563750

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
创建一个新的 Ego web 项目

```
$ re new my-webapp
```

### re run

运行我们创建的 web 项目, 你可以导航到应用程序文件夹并执行:
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
### 渲染效果:

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
  router.TestHtml(strUrl, paramMap) // http url, http 参数, args (可选参数): 默认为 "data"

  router.Run(":3100")
}

```
### 效果图:

<p align="center">
    <img src="https://github.com/go-ego/ego/blob/master/img/test.png" width="700" hight="500">
</p>

[更多说明](https://github.com/go-ego/RESTest)

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