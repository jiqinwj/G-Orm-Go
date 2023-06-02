package metadata

import (
	"G-Orm-go/modules/internal"
	"reflect"
	"strings"
	"sync"
	"unicode"
)

// 元数据注册中心抽象
type I9Registry interface {
	// 查找元数据
	F8Get(p7s6Strut any) (*S6Model, error)
	// 注册元数据
	F8Register(p7s6Model any, s5f8Option ...F8S6ModelOption) (*S6Model, error)
}

// F8NewI9Registry 构造 I9Registry
// S6Registry 是 I9Registry 的实例
func F8NewI9Registry() I9Registry {
	return &S6Registry{}
}

func (p7this *S6Registry) F8Get(p7s6Struct any) (*S6Model, error) {
	// 看看传进来的结构体解析过没有
	i9type := reflect.TypeOf(p7s6Struct)
	value, ok := p7this.m3Model.Load(i9type.String())
	// 如果结构体已经解析过了，就直接返回
	if ok {
		return value.(*S6Model), nil
	}
	// 否则需要解析并注册映射模型
	return p7this.F8Register(p7s6Struct)
}

func (p7this *S6Registry) F8Register(p7s6Struct any, s5f8Option ...F8S6ModelOption) (*S6Model, error) {
	p7s6om, err := p7this.f8ParseModel(p7s6Struct)
	if nil != err {
		return nil, err
	}
	// 执行 Option 设计模式
	for _, t4f8 := range s5f8Option {
		err = t4f8(p7s6om)
		if nil != err {
			return nil, err
		}
	}
	// 注册映射模型
	i9type := reflect.TypeOf(p7s6Struct)
	p7this.m3Model.Store(i9type.String(), p7s6om)
	return p7s6om, nil
}

// S6Registry 元数据注册中心：实现 I9Registry 接口
type S6Registry struct {
	// m3Model map：解析好的映射模型
	// 这里可以预见会出现并发操作，所以用 sync.Map
	m3Model sync.Map
}

// f8ParseModel 解析结构体获取元数据
func (p7this *S6Registry) f8ParseModel(p7s6Model any) (*S6Model, error) {
	i9ModelType := reflect.TypeOf(p7s6Model)
	// 只接受一级结构体指针
	if reflect.Ptr != i9ModelType.Kind() || reflect.Struct != i9ModelType.Elem().Kind() {
		return nil, internal.F8NewErrInputOnlyStructPointer()
	}
	i9ModelType = i9ModelType.Elem()

	// 获取表名
	var tableName string
	t4i9TableName, ok := p7s6Model.(I9TableName)
	if ok {
		tableName = t4i9TableName.F8TableName()
	}
	if "" == tableName {
		tableName = f8CamelCaseToSnakeCase(i9ModelType.Name())
	}

	// 获取结构体属性数量
	fieldNum := i9ModelType.NumField()
	m3SToF := make(map[string]*S6ModelField, fieldNum)
	m3FToS := make(map[string]*S6ModelField, fieldNum)
	s5p7s6Field := make([]*S6ModelField, 0, fieldNum)
	// 解析结构体的每个字段
	for i := 0; i < fieldNum; i++ {
		s6FieldType := i9ModelType.Field(i)
		m3tag, err := p7this.f8ParseTag(s6FieldType.Tag)
		if nil != err {
			return nil, err
		}
		// 从标签里获取设置的数据库列名
		fieldName := m3tag[tagKeyColumnName]
		// 如果没有设置数据库列名，默认用转换成小驼峰的结构体属性名
		if "" == fieldName {
			fieldName = f8CamelCaseToSnakeCase(s6FieldType.Name)
		}
		// 正反方向都要存一份
		p7s6mf := &S6ModelField{
			FieldName:  s6FieldType.Name,
			I9Type:     s6FieldType.Type,
			Offset:     s6FieldType.Offset,
			ColumnName: fieldName,
			Index:      i,
		}
		s5p7s6Field = append(s5p7s6Field, p7s6mf)
		m3SToF[s6FieldType.Name] = p7s6mf
		m3FToS[fieldName] = p7s6mf
	}

	p7s6om := &S6Model{
		TableName:        tableName,
		M3FieldToColumn:  m3SToF,
		M3ColumnToField:  m3FToS,
		S5P7S6ModelField: s5p7s6Field,
	}

	return p7s6om, nil
}

// f8ToUnderscore 驼峰转蛇形
func f8CamelCaseToSnakeCase(oldString string) string {
	var s5NewString []byte
	for i, char := range oldString {
		// 如果是大写字母，前面加一个下划线，然后转小写字母
		if unicode.IsUpper(char) {
			// 如果首字母是大写字母，不用加下划线
			if 0 != i {
				s5NewString = append(s5NewString, '_')
			}
			s5NewString = append(s5NewString, byte(unicode.ToLower(char)))
		} else {
			s5NewString = append(s5NewString, byte(char))
		}
	}
	return string(s5NewString)
}

// f8ParseTag 解析结构体属性的标签
// 标签格式：`orm:"key1=value1,key2=value2"`
func (p7this *S6Registry) f8ParseTag(s6tag reflect.StructTag) (map[string]string, error) {
	// 从 tag 里面拿 orm 标签
	orm := s6tag.Get("orm")
	if "" == orm {
		return map[string]string{}, nil
	}
	// 解析 tag，其实就是解析字符串
	s5kv := strings.Split(orm, ",")
	m3tag := make(map[string]string, tagNum)
	for _, kv := range s5kv {
		t4s5kv := strings.Split(kv, "=")
		// 判断标签格式正不正确
		if 2 != len(t4s5kv) {
			return nil, internal.F8NewErrInvalidTagContent(kv)
		}
		m3tag[t4s5kv[0]] = t4s5kv[1]
	}
	return m3tag, nil
}
