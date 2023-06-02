package modules

// S6Value 对应 value
// 即语句中的占位符对应的值
type S6Value struct {
	// Value 值
	Value any
}

func (this S6Value) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	// sql 里加一个占位符
	p7s6Builder.sqlString.WriteByte('?')
	// sql 参数里加一个值
	p7s6Builder.f8AddParameter(this.Value)
	return nil
}

// F8NewS6Value 把输入转换成查询语句里占位符对应的值
func F8NewS6Value(input any) S6Value {
	return S6Value{Value: input}
}
