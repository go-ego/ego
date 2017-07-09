package render

import (
	"bytes"
	"net/http"

	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type IJSON struct {
	Data interface{}
}

type IIndentedJSON struct {
	Data interface{}
}

type SecureJSON struct {
	Prefix string
	Data   interface{}
}

type SecureJSONPrefix string

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
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	w.Write(jsonBytes)
	return nil
}

func (r IIndentedJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := json.MarshalIndent(r.Data, "", "    ")
	if err != nil {
		return err
	}
	w.Write(jsonBytes)
	return nil
}

func (r IIndentedJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func (r SecureJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := json.Marshal(r.Data)
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

func (r SecureJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
