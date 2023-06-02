package modules

import "G-Orm-go/modules/internal"

// i9Assignment 标记接口，对应 INSERT 和 UPDATE 的赋值语句
// 即 INSERT Statement 和 UPDATE Statement 里的 assignment
type i9Assignment interface {
	// 构造赋值语句的 SQL
	f8BuildAssignment(*s6QueryBuilder) error
}

// S6Assignment 赋值语句
type S6Assignment struct {
	// s6Column 列
	s6Column S6Column
	// i9Expr 表达式
	i9Expr i9Expression
}

// f8BuildAssignment 构建赋值语句 SQL，对应，列 = 表达式，这种
func (this S6Assignment) f8BuildAssignment(p7s6Builder *s6QueryBuilder) error {
	p7s6ModelField, ok := p7s6Builder.p7s6Model.M3FieldToColumn[this.s6Column.fieldName]
	if !ok {
		return internal.F8NewErrUnknownField(this.s6Column.fieldName)
	}
	p7s6Builder.f8WrapWithQuote(p7s6ModelField.ColumnName)
	p7s6Builder.sqlString.WriteByte('=')
	err := p7s6Builder.f8BuildExpression(this.i9Expr)
	if nil != err {
		return err
	}
	return nil
}
