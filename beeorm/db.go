package beeorm

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/winkb/ddup-go-beego/logs"
)

func Transcation(f func() error) (err error) {
	db := orm.NewOrm()

	err = db.Begin()
	if err != nil {
		logs.Log().Debug("事务begin出错:\n" + err.Error())
		return errors.New("事务失败")
	}

	err = f()

	// 业务处理失败，事务回滚
	if err != nil {
		errRoll := db.Rollback()

		if errRoll != nil {
			logs.Log().Debug("事务rollback出错:\n" + errRoll.Error())
		}

		return
	}

	err = db.Commit()

	// 事务提交成功
	if err == nil {
		return nil
	}

	logs.Log().Debug("事务commit出错:\n" + err.Error())
	_ = db.Rollback()

	return err
}
