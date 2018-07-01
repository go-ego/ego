// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ego

import (
	"bytes"
	"html/template"
	"log"
)

func init() {
	log.SetFlags(0)
}

// IsDebugging returns true if the framework is running in debug mode.
// Use SetMode(ego.ReleaseMode) to disable debug mode.
func IsDebugging() bool {
	return egoMode == debugCode
}

func debugPrintRoute(httpMethod, absolutePath string, handlers HandlersChain) {
	if IsDebugging() {
		nuHandlers := len(handlers)
		handlerName := nameOfFunction(handlers.Last())
		debugPrint("%-6s %-25s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func debugPrintLoadTemplate(tmpl *template.Template) {
	if IsDebugging() {
		var buf bytes.Buffer
		for _, tmpl := range tmpl.Templates() {
			buf.WriteString("\t- ")
			buf.WriteString(tmpl.Name())
			buf.WriteString("\n")
		}
		debugPrint("Loaded HTML Templates (%d): \n%s\n", len(tmpl.Templates()), buf.String())
	}
}

func debugPrint(format string, values ...interface{}) {
	if IsDebugging() {
		log.Printf("[EGO-debug] "+format, values...)
	}
}

func debugPrintWARNINGDefault() {
	debugPrint(`[WARNING] Now Gin requires Go 1.8 or later and Go 1.x will be required soon.
	`)

	debugPrint(`[WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
	
	`)
}

func debugPrintWARNINGNew() {
	debugPrint(`[WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export EGO_MODE=release
 - using code:	ego.SetMode(ego.ReleaseMode)

`)
}

func debugPrintWARNINGSetHTMLTemplate() {
	debugPrint(`[WARNING] Since SetHTMLTemplate() is NOT thread-safe. It should only be called
at initialization. ie. before any route is registered or the router is listening in a socket:

	router := ego.Default()
	router.SetHTMLTemplate(template) // << good place

`)
}

func debugPrintError(err error) {
	if err != nil {
		debugPrint("[ERROR] %v\n", err)
	}
}
