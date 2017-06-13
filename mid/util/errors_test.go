// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	baseError := errors.New("test error")
	err := &Error{
		Err:  baseError,
		Type: ErrorTypePrivate,
	}
	assert.Equal(t, err.Error(), baseError.Error())
	assert.Equal(t, err.JSON(), Map{"error": baseError.Error()})

	assert.Equal(t, err.SetType(ErrorTypePublic), err)
	assert.Equal(t, err.Type, ErrorTypePublic)

	assert.Equal(t, err.SetMeta("some data"), err)
	assert.Equal(t, err.Meta, "some data")
	assert.Equal(t, err.JSON(), Map{
		"error": baseError.Error(),
		"meta":  "some data",
	})

	jsonBytes, _ := json.Marshal(err)
	assert.Equal(t, string(jsonBytes), "{\"error\":\"test error\",\"meta\":\"some data\"}")

	err.SetMeta(Map{
		"status": "200",
		"data":   "some data",
	})
	assert.Equal(t, err.JSON(), Map{
		"error":  baseError.Error(),
		"status": "200",
		"data":   "some data",
	})

	err.SetMeta(Map{
		"error":  "custom error",
		"status": "200",
		"data":   "some data",
	})
	assert.Equal(t, err.JSON(), Map{
		"error":  "custom error",
		"status": "200",
		"data":   "some data",
	})

	type customError struct {
		status string
		data   string
	}
	err.SetMeta(customError{status: "200", data: "other data"})
	assert.Equal(t, err.JSON(), customError{status: "200", data: "other data"})
}

func TestErrorSlice(t *testing.T) {
	errs := ErrorMsgs{
		{Err: errors.New("first"), Type: ErrorTypePrivate},
		{Err: errors.New("second"), Type: ErrorTypePrivate, Meta: "some data"},
		{Err: errors.New("third"), Type: ErrorTypePublic, Meta: Map{"status": "400"}},
	}

	assert.Equal(t, errs, errs.ByType(ErrorTypeAny))
	assert.Equal(t, errs.Last().Error(), "third")
	assert.Equal(t, errs.Errors(), []string{"first", "second", "third"})
	assert.Equal(t, errs.ByType(ErrorTypePublic).Errors(), []string{"third"})
	assert.Equal(t, errs.ByType(ErrorTypePrivate).Errors(), []string{"first", "second"})
	assert.Equal(t, errs.ByType(ErrorTypePublic|ErrorTypePrivate).Errors(), []string{"first", "second", "third"})
	assert.Empty(t, errs.ByType(ErrorTypeBind))
	assert.Empty(t, errs.ByType(ErrorTypeBind).String())

	assert.Equal(t, errs.String(), `Error #01: first
Error #02: second
     Meta: some data
Error #03: third
     Meta: map[status:400]
`)
	assert.Equal(t, errs.JSON(), []interface{}{
		Map{"error": "first"},
		Map{"error": "second", "meta": "some data"},
		Map{"error": "third", "status": "400"},
	})
	jsonBytes, _ := json.Marshal(errs)
	assert.Equal(t, string(jsonBytes), "[{\"error\":\"first\"},{\"error\":\"second\",\"meta\":\"some data\"},{\"error\":\"third\",\"status\":\"400\"}]")
	errs = ErrorMsgs{
		{Err: errors.New("first"), Type: ErrorTypePrivate},
	}
	assert.Equal(t, errs.JSON(), Map{"error": "first"})
	jsonBytes, _ = json.Marshal(errs)
	assert.Equal(t, string(jsonBytes), "{\"error\":\"first\"}")

	errs = ErrorMsgs{}
	assert.Nil(t, errs.Last())
	assert.Nil(t, errs.JSON())
	assert.Empty(t, errs.String())
}
