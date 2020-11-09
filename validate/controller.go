package validate

import (
	"ddup-go-beego/logs"
	"github.com/astaxie/beego/validation"
)

type ControllerValidate struct {
}

func (c *ControllerValidate) Validate(dtoObj interface{}) error {
	valid := validation.Validation{}
	b, err := valid.Valid(dtoObj)
	if err != nil {
		logs.LogOnError(err, "验证失败")
		return err
	}

	if b {
		return nil
	}

	for _, err := range valid.Errors {
		logs.LogOnError(err, "数据验证未通过")
		return err
	}

	return nil
}
