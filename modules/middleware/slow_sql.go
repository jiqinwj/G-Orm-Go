package middleware

import (
	"G-Orm-go/modules"
	"context"
	"log"
	"time"
)

// SlowLogMiddlewareBuild 计算查询执行时间，用于捕获慢 SQL
func SlowLogMiddlewareBuild() modules.F8Middleware {
	return func(next modules.F8MiddlewareHandle) modules.F8MiddlewareHandle {
		return func(ctx context.Context, p7s6Context *modules.S6QueryContext) *modules.S6QueryResult {
			timeStart := time.Now()
			t4 := next(ctx, p7s6Context)
			timeEnd := time.Now()
			timeCost := timeEnd.Sub(timeStart).Milliseconds()
			log.Printf("time pass %d ms\r\n", timeCost)
			if 200 < timeCost {
				log.Printf("slow sql, time pass %d ms\r\n", timeCost)
			}
			return t4
		}
	}
}
