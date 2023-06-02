package modules

// i9Expression 对应表达式
// 即 SELECT Statement 里的 expr。可以是：列、聚合函数、查询条件、值。
type i9Expression interface {
	// f8BuildExpression 构造表达式 SQL
	f8BuildExpression(p7s6Builder *s6QueryBuilder) error
}

// f8NewI9Expression 把输入转换成表达式
func f8NewI9Expression(input any) i9Expression {

	//fmt.Println("转换成什么类型", input)
	switch input.(type) {
	case i9Expression:
		// 如果是表达式，就断言一下丢回去
		return input.(i9Expression)
	default:
		//fmt.Println("普通类型", input)
		// 如果不是表达式，就转换成值
		return F8NewS6Value(input)
	}
}
