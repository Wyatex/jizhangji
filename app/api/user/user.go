package user

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"jizhangji/app/service/user"
	"jizhangji/lib/respond"
	"jizhangji/lib/structure"
)

var Token *gtoken.GfToken

// 注册请求参数，用于前后端交互参数格式约定
type RegisterRequest struct {
	user.RegisterInput
}

func Register(r *ghttp.Request) {
	var data *RegisterRequest

	// 使用GetStruct获取对象
	if err := r.GetStruct(&data); err != nil {
		response.Error(r, err.Error())
	}

	// 交给service层统一处理
	if err, code := user.Regiseter(&data.RegisterInput); err != nil {
		if code == -1 {
			response.Fail(r, err.Error())
		} else {
			response.Error(r, err.Error())
		}
	} else {
		response.Succeed(r, "注册成功")
	}
}

func Login(r *ghttp.Request) (string, interface{}) {
	passport := r.GetString("passport")
	password := r.GetString("password")

	if password == "" || passport == "" {
		response.Fail(r, "请输入正确的账号或者密码")
	}

	model, err := user.GetByPassport(passport)

	if err != nil {
		response.Fail(r, err.Error())
	}

	if model == nil || model.Uid <= 0 {
		response.Fail(r, "用户名或密码错误.")
	}

	ledgerList, num, err := user.GetLedgerById(gconv.Int64(model.Uid))
	if err != nil {
		response.Fail(r, err.Error())
	}

	sessionUser := structure.SessionUser {
		Uid:		model.Uid,
		Passport: 	model.Passport,
		Nickname:  	model.Nickname,
		LedgerNum: 	num,
		Ledger: 	ledgerList,
	}

	return model.Passport, sessionUser
}

func LogoutBefore(r *ghttp.Request) bool {
	userId := GetUserTokenData(r).Uid
	model, err := user.GetUserById(gconv.Int64(userId))
	if err != nil {
		glog.Warning("logout getUser error", err)
		return false
	} else if model.Uid != userId {
		// 登出用户不存在
		glog.Warning("logout userId error", userId)
		return false
	}
	return true
}

// 获取用户session信息
func GetUserTokenData(r *ghttp.Request) structure.SessionUser {
	resp := Token.GetTokenData(r)
	if !resp.Success() {
		return structure.SessionUser{}
	}

	var sessionUser structure.SessionUser
	err := gjson.DecodeTo(resp.GetString("data"), &sessionUser)
	if err != nil {
		glog.Error("get session user error", err)
	}

	return sessionUser
}

// 返回用户信息
func GetUserInfo(r *ghttp.Request) {
	data := GetUserTokenData(r)
	null := structure.SessionUser{}
	if data == null {
		response.Fail(r, "获取信息失败")
	} else {
		response.JsonExit(r, 0, "获取信息成功", data)
	}
}

func getNum() int {
	return 2
}