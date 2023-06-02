package metadata

import "reflect"

// orm 支持的结构体属性的 tag 上的 key 都放在这里
const (
	// 支持的 key 的数量
	tagNum           int    = 1
	tagKeyColumnName string = "column_name"
)

// 映射模型
// 处理结构体属性和数据库列名的互相转换
type S6Model struct {
	// 结构体对应的表名
	TableName string
	// 切片：index =>结构体属性
	S5P7S6ModelField []*S6ModelField
	// M3FieldToColumn map：结构体属性 => 数据库列名
	M3FieldToColumn map[string]*S6ModelField
	// M3ColumnToField map：数据库列名 => 结构体属性
	M3ColumnToField map[string]*S6ModelField
}

// F8S6ModelOption 方法抽象：针对 S6Model 的 Option 设计模式
type F8S6ModelOption func(p7s6om *S6Model) error

// 映射模型的每个属性
type S6ModelField struct {
	// 结构体属性名
	FieldName string
	//结构体属性反射类型
	I9Type reflect.Type
	// 结构体属性相对于对象的起始地址的偏移量
	Offset uintptr
	// 数据库列名
	ColumnName string
	// 反射时得到的下标
	Index int
}

// I9TableName 接口抽象：实现这个接口来返回自定义的表名
type I9TableName interface {
	// F8TableName 方法抽象：返回自定义的表名
	F8TableName() string
}
