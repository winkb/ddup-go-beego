package beeorm

import (
	"github.com/astaxie/beego/orm"
	"github.com/winkb/ddup-go-util/orm/where"
	"reflect"
	"strconv"
)

func ForBeeGoOrm(wr *where.Where, db orm.QuerySeter) orm.QuerySeter {

	switch wr.Con {
	case "=":
		return db.Filter(wr.Name, wr.Val)
	case "like":
		return db.Filter(wr.Name+"__contains", wr.Val)
	case "<>":
		v := reflect.ValueOf(wr.Val)
		switch v.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return db.FilterRaw(wr.Name, "<> "+strconv.FormatInt(v.Int(), 10))
		}

		return db.FilterRaw(wr.Name, "<> '"+v.String()+"'")
	case ">":
		return db.Filter(wr.Name+"__gt", wr.Val)
	case ">=":
		return db.Filter(wr.Name+"__gte", wr.Val)
	case "<":
		return db.Filter(wr.Name+"__lt", wr.Val)
	case "<=":
		return db.Filter(wr.Name+"__lte", wr.Val)
	}

	return db.Filter(wr.Name, wr.Val)
}
