package main

import "fmt"

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
}

func loadInt(fileName string, v interface{}) (err error) {
	return nil
}

func main() {
	var mc MysqlConfig
	var rc RedisConfig
	err := loadInt("./conf.ini", &mc)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
	fmt.Println(rc.Port, rc.Port, rc.Password, rc.Database)
}
