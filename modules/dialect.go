package modules

import "G-Orm-go/modules/internal"

var S6MySQLDialect I9Dialect = &s6MySQLDialect{}
var S6SQLite3Dialect I9Dialect = &s6SQLite3Dialect{}

type I9Dialect interface {
	//返回一个引号，引用列名，表名的引号
	f8GetQuoter() byte
	// f8BuildOnConflict 构造 On CONFLICT
	// MySQL 里的 UPSERT，SQLite3 里的 UPSERT 不一样
	f8BuildOnConflict(*s6QueryBuilder, *S6Conflict) error
}

// #### MySQL ####

type s6MySQLDialect struct {
}

func (p7this *s6MySQLDialect) f8GetQuoter() byte {
	return '`'
}

func (p7this *s6MySQLDialect) f8BuildOnConflict(p7s6Builder *s6QueryBuilder, p7s6Conflict *S6Conflict) error {
	p7s6Builder.sqlString.WriteString(" ON DUPLICATE KEY UPDATE ")
	for i, t4value := range p7s6Conflict.S5Assignment {
		if 0 < i {
			p7s6Builder.sqlString.WriteByte(',')
		}
		switch t4value2 := t4value.(type) {
		case S6Column:
			p7s6ModelField, ok := p7s6Builder.p7s6Model.M3FieldToColumn[t4value2.fieldName]
			if !ok {
				return internal.F8NewErrUnknownField(t4value2.fieldName)
			}
			p7s6Builder.f8WrapWithQuote(p7s6ModelField.ColumnName)
			p7s6Builder.sqlString.WriteString("=VALUES(")
			p7s6Builder.f8WrapWithQuote(p7s6ModelField.ColumnName)
			p7s6Builder.sqlString.WriteString(")")
		case S6Assignment:
			err := t4value2.f8BuildAssignment(p7s6Builder)
			if nil != err {
				return err
			}
		default:
			return internal.NewErrUnsupportedExpressionType(t4value2)
		}
	}
	return nil
}

type s6SQLite3Dialect struct {
}

func (p7this *s6SQLite3Dialect) f8GetQuoter() byte {
	return '`'
}

func (p7this *s6SQLite3Dialect) f8BuildOnConflict(p7s6Builder *s6QueryBuilder, p7s6Conflict *S6Conflict) error {
	p7s6Builder.sqlString.WriteString(" ON CONFLICT ")

	if 0 < len(p7s6Conflict.S5ConflictColumn) {
		p7s6Builder.sqlString.WriteByte('(')
		for i, t4value := range p7s6Conflict.S5ConflictColumn {
			if 0 < i {
				p7s6Builder.sqlString.WriteByte(',')
			}
			err := t4value.f8BuildColumn(p7s6Builder, false)
			if nil != err {
				return err
			}
		}
		p7s6Builder.sqlString.WriteByte(')')
	}

	p7s6Builder.sqlString.WriteString(" DO UPDATE SET ")

	for i, t4value := range p7s6Conflict.S5Assignment {
		if 0 < i {
			p7s6Builder.sqlString.WriteByte(',')
		}
		switch t4value2 := t4value.(type) {
		case S6Column:
			p7s6ModelField, ok := p7s6Builder.p7s6Model.M3FieldToColumn[t4value2.fieldName]
			if !ok {
				return internal.F8NewErrUnknownField(t4value2.fieldName)
			}
			p7s6Builder.f8WrapWithQuote(p7s6ModelField.ColumnName)
			p7s6Builder.sqlString.WriteString("=excluded.")
			p7s6Builder.f8WrapWithQuote(p7s6ModelField.ColumnName)
		case S6Assignment:
			err := t4value2.f8BuildAssignment(p7s6Builder)
			if nil != err {
				return err
			}
		default:
			return internal.NewErrUnsupportedExpressionType(t4value2)
		}
	}

	return nil
}
