package modules

import "database/sql"

// S6Result 框架的数据库结果对象：封装真正的数据库结果对象
type S6Result struct {
	I9SQLResult sql.Result
	I9Err       error
}
