package cors

import "github.com/gogf/gf/net/ghttp"

// 允许接口跨域请求
func DefaultCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}