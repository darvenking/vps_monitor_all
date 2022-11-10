package param

type UserLoginParam struct {
	UserName string `v:"required#请输入用户姓名"`
	PassWord string `v:"required#请输入密码"`
}
