package main

import (
	"fmt"
	"strconv"
)

// strconv

func f1() {
	str := "10000"
	str1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("paseint failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v %T\n", str1, int(str1))

	retInt, _ := strconv.Atoi(str)
	fmt.Println(retInt, "<---------")

	i := int32(2316)
	ret := string(i)
	fmt.Println(ret)
	ret3 := strconv.Itoa(int(i))
	fmt.Printf("%#v\n", ret3)

	// 将字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 将字符串中解析出浮点数
	floatStr := "1.2334"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)
}

func main() {
	f1()
}
