package user

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"jizhangji/app/model/ledger"
	"jizhangji/app/model/user"
	"jizhangji/lib/structure"
)

const (
	SUCCEED		=	0
	FAIL		=	-1
	ERROR		= 	-2
)

// 注册输入参数
type RegisterInput struct {
	Passport string `v:"required|length:4,16#账号不能为空|账号长度为:min到:max位"`
	Password string `v:"required|length:6,16#请输入密码|密码长度不够"`
	Nickname string
	Email    string `v:"email#请输入正确的邮箱"`
	Phone	 string `v:"phone#请输入正确的手机号"`
}

// 用户注册处理
func Regiseter(data *RegisterInput) (error, int) {
	// 先执行参数检查
	if err := gvalid.CheckStruct(data, nil); err != nil {
		return errors.New(err.FirstString()), ERROR
	}

	// 昵称非必选项，注册时为空则使用账号名称
	if data.Nickname == "" {
		data.Nickname = data. Passport
	}

	// 账号唯一性检查
	if !CheckPassport(data.Passport) {
		return errors.New(fmt.Sprintf("账号 %s 已被注册", data.Passport)), FAIL
	}

	// 邮箱唯一性检查
	if data.Email != "" {
		if !CheckEmail(data.Email) {
			return errors.New(fmt.Sprintf("邮箱 %s 已被注册", data.Email)), FAIL
		}
	}

	// 手机号唯一性检查
	if data.Phone != "" {
		if !CheckPhone(data.Phone) {
			return errors.New(fmt.Sprintf("手机 %s 已被注册", data.Phone)), FAIL
		}
	}

	// 改成前端加密
	//// 对密码进行加密
	//data.Password, _ = gmd5.Encrypt(data.Password)

	// 将输入参数赋值到数据库实体对象上
	var entity user.Entity
	if err := gconv.Struct(data, &entity); err != nil {
		return err, ERROR
	}

	// 一些其他的操作

	// 写进数据库
	if _, err := user.Save(entity); err != nil {
		return err, ERROR
	}

	return nil, SUCCEED
}

// 检查账号passport是否唯一,存在返回false,否则true
func CheckPassport(passport string) bool {
	if i, err := user.FindCount("passport", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查email是否唯一,存在返回false,否则true
func CheckEmail(email string) bool {
	if i, err := user.FindCount("email", email); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查phone是否唯一,存在返回false,否则true
func CheckPhone(phone string) bool {
	if i, err := user.FindCount("phone", phone); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 根据账号获取实体
func GetByPassport(p string) (*user.Entity, error) {
	if p == "" {
		return new(user.Entity), errors.New("没找到这个用户")
	}

	return user.Model.FindOne("passport = ?", p)
}

// 根据用户id获取实体
func GetUserById(uid int64)  (*user.Entity, error) {
	if uid <= 0 {
		glog.Error(" get id error")
		return new(user.Entity), errors.New("参数不合法")
	}

	return user.Model.FindOne("uid = ?", uid)
}

// 根据用户id获取账本信息
func GetLedgerById(uid int64) ([]structure.Ledger, int, error) {
	if uid <= 0 {
		glog.Error(" get id error")
		return []structure.Ledger{}, 0, errors.New("参数不合法")
	}
	if num, err := ledger.Model.FindCount("uid = ?", uid); err != nil {
		return []structure.Ledger{}, num, err
	} else {
		if result, err := ledger.Model.FindAll("uid = ?", uid); err != nil {
			return []structure.Ledger{}, num, err
		} else {
			ledgerList := make([]structure.Ledger, num)
			for n := 0; n < num; n++ {
				ledgerList[n].Id = result[n].Id
				ledgerList[n].LedgerName = result[n].Name
			}
			return ledgerList, num, nil
		}
	}

}