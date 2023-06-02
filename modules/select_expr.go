package modules

// i9SelectExpr 对应查询表达式
// 即 SELECT Statement 里的 select_expr
// SELECT 语句 SELECT 后面 FROM 前面的那部分
type i9SelectExpr interface {
	// f8BuildSelectExpr 构造查询表达式
	f8BuildSelectExpr(p7s6Builder *s6QueryBuilder) error
	// f8GetFieldName 获取结构体属性名，构造子查询用的
	f8GetFieldName() string
	// f8GetAlias 获取别名，构造子查询用的
	f8GetAlias() string
}
