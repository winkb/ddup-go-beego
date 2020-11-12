package beeorm

import (
	"github.com/astaxie/beego/orm"
	"github.com/winkb/ddup-go-util/orm/where"
	"reflect"
	"strconv"
)

func ScoreWhere(qSelect orm.QuerySeter, wh []*where.Where) orm.QuerySeter {
	for _, wr := range wh {
		qSelect = ForBeeGoOrm(wr, qSelect)
	}

	return qSelect
}

func ScopeWhere(wh []*where.Where, q orm.QueryBuilder) (orm.QueryBuilder, []interface{}) {
	vals := []interface{}{}

	for _, v := range wh {
		q = ForBeeGoBuilder(v, q)
		vals = append(vals, v.Val)
	}

	return q, vals
}

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

func ForBeeGoBuilder(wr *where.Where, db orm.QueryBuilder) orm.QueryBuilder {

	switch wr.Con {
	case "=":
		return db.Where(wr.Name + " = ?")
	case "like":
		return db.Where(wr.Name + " like %?%")
	case "<>":
		return db.Where(wr.Name + " <> ?")
	case ">":
		return db.Where(wr.Name + " > ?")
	case ">=":
		return db.Where(wr.Name + " >= ?")
	case "<":
		return db.Where(wr.Name + " < ?")
	case "<=":
		return db.Where(wr.Name + " <= ?")
	}

	return db.Where(wr.Name + " = ?")
}
