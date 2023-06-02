package modules

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
