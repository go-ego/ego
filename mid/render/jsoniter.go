package render

import (
	"bytes"
	"net/http"

	"github.com/json-iterator/go"
)

var ijson = jsoniter.ConfigCompatibleWithStandardLibrary

type IJSON struct {
	Data interface{}
}

type IIndentedJSON struct {
	Data interface{}
}

type ISecureJSON struct {
	Prefix string
	Data   interface{}
}

type ISecureJSONPrefix string

// var jsonContentType = []string{"application/json; charset=utf-8"}

func (r IJSON) Render(w http.ResponseWriter) (err error) {
	if err = IWriteJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

func (r IJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func IWriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := ijson.Marshal(obj)
	if err != nil {
		return err
	}
	w.Write(jsonBytes)
	return nil
}

func (r IIndentedJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := ijson.MarshalIndent(r.Data, "", "    ")
	if err != nil {
		return err
	}
	w.Write(jsonBytes)
	return nil
}

func (r IIndentedJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func (r ISecureJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := ijson.Marshal(r.Data)
	if err != nil {
		return err
	}
	// if the jsonBytes is array values
	if bytes.HasPrefix(jsonBytes, []byte("[")) && bytes.HasSuffix(jsonBytes, []byte("]")) {
		w.Write([]byte(r.Prefix))
	}
	w.Write(jsonBytes)
	return nil
}

func (r ISecureJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
