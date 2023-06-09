package modules

import (
	"context"
	"fmt"
)

// S6SelectBuilder 用于构造 SELECT 语句
type S6SelectBuilder[T any] struct {
	// s5select SELECT 后面的
	s5select []i9SelectExpr
	// i9from FROM 后面的
	i9from i9TableReference
	// s5where WHERE 后面的
	s5where []S6WhereCondition
	// s5GroupBy GROUP BY 后面的
	s5GroupBy []S6Column
	// s5having GROUP BY ... HAVING 后面的
	s5having []S6WhereCondition
	// s5OrderBy ORDER BY 后面的
	s5OrderBy []S6OrderBy
	// limit LIMIT 行数
	limit int
	// offset OFFSET 行数
	offset int

	i9Session I9Session
	s6QueryBuilder
}

func F8NewS6SelectBuilder[T any](i9Session I9Session) *S6SelectBuilder[T] {
	t4p7s6monitor := i9Session.f8GetS6Monitor()
	return &S6SelectBuilder[T]{
		i9Session: i9Session,
		s6QueryBuilder: s6QueryBuilder{
			s6Monitor: t4p7s6monitor,
			quote:     t4p7s6monitor.i9Dialect.f8GetQuoter(),
		},
	}
}

// F8Where 添加 where 子句
func (p7this *S6SelectBuilder[T]) F8Where(s5condition ...S6WhereCondition) *S6SelectBuilder[T] {
	if 0 >= len(s5condition) {
		return p7this
	}
	if nil == p7this.s5where {
		p7this.s5where = s5condition
		return p7this
	}
	p7this.s5where = append(p7this.s5where, s5condition...)
	return p7this
}

// F8First 执行查询获取一条数据，用映射关系
func (p7this *S6SelectBuilder[T]) F8First(i9ctx context.Context) (*T, error) {
	p7s6Context := &S6QueryContext{
		QueryType: "SELECT",
		i9Builder: p7this,
		p7s6Model: p7this.s6QueryBuilder.p7s6Model,
		p7s6Query: nil,
	}
	p7s6Result := f8DoFirst[T](i9ctx, p7this.i9Session, &p7this.s6Monitor, p7s6Context)
	if nil != p7s6Result.AnyResult {
		return p7s6Result.AnyResult.(*T), p7s6Result.I9Err
	}
	return nil, p7s6Result.I9Err
}

// f8BuildSelect 处理 SELECT 后面的
func (p7this *S6SelectBuilder[T]) f8BuildSelect() error {
	if 0 >= len(p7this.s5select) {
		p7this.sqlString.WriteByte('*')
		return nil
	}
	for i, t4value := range p7this.s5select {
		if 0 < i {
			p7this.sqlString.WriteByte(',')
		}
		err := t4value.f8BuildSelectExpr(&p7this.s6QueryBuilder)
		if nil != err {
			return err
		}
	}
	return nil
}

func (p7this *S6SelectBuilder[T]) f8BuildTableReference(reference i9TableReference) error {
	if nil == reference {
		p7this.f8WrapWithQuote(p7this.p7s6Model.TableName)
		return nil
	}
	return reference.f8BuildTableReference(&p7this.s6QueryBuilder)
}

// f8BuildWhereCondition 处理查询条件
func (p7this *s6QueryBuilder) f8BuildWhereCondition(s5p []S6WhereCondition) error {
	t4p := s5p[0]
	for i := 1; i < len(s5p); i++ {
		t4p = t4p.F8And(s5p[i])
	}
	return p7this.f8BuildExpression(t4p)
}

func (p7this *S6SelectBuilder[T]) F8BuildQuery() (*S6Query, error) {
	var err error = nil

	p7this.s6QueryBuilder.p7s6Model, err = p7this.s6Monitor.i9Registry.F8Get(new(T))
	if nil != err {
		return nil, err
	}

	p7this.sqlString.WriteString("SELECT ")

	err = p7this.f8BuildSelect()
	if nil != err {
		return nil, err
	}

	p7this.sqlString.WriteString(" FROM ")

	// 处理 FROM 后面的
	err = p7this.f8BuildTableReference(p7this.i9from)
	if nil != err {
		return nil, err
	}

	// 处理 where
	if 0 < len(p7this.s5where) {
		p7this.sqlString.WriteString(" WHERE ")
		err = p7this.f8BuildWhereCondition(p7this.s5where)
		if nil != err {
			return nil, err
		}
	}

	// 处理 group by
	if 0 < len(p7this.s5GroupBy) {
		p7this.sqlString.WriteString(" GROUP BY ")
		for i, t4value := range p7this.s5GroupBy {
			if 0 < i {
				p7this.sqlString.WriteByte(',')
			}
			err = t4value.f8BuildColumn(&p7this.s6QueryBuilder, false)
			if nil != err {
				return nil, err
			}
		}

		// 在有 group by 的情况下，才处理 having
		if 0 < len(p7this.s5having) {
			p7this.sqlString.WriteString(" HAVING ")
			err = p7this.f8BuildWhereCondition(p7this.s5having)
			if nil != err {
				return nil, err
			}
		}
	}

	// 处理 order by
	if 0 < len(p7this.s5OrderBy) {
		p7this.sqlString.WriteString(" ORDER BY ")
		for i, t4value := range p7this.s5OrderBy {
			if 0 < i {
				p7this.sqlString.WriteByte(',')
			}
			err = t4value.F8BuildOrderBy(&p7this.s6QueryBuilder)
			if nil != err {
				return nil, err
			}
		}
	}

	// 处理 limit offset
	if 0 < p7this.limit {
		p7this.sqlString.WriteString(" LIMIT ?")
		p7this.f8AddParameter(p7this.limit)
	}
	if 0 < p7this.offset {
		p7this.sqlString.WriteString(" OFFSET ?")
		p7this.f8AddParameter(p7this.offset)
	}

	p7this.sqlString.WriteByte(';')

	fmt.Println("打印构造器的sql---", p7this.sqlString.String())

	p7s6query := &S6Query{
		SQLString: p7this.sqlString.String(),
		S5Value:   p7this.s5Value,
	}

	return p7s6query, nil
}

func f8DoGetList[T any](i9ctx context.Context, i9Session I9Session, p7s6Monitor *s6Monitor, p7s6Context *S6QueryContext) *S6QueryResult {
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
		t4s5p7T := make([]*T, 0, 4)
		for p7SqlRows.Next() {
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

			t4s5p7T = append(t4s5p7T, t4p7T)
		}

		return &S6QueryResult{
			AnyResult: t4s5p7T,
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

// F8GetList 执行查询获取多条数据，用映射关系
func (p7this *S6SelectBuilder[T]) F8GetList(i9ctx context.Context) ([]*T, error) {
	p7s6Context := &S6QueryContext{
		QueryType: "SELECT",
		i9Builder: p7this,
		p7s6Model: p7this.s6QueryBuilder.p7s6Model,
		p7s6Query: nil,
	}
	p7s6Result := f8DoGetList[T](i9ctx, p7this.i9Session, &p7this.s6Monitor, p7s6Context)
	if nil != p7s6Result.AnyResult {
		return p7s6Result.AnyResult.([]*T), p7s6Result.I9Err
	}
	return nil, p7s6Result.I9Err
}
