package modules

// S6Conflict 对应 INSERT 语句中处理冲突的部分
// 即 INSERT Statement 里的 ON CONFLICT
type S6Conflict struct {
	//用于判断冲突的列
	// S5ConflictColumn 用于判断冲突的列
	S5ConflictColumn []S6Column
	// S5Assignment 赋值语句
	S5Assignment []i9Assignment
}

// S6ConflictBuilder ON Conflict 查询构造器
// 设计这玩意，主要是为了把常规的 INSERT 语句和带 ON CONFLICT 的 INSERT 语句区分开
type S6ConflictBuilder[T any] struct {
	// p7s6Insert INSERT 查询构造器
	p7s6Insert *S6InsertBuilder[T]
	// S5ConflictColumn 用于判断冲突的列
	S5ConflictColumn []S6Column
}

// F8SetConflictColumn 设置用于判断冲突的列
func (p7this *S6ConflictBuilder[T]) F8SetConflictColumn(s5column ...S6Column) *S6ConflictBuilder[T] {
	p7this.S5ConflictColumn = s5column
	return p7this
}

// F8SetUpdate 设置赋值语句
func (p7this *S6ConflictBuilder[T]) F8SetUpdate(s5i9assignment ...i9Assignment) *S6InsertBuilder[T] {
	p7this.p7s6Insert.p7s6Conflict = &S6Conflict{
		S5ConflictColumn: p7this.S5ConflictColumn,
		S5Assignment:     s5i9assignment,
	}
	// 设置完赋值语句后，理论上关于冲突的部分就设置完了，应该返回 INSERT 构造器
	return p7this.p7s6Insert
}
