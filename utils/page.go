package utils

import (
	"reflect"
)

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

/*
	count 总数量
	pageNo 第几页
	pageSize 每页的数量
	list  数据

*/
func PageUtil(count int, pageNo int, pageSize int) Page {
	if pageNo <= 0 {
		pageNo =1
	}

	if pageSize <=0 {
		pageSize =5
	}

	tp := count / pageSize
	if count % pageSize > 0 {
		tp = count / pageSize + 1
	}

	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp}
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
