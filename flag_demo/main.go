package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// 创建一个标志位参数
	var name string
	flag.StringVar(&name, "name", "悟空", "please input name")
	// name := flag.String("name", "悟空", "please input you name")
	age := flag.Int("age", 9000, "plase input you realist age")
	married := flag.Bool("married", false, "plase input you married info")
	mt := flag.Duration("mt", time.Second, "you married years")
	// 使用flag(先解析)
	flag.Parse()
	fmt.Println(name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*mt)
	fmt.Println("-------------------")

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}
