package result

import (
	"G-Orm-go/modules/internal"
	"G-Orm-go/modules/metadata"
	"database/sql"
	"reflect"
	"unsafe"
)

// 确保 F8NewS6ResultUseUnsafe 实现的是 F8NewI9Result
var _ F8NewI9Result = F8NewS6ResultUseUnsafe

// s6ResultUseUnsafe 用 unsafe 实现 I9Result
type s6ResultUseUnsafe struct {
	// p7pointer 存储数据库返回的查询结果的结构体的起始地址
	p7pointer unsafe.Pointer
	// p7s6Model 映射模型z
	p7s6Model *metadata.S6Model
}

// F8NewS6ResultUseUnsafe 构造 s6ResultUseUnsafe
func F8NewS6ResultUseUnsafe(value any, p7s5OrmModel *metadata.S6Model) I9Result {
	return &s6ResultUseUnsafe{
		p7pointer: unsafe.Pointer(reflect.ValueOf(value).Pointer()),
		p7s6Model: p7s5OrmModel,
	}
}

func (p7this s6ResultUseUnsafe) F8SetField(rows *sql.Rows) error {
	// 返回数据库列名
	s5ColumnName, err := rows.Columns()
	if nil != err {
		return err
	}
	if len(s5ColumnName) > len(p7this.p7s6Model.M3ColumnToField) {
		return internal.ErrTooManyReturnedColumns
	}

	// 这里初始化的时候需要长度
	s5ColumnValue := make([]any, len(s5ColumnName))
	for i, t4ColumnName := range s5ColumnName {
		// 通过数据库列名找到对应的结构体属性
		p7s6ModelField, ok := p7this.p7s6Model.M3ColumnToField[t4ColumnName]
		if !ok {
			return internal.F8NewErrUnknownColumn(t4ColumnName)
		}
		// 通过结构体属性的内存偏移量，找到结构体属性的位置
		t4p7pointer := unsafe.Pointer(uintptr(p7this.p7pointer) + p7s6ModelField.Offset)
		// 在找到的找到结构体属性的位置上构造结构体属性
		t4value := reflect.NewAt(p7s6ModelField.I9Type, t4p7pointer)
		s5ColumnValue[i] = t4value.Interface()
	}
	// 从数据库返回的查询结果里取数据
	if err = rows.Scan(s5ColumnValue...); err != nil {
		return err
	}
	return nil
}

func (p7this s6ResultUseUnsafe) F8GetField(name string) (any, error) {
	fd, ok := p7this.p7s6Model.M3FieldToColumn[name]
	if !ok {
		return nil, internal.F8NewErrUnknownColumn(name)
	}
	ptr := unsafe.Pointer(uintptr(p7this.p7pointer) + fd.Offset)
	val := reflect.NewAt(fd.I9Type, ptr).Elem()
	return val.Interface(), nil
}
