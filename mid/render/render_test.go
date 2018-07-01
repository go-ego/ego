// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
)

// TODO unit tests
// test errors

func TestRenderMsgPack(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
	}

	err := (MsgPack{data}).Render(w)

	assert.NoError(t, err)

	h := new(codec.MsgpackHandle)
	assert.NotNil(t, h)
	buf := bytes.NewBuffer([]byte{})
	assert.NotNil(t, buf)
	err = codec.NewEncoder(buf, h).Encode(data)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), string(buf.Bytes()))
	assert.Equal(t, w.Header().Get("Content-Type"), "application/msgpack; charset=utf-8")
}

func TestRenderJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
	}

	err := (JSON{data}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, "{\"foo\":\"bar\"}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestRenderIndentedJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
		"bar": "foo",
	}

	err := (IndentedJSON{data}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), "{\n    \"bar\": \"foo\",\n    \"foo\": \"bar\"\n}")
	assert.Equal(t, w.Header().Get("Content-Type"), "application/json; charset=utf-8")
}

type xmlmap map[string]interface{}

// Allows type H to be used with xml.Marshal
func (h xmlmap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func TestRenderXML(t *testing.T) {
	w := httptest.NewRecorder()
	data := xmlmap{
		"foo": "bar",
	}

	err := (XML{data}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), "<map><foo>bar</foo></map>")
	assert.Equal(t, w.Header().Get("Content-Type"), "application/xml; charset=utf-8")
}

func TestRenderRedirect(t *testing.T) {
	// TODO
}

func TestRenderData(t *testing.T) {
	w := httptest.NewRecorder()
	data := []byte("#!PNG some raw data")

	err := (Data{
		ContentType: "image/png",
		Data:        data,
	}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), "#!PNG some raw data")
	assert.Equal(t, w.Header().Get("Content-Type"), "image/png")
}

func TestRenderString(t *testing.T) {
	w := httptest.NewRecorder()

	err := (String{
		Format: "hola %s %d",
		Data:   []interface{}{"manu", 2},
	}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), "hola manu 2")
	assert.Equal(t, w.Header().Get("Content-Type"), "text/plain; charset=utf-8")
}

func TestRenderHTMLTemplate(t *testing.T) {
	w := httptest.NewRecorder()
	templ := template.Must(template.New("t").Parse(`Hello {{.name}}`))

	htmlRender := HTMLProduction{Template: templ}
	instance := htmlRender.Instance("t", map[string]interface{}{
		"name": "alexandernyquist",
	})

	err := instance.Render(w)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), "Hello alexandernyquist")
	assert.Equal(t, w.Header().Get("Content-Type"), "text/html; charset=utf-8")
}

func TestRenderReader(t *testing.T) {
	w := httptest.NewRecorder()

	body := "#!PNG some raw data"
	headers := make(map[string]string)
	headers["Content-Disposition"] = `attachment; filename="filename.png"`

	err := (Reader{
		ContentLength: int64(len(body)),
		ContentType:   "image/png",
		Reader:        strings.NewReader(body),
		Headers:       headers,
	}).Render(w)

	assert.NoError(t, err)
	assert.Equal(t, body, w.Body.String())
	assert.Equal(t, "image/png", w.HeaderMap.Get("Content-Type"))
	assert.Equal(t, strconv.Itoa(len(body)), w.HeaderMap.Get("Content-Length"))
	assert.Equal(t, headers["Content-Disposition"], w.HeaderMap.Get("Content-Disposition"))
}
