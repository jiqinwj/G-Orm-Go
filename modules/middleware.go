package modules

import (
	"G-Orm-go/modules/metadata"
	"context"
)

// S6QueryContext 查询上下文，给中间件用的
type S6QueryContext struct {
	// 查询类型[SELECT、INSERT、UPDATE、DELETE]
	QueryType string
	// 查询构造器[S6SelectBuilder、S6InsertBuilder、S6UpdateBuilder、S6DeleteBuilder]
	// 如果想在中间件里使用查询构造器，需要先进行类型断言
	i9Builder I9QueryBuilder
	// p7s6Model 映射模型
	p7s6Model *metadata.S6Model
	// 构造出来的查询语句和参数，如果中间件需要用就可以提前构造
	p7s6Query *S6Query
}

// S6QueryResult 查询结果，给中间件用的
type S6QueryResult struct {
	// AnyResult 查询结果
	// 不同的查询，结果类型不一样，需要进行类型断言
	// S6SELECT.First() => *T
	// S6SELECT.List() => []*T
	// S6SELECT.Get() => map[string]any
	// INSERT.EXEC() => sql.Result
	// UPDATE.EXEC() => sql.Result
	// DELETE.EXEC() => sql.Result
	AnyResult any
	I9Err     error
}

// F8CTXBuildQuery 用查询构造器构造查询
// 查询构造器不会主动构造查询，这玩意是在最后要执行查询的时候构造的
func (p7this *S6QueryContext) F8CTXBuildQuery() (*S6Query, error) {
	var err error = nil

	if nil == p7this.p7s6Query {
		p7this.p7s6Query, err = p7this.i9Builder.F8BuildQuery()
	}
	return p7this.p7s6Query, err
}

// F8MiddlewareHandle 中间件的处理方法
type F8MiddlewareHandle func(ctx context.Context, p7s6Context *S6QueryContext) *S6QueryResult

// F8Middleware 中间件
type F8Middleware func(next F8MiddlewareHandle) F8MiddlewareHandle
