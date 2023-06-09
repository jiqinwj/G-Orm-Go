package middleware

import (
	"G-Orm-go/modules"
	"context"
	"time"
)

// SlowLogTriggerMiddlewareBuild 触发慢 SQL 用的
func SlowLogTriggerMiddlewareBuild() modules.F8Middleware {
	return func(next modules.F8MiddlewareHandle) modules.F8MiddlewareHandle {
		return func(ctx context.Context, p7s6Context *modules.S6QueryContext) *modules.S6QueryResult {
			time.Sleep(500 * time.Millisecond)
			return next(ctx, p7s6Context)
		}
	}
}
