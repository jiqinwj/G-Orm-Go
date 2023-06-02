package modules

// S6WhereCondition 对应 where_condition
// SELECT Statement 里的 where_condition
// SELECT ... WHERE 后面的部分
// SELECT ... GROUP BY ... HAVING 后面的部分
// 可以通过嵌套组成复杂的查询条件
type S6WhereCondition s6BinaryExpression

func (this S6WhereCondition) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	expr := s6BinaryExpression(this)
	return expr.f8BuildExpression(p7s6Builder)
}

// F8And 与，左查询条件 `与` 右查询条件 => (`Id` = 11) AND (S6Column = 'aa')
func (this S6WhereCondition) F8And(p S6WhereCondition) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorAND,
		i9RightExpr: p,
	}
}

// F8Or 或，左查询条件 `或` 右查询条件 => (`Id` = 11) OR (S6Column = 'aa')
func (this S6WhereCondition) F8Or(p S6WhereCondition) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorOR,
		i9RightExpr: p,
	}
}

// F8Not 非，`非` 右查询条件 => NOT (`id` = 11)
// 注意 F8Not 条件只有操作符右边的查询条件
func F8Not(p S6WhereCondition) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  nil,
		s6Operator:  c5OperatorNOT,
		i9RightExpr: p,
	}
}
