// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ego

import (
	"io"
	"os"

	"github.com/go-ego/ego/mid/binding"
)

const ENV_EGO_MODE = "EGO_MODE"

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter is the default io.Writer used the Ego for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
// 		import "github.com/mattn/go-colorable"
// 		ego.DefaultWriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout
var DefaultErrorWriter io.Writer = os.Stderr

var egoMode = debugCode
var modeName = DebugMode

func init() {
	mode := os.Getenv(ENV_EGO_MODE)
	SetMode(mode)
}

func SetMode(value string) {
	switch value {
	case DebugMode, "":
		egoMode = debugCode
	case ReleaseMode:
		egoMode = releaseCode
	case TestMode:
		egoMode = testCode
	default:
		panic("ego mode unknown: " + value)
	}
	if value == "" {
		value = DebugMode
	}
	modeName = value
}

func DisableBindValidation() {
	binding.Validator = nil
}

func EnableJsonDecoderUseNumber() {
	binding.EnableDecoderUseNumber = true
}

func Mode() string {
	return modeName
}
