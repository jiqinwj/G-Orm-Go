package main

import (
	"G-Orm-go/modules"
	"G-Orm-go/modules/middleware"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	testSelectFirst()
}

func testSelectFirst() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := modules.F8NewS6DB(p7s6SQLDB, modules.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	sqlResult, err := modules.F8NewS6SelectBuilder[modules.S6APPUserModel](p7s6DB).
		F8Where(modules.F8NewS6Column("Id").F8Equal(1)).F8First(context.Background())
	fmt.Println(sqlResult, err)
}
