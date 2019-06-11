package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v8"
)

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	log.Info(b, c.Request.Method, c.ContentType())
	log.Info(b.Name())
	return c.ShouldBindWith(obj, b)
}

// My own Error type that will help return my customized Error info
//  {"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, v := range errs {
			// can translate each error one at a time.
			//fmt.Println("gg",v.NameNamespace)
			if v.Param != "" {
				res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
			} else {
				res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
			}

		}
	} else {
		res.Errors["general"] = err.Error()
	}

	return res
}
