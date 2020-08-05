package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// MysqlConfig Mysql配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig Redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

// Config Config配置结构体
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadInt(fileName string, v interface{}) (err error) {
	t := reflect.TypeOf(v)
	// 传进来的v参数必须是指针类型（因为需要在函数中对其赋值修改）
	if t.Kind() != reflect.Ptr {
		// err = fmt.Errorf("data should be a pointer") // 格式化输出后返回一个error类型
		err = errors.New("data param should be a pointer") // 新创建一个error
		return
	}

	// 传入的v必须为结构体指针
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return
	}

	// 读取文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	// string(b) // 将文件内容转换成字符串
	// windows下为 "\r\n"  (按行读取数据)
	lineSlice := strings.Split(string(b), "\n")

	var structName string
	for idx, line := range lineSlice {
		//去掉首尾空格
		line = strings.TrimSpace(line)
		// 如果是注释跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "[") {
			// 节(section)判断
			if line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName在data中根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("find %s 对应的嵌套体%s\n", sectionName, structName)
				}
			}

		} else {
			// 分割兼职对
			// 异常情况判断
			index := strings.Index(line, "=")
			if index == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			rv := reflect.ValueOf(v)
			sValue := rv.Elem().FieldByName(structName) // get 嵌套结构体的值信息
			sType := sValue.Type()                      // get 嵌套结构体的类型信息
			structObj := rv.Elem().FieldByName(structName)
			if structObj.Kind() != reflect.Struct {
				err = fmt.Errorf("配置中的%s字段应该为一个结构体", structName)
				return
			}

			var fieldName string
			var fieldType reflect.Kind
			// 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) // tag信息是存储在类型信息中
				fieldType = field.Type.Kind()
				if field.Tag.Get("ini") == key {
					// find 对应的字段
					fieldName = field.Name
					break
				}
			}

			if len(fieldName) == 0 {
				// 结构体中找不到对应的字符
				continue
			}

			fieldObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName, fieldType)

			switch fieldType {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fieldObj.SetBool(valueBool)
			}
		}
	}
	return nil
}

func main() {
	var cfg Config
	err := loadInt("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg)
}
