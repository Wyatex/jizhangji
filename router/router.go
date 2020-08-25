package router

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"jizhangji/app/api/hello"
	"jizhangji/app/api/user"
	"jizhangji/app/middleware/cors"
)

var Token *gtoken.GfToken

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(cors.DefaultCORS)
		group.ALL("/", hello.Hello)

		// 用户api
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.ALL("/register", user.Register)
			group.ALL("/info", user.GetUserInfo)
		})
	})

	// 启动gtoken
	Token = &gtoken.GfToken{
		//Timeout:         10 * 1000,
		CacheMode:        g.Config().GetInt8("gtoken.cache-mode"),
		MultiLogin:       g.Config().GetBool("gtoken.multi-login"),
		LoginPath:        "/user/login",
		LoginBeforeFunc:  user.Login,
		LogoutPath:       "/user/logout",
		LogoutBeforeFunc: user.LogoutBefore,
		AuthPaths:        g.SliceStr{"/user"},
		AuthExcludePaths: g.SliceStr{"/user/register", "/user/login"},
		GlobalMiddleware: true,
	}
	user.Token = Token
	Token.Start()

}
