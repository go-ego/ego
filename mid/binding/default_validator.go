package binding

import (
	"reflect"
	"sync"

	"gopkg.in/go-playground/validator.v8"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ StructValidator = &defaultValidator{}

// ValidateStruct receives any kind of type, but only performed
// struct or pointer to struct type.
func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

func (v *defaultValidator) RegisterValidation(key string, fn validator.Func) error {
	v.lazyinit()
	return v.validate.RegisterValidation(key, fn)
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		config := &validator.Config{TagName: "binding"}
		v.validate = validator.New(config)
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
