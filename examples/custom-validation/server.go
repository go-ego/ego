package main

import (
	"net/http"
	"reflect"
	"time"

	"github.com/go-ego/ego"
	"github.com/go-ego/ego/mid/binding"
	validator "gopkg.in/go-playground/validator.v8"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func main() {
	route := ego.Default()
	// binding.Validator.RegisterValidation("bookabledate", bookableDate)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	route.GET("/bookable", getBookable)
	route.Run(":8085")
}

func getBookable(c *ego.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, ego.Map{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, ego.Map{"error": err.Error()})
	}
}
