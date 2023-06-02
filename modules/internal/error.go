package internal

import (
	"errors"
	"fmt"
)

// #### 元数据 ####

func F8NewErrInputOnlyStructPointer() error {
	return errors.New("ORM: 只支持一级结构体指针作为输入\r\n")
}

func F8NewErrInvalidTagContent(tag string) error {
	return fmt.Errorf("ORM: 标签 [%s] 格式错误\r\n", tag)
}

func F8NewErrUnknownField(name string) error {
	return fmt.Errorf("ORM: 元数据中不存在 %v 属性。", name)
}

func F8NewErrUnknownColumn(column string) error {
	return fmt.Errorf("ORM: 元数据中不存在 %s 列", column)
}

// #### 结果集 ####

var ErrSqlRowsIsEmpty = errors.New("ORM: SELECT 查询结果为空")
var ErrTooManyReturnedColumns = errors.New("ORM: 返回的列过多")

// #### ORM ####

var ErrUpdateWithoutColumn = errors.New("ORM: UPDATE 没有设置更新的列")
var ErrUpdateWithoutWhere = errors.New("ORM: UPDATE 没有 WHERE")

var ErrDeleteWithoutWhere = errors.New("ORM: DELETE 没有 WHERE")

func NewErrUnsupportedExpressionType(e any) error {
	return fmt.Errorf("ORM: 不支持的表达式 %v", e)
}
