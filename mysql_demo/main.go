package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Go链接MySQL示例

var db *sql.DB

func initDB() (err error) {
	dsn := "viclilei:1198072529@tcp(182.61.13.234:3306)/viclilei"
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

// 查询
type user struct {
	id   int
	name string
	age  int
}

func queryOne(id int) user {
	var u1 user
	// 1.写查询单条记录的sql
	sqlStr := `select id, name, age from user where id=?;`
	// 2.执行(取到rowObj后必须调用Scan方法)
	// rowObj := db.QueryRow(sqlStr, 2) // 从连接池里拿一个连接出来去数据库查询单条记录
	// 3.拿到结果（释放连接）
	// rowObj.Scan(&u1.id, &u1.name, &u1.age)

	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)

	// 打印结果
	return u1
}

func queryMore(n int) {
	// 1.SQL语句
	sqlStr := `select id, name, age from user where id > ?;`
	// 2.执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
		return
	}
	// 一定要关闭
	defer rows.Close()

	// 循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("rows scan failed, err:%v\n", err)
			break
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

func insert() {
	sqlStr := `insert into user (name, age) values ("沙和尚", 4000)`
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 如果是插入数据的 操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get insert id failed, err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

func updateRow(name string, id int) {
	sqlStr := `update user set name= ? where id = ?;`
	ret, err := db.Exec(sqlStr, name, id)
	if err != nil {
		fmt.Printf("update user failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Println("update total rows:", n)
}

func deleteRow(id int) {
	sqlStr := `delete from user where id = ?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Println("delete total rows:", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into user (name, age) values (?, ?)`
	// 把SQL语句先发给MySQL预处理下
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"唐僧": 1000,
		"玉兔": 1200,
		"金角": 1300,
		"牛魔": 5000,
	}
	for k, v := range m {
		// 后续只需要传值
		stmt.Exec(k, v)
	}
}

func transactionDemo() {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("db begin failed, err:%v\n", err)
		return
	}
	// 执行多个SQL操作
	sqlStr1 := `update user set age = age + 2 where id = 1`
	sqlStr2 := `update user set age = age - 2 where id = 2`
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// 要回滚
		tx.Rollback()
		fmt.Println("the first step failed")
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Println("the second step failed")
		return
	}
	// 上面两步都执行成功，就提交本次事务
	err = tx.Commit()
	if err != nil {
		// 回滚
		tx.Rollback()
		fmt.Println("commit step failed")
		return
	}
	fmt.Println("all success")
}

func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name=%s", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	var u1 user
	err := db.QueryRow(sqlStr).Scan(&u1.id, &u1.name, &u1.age)
	if err != nil {
		fmt.Printf("select error, err:%v\n", err)
		return
	}
	fmt.Printf("user%#v\n", u1)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect mysql failed, err:%v\n", err)
		return
	}

	fmt.Println("connect success!")
	fmt.Println(queryOne(1))
	queryMore(0)
	// insert()
	// updateRow("沙僧", 3)
	// deleteRow(4)
	// prepareInsert()
	// transactionDemo()
	sqlInjectDemo("'悟空'")
	sqlInjectDemo("'xxx' or 1=1 #")
	sqlInjectDemo("'xxx' union select * from user #")
}
