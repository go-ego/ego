package render

import (
	"encoding/json"
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

type (
	FFJSON struct {
		Data interface{}
	}

	FFIndentedJSON struct {
		Data interface{}
	}
)

func (r FFJSON) Render(w http.ResponseWriter) (err error) {
	if err = FFWriteJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

func (r FFJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func FFWriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	ffjsonBytes, err := ffjson.Marshal(obj)
	if err != nil {
		return err
	}
	w.Write(ffjsonBytes)
	return nil
}

func (r FFIndentedJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := json.MarshalIndent(r.Data, "", "    ")
	if err != nil {
		return err
	}
	w.Write(jsonBytes)
	return nil
}

func (r FFIndentedJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
