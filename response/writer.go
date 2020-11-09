package response

import (
	"github.com/astaxie/beego"
)

var codeSuccess = 200
var codeFail = 606

func SetResponseDefaultCode(success, fail int) {
	codeSuccess = success
	codeFail = fail
}

type ControllerWriter struct {
	beego.Controller
	code int
}

func (c *ControllerWriter) Success(data interface{}) {
	r := &responseJson{}

	r.Code = codeSuccess
	r.Data = data

	if data == nil {
		r.Data = "OK"
	}

	write(c, r)
}

func (c *ControllerWriter) Code(code int) *ControllerWriter {
	c.code = code

	return c
}

func (c *ControllerWriter) FailError(err error) {
	r := &responseJson{}

	r.Data = err.Error()
	r.Message = err.Error()
	r.Code = codeFail

	write(c, r)
}

func (c *ControllerWriter) Fail(message string) {
	r := &responseJson{}

	r.Data = message
	r.Message = message
	r.Code = codeFail

	write(c, r)
}

func write(c *ControllerWriter, r *responseJson) {
	if c.code != 0 {
		r.Code = c.code
	}

	c.Controller.Data["json"] = r.Data

	c.Ctx.ResponseWriter.WriteHeader(r.Code)

	c.ServeJSON()
}
