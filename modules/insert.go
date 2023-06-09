package modules

import (
	"G-Orm-go/modules/internal"
	"G-Orm-go/modules/metadata"
	"context"
	"database/sql"
)

// S6InsertBuilder INSERT 查询构造器
type S6InsertBuilder[T any] struct {
	// s5p7Entity 代表要插入的数据，解析它得到原数据
	s5p7Entity []*T
	// s5FieldName 结构体属性名，代表要插入的字段
	s5FieldName []string
	// p7s6Conflict ON CONFLICT 后面的
	// MySQL 中是 ON DUPLICATE KEY 后面的
	// SQLite3 中是 ON CONFLICT 后面的
	p7s6Conflict *S6Conflict

	i9Session I9Session
	s6QueryBuilder
}

func F8NewS6InsertBuilder[T any](i9Session I9Session) *S6InsertBuilder[T] {
	p7s6monitor := i9Session.f8GetS6Monitor()
	return &S6InsertBuilder[T]{
		i9Session: i9Session,
		s6QueryBuilder: s6QueryBuilder{
			s6Monitor: p7s6monitor,
			quote:     p7s6monitor.i9Dialect.f8GetQuoter(),
		},
	}
}

func (p7this *S6InsertBuilder[T]) F8SetEntity(s5Entity ...*T) *S6InsertBuilder[T] {
	if 0 >= len(s5Entity) {
		return p7this
	}
	if nil == p7this.s5p7Entity {
		p7this.s5p7Entity = s5Entity
		return p7this
	}
	p7this.s5p7Entity = append(p7this.s5p7Entity, s5Entity...)
	return p7this
}

func (p7this *S6InsertBuilder[T]) F8BuildQuery() (*S6Query, error) {
	var err error = nil

	p7this.s6QueryBuilder.p7s6Model, err = p7this.s6Monitor.i9Registry.F8Get(p7this.s5p7Entity[0])
	if nil != err {
		return nil, err
	}

	// 处理要插入的表
	p7this.s6QueryBuilder.sqlString.WriteString("INSERT INTO ")
	p7this.f8WrapWithQuote(p7this.s6QueryBuilder.p7s6Model.TableName)

	// 处理要插入的字段
	p7this.s6QueryBuilder.sqlString.WriteByte('(')
	// 如果没有设置要插入的字段，默认映射模型里面的全要
	s5p7s6ModelField := p7this.s6QueryBuilder.p7s6Model.S5P7S6ModelField
	if 0 != len(p7this.s5FieldName) {
		// 设置了要插入的字段，这里就要从映射模型里挑一遍，重新赋值
		s5p7s6ModelField = make([]*metadata.S6ModelField, 0, len(p7this.s5FieldName))
		// 检查一下设置的要插入的字段在不在映射模型里面
		for _, t4FieldName := range p7this.s5FieldName {
			t4p7s6ModelField, ok := p7this.p7s6Model.M3FieldToColumn[t4FieldName]
			if !ok {
				return nil, internal.F8NewErrUnknownField(t4FieldName)
			}
			s5p7s6ModelField = append(s5p7s6ModelField, t4p7s6ModelField)
		}
	}

	// 切片大小 = 要插入的字段数量 * 要插入的数据数量 + 1
	// ON CONFLICT 语句（MySQL UPSERT 语句）会传递额外的参数，所以要 +1
	p7this.s6QueryBuilder.s5Value = make([]any, 0, len(s5p7s6ModelField)*(len(p7this.s5p7Entity)+1))
	for i, t4value := range s5p7s6ModelField {
		if 0 < i {
			p7this.s6QueryBuilder.sqlString.WriteByte(',')
		}
		p7this.f8WrapWithQuote(t4value.ColumnName)
	}

	p7this.s6QueryBuilder.sqlString.WriteString(") VALUES")

	// 处理要插入的数据
	for i, t4value := range p7this.s5p7Entity {
		if 0 < i {
			p7this.s6QueryBuilder.sqlString.WriteByte(',')
		}
		// 通过反射拿到结构体属性的值
		t4i9result := p7this.f8NewI9Result(t4value, p7this.p7s6Model)
		p7this.s6QueryBuilder.sqlString.WriteByte('(')
		for j, t4value2 := range s5p7s6ModelField {
			if 0 < j {
				p7this.s6QueryBuilder.sqlString.WriteByte(',')
			}
			p7this.s6QueryBuilder.sqlString.WriteByte('?')
			t4EntityValue, err2 := t4i9result.F8GetField(t4value2.FieldName)
			if err2 != nil {
				return nil, err2
			}
			p7this.f8AddParameter(t4EntityValue)
		}

		p7this.s6QueryBuilder.sqlString.WriteByte(')')
	}

	// 处理 ON CONFLICT 部分
	if nil != p7this.p7s6Conflict {
		err = p7this.s6QueryBuilder.s6Monitor.i9Dialect.f8BuildOnConflict(&p7this.s6QueryBuilder, p7this.p7s6Conflict)
		if err != nil {
			return nil, err
		}
	}

	p7this.s6QueryBuilder.sqlString.WriteByte(';')

	return &S6Query{
		SQLString: p7this.s6QueryBuilder.sqlString.String(),
		S5Value:   p7this.s6QueryBuilder.s5Value,
	}, nil
}

func (p7this *S6InsertBuilder[T]) F8EXEC(ctx context.Context) (sql.Result, error) {
	p7s6Context := &S6QueryContext{
		QueryType: "INSERT",
		i9Builder: p7this,
		p7s6Model: p7this.s6QueryBuilder.p7s6Model,
		p7s6Query: nil,
	}
	p7s6Result := f8DoEXEC(ctx, p7this.i9Session, &p7this.s6Monitor, p7s6Context)
	return p7s6Result.I9SQLResult, p7s6Result.I9Err
}

// F8OnConflictBuilder 跳到中间 builder，处理 On Conflict 的内容
func (p7this *S6InsertBuilder[T]) F8OnConflictBuilder() *S6ConflictBuilder[T] {
	return &S6ConflictBuilder[T]{
		p7s6Insert: p7this,
	}
}
