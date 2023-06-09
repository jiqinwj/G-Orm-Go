package modules

import (
	"G-Orm-go/modules/metadata"
	"G-Orm-go/modules/result"
	"context"
	"database/sql"
	"fmt"
)

// F8S6DBOption 方法抽象：针对 S6DB 的 Option 设计模式
type F8S6DBOption func(*S6DB)

// S6DB 框架的数据库对象:封装真正的数据库对象
type S6DB struct {
	//真正的数据库对象
	p7s6SqlDB *sql.DB
	// 控制器
	s6Monitor
}

// 构造一个db
// F8NewS6DB 构造 S6DB
func F8NewS6DB(p7s6SqlDB *sql.DB, s5Option ...F8S6DBOption) *S6DB {
	p7s6DB := &S6DB{
		p7s6SqlDB: p7s6SqlDB,
		s6Monitor: s6Monitor{
			i9Registry:     metadata.F8NewI9Registry(),
			f8NewI9Result:  result.F8NewS6ResultUseUnsafe,
			i9Dialect:      S6MySQLDialect,
			s5f8Middleware: nil,
		},
	}
	// Option 设计模式
	for _, t4value := range s5Option {
		t4value(p7s6DB)
	}
	return p7s6DB
}

// F8DBWithMiddleware 设置中间件
func F8DBWithMiddleware(s5f8Middleware ...F8Middleware) F8S6DBOption {
	return func(p7s6DB *S6DB) {
		p7s6DB.s6Monitor.s5f8Middleware = s5f8Middleware
	}
}

// #### struct func ####

func (p7this *S6DB) f8GetS6Monitor() s6Monitor {
	return p7this.s6Monitor
}

func (p7this *S6DB) f8DoQueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	fmt.Println("query sql", query)
	return p7this.p7s6SqlDB.QueryContext(ctx, query, args...)
}

func (p7this *S6DB) f8DoEXECContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return p7this.p7s6SqlDB.ExecContext(ctx, query, args...)
}
