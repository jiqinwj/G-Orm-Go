package modules

// s6Operator 对应操作符
type s6Operator string

func (this s6Operator) String() string {
	return string(this)
}

const (
	c5OperatorEqual            s6Operator = "="
	c5OperatorGreaterThan      s6Operator = ">"
	c5OperatorGreaterEqualThan s6Operator = ">="
	c5OperatorLessThan         s6Operator = "<"
	c5OperatorLessEqualThan    s6Operator = "<="
	c5OperatorAND              s6Operator = "AND"
	c5OperatorOR               s6Operator = "OR"
	c5OperatorNOT              s6Operator = "NOT"
	c5OperatorLike             s6Operator = "LIKE"
	c5OperatorIn               s6Operator = "IN"
	c5OperatorAdd              s6Operator = "+"
	c5OperatorMulti            s6Operator = "*"
)
