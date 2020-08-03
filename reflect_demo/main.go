package main

import (
	"fmt"
	"reflect"
	"time"
)

// Cat ...
type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() // 值的类型
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func reflectStruct() {
	stu1 := student{
		Name:  "littie prince",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", filed.Name, filed.Index, filed.Type, filed.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
	var c = Cat{}
	reflectType(c)

	reflectValue(a)

	reflectSetValue2(&b)
	fmt.Println(b)

	fmt.Println("--------------------")
	reflectStruct()

	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t)
	}
}
