package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻 自定义validator
	Type   uint   `form:"type" json:"type" binding:"required,oneof=1 2"`  //1注册 2表示动态验证码登录
}
