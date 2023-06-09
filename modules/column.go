package modules

import "G-Orm-go/modules/internal"

// 即各个 Statement 里的col_name ,语句中表示【表，JOIN,子查询】中列的部分
type S6Column struct {
	// i9From 列对应的[表、JOIN、子查询]
	i9From i9TableReference
	// fieldName 结构体属性名
	fieldName string
	// alias 数据库列名的别名
	alias string
}

// #### func ####

func F8NewS6Column(name string) S6Column {
	return S6Column{
		i9From:    nil,
		fieldName: name,
		alias:     "",
	}
}

func (this S6Column) F8Equal(p any) S6WhereCondition {
	return S6WhereCondition{
		i9LeftExpr:  this,
		s6Operator:  c5OperatorEqual,
		i9RightExpr: f8NewI9Expression(p),
	}
}

func (this S6Column) f8BuildSelectExpr(p7s6qb *s6QueryBuilder) error {
	return this.f8BuildColumn(p7s6qb, true)
}

// f8BuildColumn 构造列 SQL
// p7s6Builder 查询构造器
// isUseAlias 用不用别名
func (this S6Column) f8BuildColumn(p7s6Builder *s6QueryBuilder, isUseAlias bool) error {
	// 初始化列名为空，默认找不到列
	var columnName string = ""
	var err error = internal.F8NewErrUnknownField(this.fieldName)

	// 处理表
	if nil != this.i9From {
		columnName, err = this.i9From.f8CheckColumn(p7s6Builder, this)
		// 处理表的别名
		alies := this.i9From.f8GetTableReferenceAlies()
		if "" != alies {
			p7s6Builder.f8WrapWithQuote(alies)
			p7s6Builder.sqlString.WriteByte('.')
		}
	}
	// 上面的逻辑没找到属性，就走默认逻辑，再校验一次
	if nil != err {
		// 校验属性存不存在，存在转换成数据库列名
		p7s6ModelField, ok := p7s6Builder.p7s6Model.M3FieldToColumn[this.fieldName]
		if !ok {
			return internal.F8NewErrUnknownField(this.fieldName)
		}
		columnName = p7s6ModelField.ColumnName
	}
	p7s6Builder.f8WrapWithQuote(columnName)
	// 处理列的别名
	if isUseAlias && "" != this.alias {
		p7s6Builder.sqlString.WriteString(" AS ")
		p7s6Builder.f8WrapWithQuote(this.alias)
	}
	return nil
}

// f8BuildAssignment 赋值语句，对应，列 = 列
func (this S6Column) f8BuildAssignment(*s6QueryBuilder) error { return nil }

func (this S6Column) f8GetFieldName() string {
	return this.fieldName
}

func (this S6Column) f8GetAlias() string {
	return this.alias
}

func (this S6Column) f8BuildExpression(p7s6Builder *s6QueryBuilder) error {
	return this.f8BuildColumn(p7s6Builder, false)
}

// ToAssignment 赋值语句，对应，列 = 表达式
func (this S6Column) ToAssignment(input any) S6Assignment {
	i9Expr, ok := input.(i9Expression)
	if !ok {
		i9Expr = S6Value{Value: input}
	}
	return S6Assignment{
		s6Column: this,
		i9Expr:   i9Expr,
	}
}
