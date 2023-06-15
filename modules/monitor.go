package modules

import (
	"G-Orm-go/modules/internal"
	"G-Orm-go/modules/metadata"
	"G-Orm-go/modules/result"
	"context"
	"database/sql"
	"fmt"
)

// 控制器，控制和数据库有关的抽象的实现

type s6Monitor struct {
	// 元数据注册中心
	i9Registry metadata.I9Registry
	// 处理"用数据库返回的查询结构构造结构体"
	f8NewI9Result result.F8NewI9Result
	//处理方言
	i9Dialect I9Dialect
	// s5f8Middleware 中间件
	s5f8Middleware []F8Middleware
}

func f8DoFirst[T any](i9ctx context.Context, i9Session I9Session, p7s6Monitor *s6Monitor, p7s6Context *S6QueryContext) *S6QueryResult {
	var f8HandleFunc F8MiddlewareHandle = func(ctx context.Context, p7s6Context *S6QueryContext) *S6QueryResult {
		// 查询构造器构造查询。注意，这里用的是查询上下文的构造方法，防止查询语句重复构造。
		p7s6Query, err := p7s6Context.F8CTXBuildQuery()
		if nil != err {
			return &S6QueryResult{
				I9Err: err,
			}
		}
		// 执行查询
		p7SqlRows, err := i9Session.f8DoQueryContext(ctx, p7s6Query.SQLString, p7s6Query.S5Value...)
		if nil != err {
			return &S6QueryResult{
				I9Err: err,
			}
		}

		// 处理数据库返回的查询结果
		if !p7SqlRows.Next() {
			return &S6QueryResult{
				I9Err: internal.ErrSqlRowsIsEmpty,
			}
		}

		// new 一个类型 T 的变量
		t4p7T := new(T)
		// 获取类型 T 对应的映射模型
		t4s6model, err := i9Session.f8GetS6Monitor().i9Registry.F8Get(t4p7T)
		if nil != err {
			return &S6QueryResult{
				I9Err: err,
			}
		}

		// 用数据库返回的查询结果构造结构体
		t4result := i9Session.f8GetS6Monitor().f8NewI9Result(t4p7T, t4s6model)
		err = t4result.F8SetField(p7SqlRows)

		return &S6QueryResult{
			AnyResult: t4p7T,
			I9Err:     err,
		}
	}

	// 中间件套娃
	for i := len(p7s6Monitor.s5f8Middleware) - 1; 0 <= i; i-- {
		f8HandleFunc = p7s6Monitor.s5f8Middleware[i](f8HandleFunc)
	}
	// 执行套娃
	p7s6Result := f8HandleFunc(i9ctx, p7s6Context)

	// 从中间件的 S6QueryResult 里面把结果捞出来
	return p7s6Result
}

func f8DoEXEC(ctx context.Context, i9Session I9Session, p7s6Monitor *s6Monitor, p7s6Context *S6QueryContext) S6Result {
	var f8HandleFunc F8MiddlewareHandle = func(ctx context.Context, p7s6Context *S6QueryContext) *S6QueryResult {
		// 查询构造器构造查询
		//p7s6Query, err := p7s6Context.i9Builder.F8BuildQuery()
		p7s6Query, err := p7s6Context.F8CTXBuildQuery()
		if nil != err {
			return &S6QueryResult{
				I9Err: err,
			}
		}
		// 执行查询
		sqlResult, err2 := i9Session.f8DoEXECContext(ctx, p7s6Query.SQLString, p7s6Query.S5Value...)
		fmt.Println("error", err2)
		return &S6QueryResult{
			AnyResult: sqlResult,
			I9Err:     err2,
		}
	}

	// 中间件套娃
	for i := len(p7s6Monitor.s5f8Middleware) - 1; 0 <= i; i-- {
		f8HandleFunc = p7s6Monitor.s5f8Middleware[i](f8HandleFunc)
	}
	// 执行套娃
	p7s6Result := f8HandleFunc(ctx, p7s6Context)

	// 从中间件的 S6QueryResult 里面把结果捞出来
	var i9SQLResult sql.Result = nil
	if nil != p7s6Result.AnyResult {
		i9SQLResult = p7s6Result.AnyResult.(sql.Result)
	}
	return S6Result{I9SQLResult: i9SQLResult, I9Err: p7s6Result.I9Err}
}
