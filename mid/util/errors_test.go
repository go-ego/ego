// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
	// "encoding/json"
	"errors"
	"testing"

	"github.com/go-ego/ego/mid/json"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	baseError := errors.New("test error")
	err := &Error{
		Err:  baseError,
		Type: ErrorTypePrivate,
	}
	assert.Equal(t, baseError.Error(), err.Error())
	assert.Equal(t, Map{"error": baseError.Error()}, err.JSON())

	assert.Equal(t, err.SetType(ErrorTypePublic), err)
	assert.Equal(t, ErrorTypePublic, err.Type)

	assert.Equal(t, err.SetMeta("some data"), err)
	assert.Equal(t, "some data", err.Meta)
	assert.Equal(t, Map{
		"error": baseError.Error(),
		"meta":  "some data",
	}, err.JSON())

	jsonBytes, _ := json.Marshal(err)
	assert.Equal(t, "{\"error\":\"test error\",\"meta\":\"some data\"}", string(jsonBytes))

	err.SetMeta(Map{
		"status": "200",
		"data":   "some data",
	})
	assert.Equal(t, Map{
		"error":  baseError.Error(),
		"status": "200",
		"data":   "some data",
	}, err.JSON())

	err.SetMeta(Map{
		"error":  "custom error",
		"status": "200",
		"data":   "some data",
	})
	assert.Equal(t, Map{
		"error":  "custom error",
		"status": "200",
		"data":   "some data",
	}, err.JSON())

	type customError struct {
		status string
		data   string
	}
	err.SetMeta(customError{status: "200", data: "other data"})
	assert.Equal(t, customError{status: "200", data: "other data"}, err.JSON())
}

func TestErrorSlice(t *testing.T) {
	errs := ErrorMsgs{
		{Err: errors.New("first"), Type: ErrorTypePrivate},
		{Err: errors.New("second"), Type: ErrorTypePrivate, Meta: "some data"},
		{Err: errors.New("third"), Type: ErrorTypePublic, Meta: Map{"status": "400"}},
	}

	assert.Equal(t, errs, errs.ByType(ErrorTypeAny))
	assert.Equal(t, "third", errs.Last().Error())
	assert.Equal(t, []string{"first", "second", "third"}, errs.Errors())
	assert.Equal(t, []string{"third"}, errs.ByType(ErrorTypePublic).Errors())
	assert.Equal(t, []string{"first", "second"}, errs.ByType(ErrorTypePrivate).Errors())
	assert.Equal(t, []string{"first", "second", "third"}, errs.ByType(ErrorTypePublic|ErrorTypePrivate).Errors())
	assert.Empty(t, errs.ByType(ErrorTypeBind))
	assert.Empty(t, errs.ByType(ErrorTypeBind).String())

	assert.Equal(t, `Error #01: first
Error #02: second
     Meta: some data
Error #03: third
     Meta: map[status:400]
`, errs.String())
	assert.Equal(t, []interface{}{
		Map{"error": "first"},
		Map{"error": "second", "meta": "some data"},
		Map{"error": "third", "status": "400"},
	}, errs.JSON())
	jsonBytes, _ := json.Marshal(errs)
	assert.Equal(t, "[{\"error\":\"first\"},{\"error\":\"second\",\"meta\":\"some data\"},{\"error\":\"third\",\"status\":\"400\"}]",
		string(jsonBytes))
	errs = ErrorMsgs{
		{Err: errors.New("first"), Type: ErrorTypePrivate},
	}
	assert.Equal(t, Map{"error": "first"}, errs.JSON())
	jsonBytes, _ = json.Marshal(errs)
	assert.Equal(t, "{\"error\":\"first\"}", string(jsonBytes))

	errs = ErrorMsgs{}
	assert.Nil(t, errs.Last())
	assert.Nil(t, errs.JSON())
	assert.Empty(t, errs.String())
}
