# API接口功能

返回数据统一格式：
```json
{
    "code": code,
    "msg":  msg,
    "data": data
}
```

## 用户功能
路由地址： `/user`

### 注册
路由地址： `/user/register`

需要发送的表单：

信息 | 属性名 | 要求
- | - | -
账号| passport | 长度位于4~16位之间，所有账号唯一，必填
密码 | password | 长度位于6~16位之间，支持弱密码，必填
昵称 | nickname | 长度小于16，选填，不发送则和账号相同
邮箱 | email | 用于密码找回，选填，如果填了会验证唯一性
手机 | phone | 用于密码找回，选填，如果填了会验证唯一性

返回数据说明：返回code为0表示成功，返回-1表示提交信息有问题，返回-2表示服务器有问题

### 登录
路由地址： `/user/login`

需要发送的表单：

信息 | 属性名 | 要求
- | - | -
账号| passport | 注册时的账号，必填
密码 | password | 注册时的密码，必填

返回返回数据说明：code：0表示成功，会在data里面带有"token":用户的token。

### 登出
路由地址： `/user/logout`

直接访问路由即可，http头需要带上 `Authorization: Bearer token`

### 用户信息
路由地址： `/user/info`

直接访问路由即可，http头需要带上 `Authorization: Bearer token`