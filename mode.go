// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ego

import (
	"io"
	"os"

	"github.com/go-ego/ego/mid/binding"
)

const (
	// EnvEgoMode set the mode env
	EnvEgoMode = "EGO_MODE"

	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

var (
	// DefaultWriter is the default io.Writer used the Ego for debug output and
	// middleware output like Logger() or Recovery().
	// Note that both Logger and Recovery provides custom ways to configure their
	// output io.Writer.
	// To support coloring in Windows use:
	// 		import "github.com/mattn/go-colorable"
	// 		ego.DefaultWriter = colorable.NewColorableStdout()
	DefaultWriter      io.Writer = os.Stdout
	DefaultErrorWriter io.Writer = os.Stderr

	egoMode  = debugCode
	modeName = DebugMode
)

func init() {
	mode := os.Getenv(EnvEgoMode)
	SetMode(mode)
}

// SetMode set ego mode
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
