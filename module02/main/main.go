// Created by Hisen at 2021/12/9.
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var db *sql.DB

var (
	UserNotFound = errors.New("用户未找到")
)

type User struct {
	Name string
	Age  int
}

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./tmp/hanxin.db")

	if err != nil {
		panic(err)
	}
	sqlTable := `CREATE TABLE IF NOT EXISTS users(id   int(11)  primary key,name varchar(191) null,age  integer(191) null);`
	_, err = db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}
}

func GetUserByName(name string) (*User, error) {
	user := &User{}
	rows := db.QueryRow("select * from users  where name=?", name)
	err := rows.Scan(&user)
	// 1.主动判断是否是因为没有数据而返回错误，而不是因为出现其他错误导致返回user为空
	// 2.sql.ErrNoRows虽然是错误，但是正常情况，属于未找到对应资源，所以需要将errNoRows与程序异常分开处理，并可
	// 将其直接映射到业务错误代码，直接返回到用户侧。
	// 3.sql.ErrNoRows使用无法打印堆栈信息,因为其是有builtin里的errors方法创建的，只有pkg/errors里才实现了打印堆栈的功能
	if err == sql.ErrNoRows {
		// 可直接返回自定义业务错误代码
		return nil, UserNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
func GetUserByNameNormal(name string) (*User, error) {
	user := &User{}
	rows := db.QueryRow("select * from users  where name=?", name)
	err := rows.Scan(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func main() {

	defer db.Close()

	// 方式一
	user, err := GetUserByName("hanxin")
	if err != nil {
		// Is方法可透传根因错误
		if errors.Is(err, UserNotFound) {
			// api接口，可返回404错误
			fmt.Printf("err:%T %v %+v\n", err, errors.Cause(err), errors.Cause(err))
			//  out:
			//  err:*errors.fundamental 用户未找到 用户未找到
			//	main.init
			//	/Users/hisen/workspace/allinmd/golang/GolangTrainingCamp/module01/main/main.go:14
			//	runtime.doInit
			//	/usr/local/go/src/runtime/proc.go:6498
			//	runtime.main
			//	/usr/local/go/src/runtime/proc.go:238
			//	runtime.goexit
			//	/usr/local/go/src/runtime/asm_amd64.s:1581

		} else {
			// api接口，可返回5xx错误
			fmt.Println("5xx")
		}
	} else {
		fmt.Println(user.Name)
		// todo business
	}

	// 方式二
	// 采用原始的处理方式，需要主动判断是否为ErrNoRows的情况,还需要导入sql包，对调用方来讲，不如第一种方式好。
	user2, err := GetUserByNameNormal("hanxin")
	if err != nil {
		fmt.Println(errors.Is(err, sql.ErrNoRows))
		// out: true
		fmt.Printf("err:%T %v %+v\n", errors.Cause(err), err, errors.Cause(err))
		// err:*errors.errorString sql: no rows in result set sql: no rows in result set
	} else {
		fmt.Println(user2.Name)
	}

}
