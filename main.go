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
	//testSelectFirst()
	//testSelectGetList()
	//testSlowSQL()
	testInsert()
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

func testSelectGetList() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := modules.F8NewS6DB(p7s6SQLDB, modules.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	s5SQLResult, err := modules.F8NewS6SelectBuilder[modules.S6APPUserModel](p7s6DB).
		F8GetList(context.Background())
	fmt.Println(s5SQLResult, err)
	for _, t4value := range s5SQLResult {
		fmt.Println(t4value)
	}
}

func testSlowSQL() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := modules.F8NewS6DB(p7s6SQLDB, modules.F8DBWithMiddleware(
		middleware.SqlLogMiddlewareBuild(),
		middleware.SlowLogMiddlewareBuild(),
		middleware.SlowLogTriggerMiddlewareBuild(),
	))

	sqlResult, err := modules.F8NewS6SelectBuilder[modules.S6APPUserModel](p7s6DB).
		F8Where(modules.F8NewS6Column("Id").F8Equal(11)).F8First(context.Background())
	fmt.Println(sqlResult, err)
}

func testInsert() {
	p7s6SQLDB, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/golang_dev?charset=utf8")
	if nil != err {
		log.Fatalln(err)
	}
	p7s6DB := modules.F8NewS6DB(p7s6SQLDB, modules.F8DBWithMiddleware(middleware.SqlLogMiddlewareBuild()))

	sqlResult, err := modules.F8NewS6InsertBuilder[modules.S6APPUserModel](p7s6DB).
		F8SetEntity(&modules.S6APPUserModel{Id: 11, Name: "aa", Age: 22, Sex: 1}).
		F8EXEC(context.Background())
	fmt.Println(sqlResult, err)

	//sqlResult, err = modules.F8NewS6InsertBuilder[modules.S6APPUserModel](p7s6DB).
	//	F8SetEntity(
	//		&modules.S6APPUserModel{Id: 22, Name: "bb", Age: 33, Sex: 2},
	//		&modules.S6APPUserModel{Id: 33, Name: "cc", Age: 44, Sex: 1},
	//	).
	//	F8EXEC(context.Background())
	//fmt.Println(sqlResult, err)
	//
	//sqlResult, err = modules.F8NewS6InsertBuilder[modules.S6APPUserModel](p7s6DB).
	//	F8SetEntity(&modules.S6APPUserModel{Id: 11, Name: "aaaa", Age: 22, Sex: 1}).
	//	F8OnConflictBuilder().
	//	F8SetUpdate(modules.F8NewS6Column("Name").ToAssignment("aaaa")).
	//	F8EXEC(context.Background())
	//fmt.Println(sqlResult, err)
	//
	//sqlResult, err = modules.F8NewS6InsertBuilder[modules.S6APPUserModel](p7s6DB).
	//	F8SetEntity(&modules.S6APPUserModel{Id: 44, Name: "dd", Age: 55, Sex: 2}).
	//	F8OnConflictBuilder().
	//	F8SetUpdate(modules.F8NewS6Column("Name")).
	//	F8EXEC(context.Background())
	//fmt.Println(sqlResult, err)
	//
	//sqlResult, err = modules.F8NewS6InsertBuilder[modules.S6APPUserModel](p7s6DB).
	//	F8SetEntity(&modules.S6APPUserModel{Id: 44, Name: "dddd", Age: 55, Sex: 2}).
	//	F8OnConflictBuilder().
	//	F8SetUpdate(modules.F8NewS6Column("Name")).
	//	F8EXEC(context.Background())
	//fmt.Println(sqlResult, err)
}
