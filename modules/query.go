package modules

import (
	"G-Orm-go/modules/metadata"
	"strings"
)

// 构造出来的查询语句和参数
type S6Query struct {
	// 带有占位符的Sql 语句
	SQLString string
	// Sql 语句中占位符对应的值
	S5Value []any
}

// 接口抽象:查询构造器
// Builder 设计模式
type I9QueryBuilder interface {
	//方法抽象：构造 s6Query
	F8BuildQuery() (*S6Query, error)
}

// 查询构造器
type s6QueryBuilder struct {
	//控制器 从s6DB 里面获取
	s6Monitor
	//查询对应的数据
	p7s6Model *metadata.S6Model
	// quote 这个东西 s6Monitor 里面有，拿到这里方便操作
	quote byte
	// sqlString 带有占位符的sql 语句
	sqlString strings.Builder
	// s5Value SQL 语句中占位符对应的参数
	s5Value []any
}

// f8AddParameter 添加占位符对应的参数
func (p7this *s6QueryBuilder) f8AddParameter(s5p ...any) {
	if nil == p7this.s5Value {
		p7this.s5Value = make([]any, 0, 8)
	}
	p7this.s5Value = append(p7this.s5Value, s5p...)
}

// f8WrapWithQuote 两边加引号
func (p7this *s6QueryBuilder) f8WrapWithQuote(name string) {
	p7this.sqlString.WriteByte(p7this.quote)
	p7this.sqlString.WriteString(name)
	p7this.sqlString.WriteByte(p7this.quote)
}

func (p7this *s6QueryBuilder) f8BuildExpression(expr i9Expression) error {
	if nil == expr {
		return nil
	}
	return expr.f8BuildExpression(p7this)
}
