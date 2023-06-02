package modules

import (
	"context"
	"database/sql"
)

// I9Session 会话抽象：代表一次数据库操作
type I9Session interface {
	// f8GetS6Monitor 获取控制器
	f8GetS6Monitor() s6Monitor
	// f8DoQueryContext 用于执行 SELECT
	f8DoQueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	// f8DoExecContext 用于执行 INSERT、UPDATE、DELETE
	f8DoEXECContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}
