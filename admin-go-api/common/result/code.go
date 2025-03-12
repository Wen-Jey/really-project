//状态语/状态码转换

package result

//Codes 定义状态

type Codes struct {
	SUCCESS                    uint
	FAILED                     uint
	Message                    map[uint]string
	NOAUTH                     uint
	AUTHFORMATERROR            uint
	MissingLoginParameter      uint
	VerificationCodeHasExpired uint
	CAPTCHANOTTRUE             uint
	PASSWORDNOTTRUE            uint
	STATUSISENABLE             uint
	POSTALREADYEXISTS          uint
}

// APICode 状态

var ApiCode = &Codes{
	SUCCESS:                    200,
	FAILED:                     501,
	NOAUTH:                     403,
	AUTHFORMATERROR:            405,
	MissingLoginParameter:      407,
	VerificationCodeHasExpired: 408,
	CAPTCHANOTTRUE:             409,
	PASSWORDNOTTRUE:            410,
	STATUSISENABLE:             411,
	POSTALREADYEXISTS:          412,
}

// 状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:                    "成功",
		ApiCode.FAILED:                     "失败",
		ApiCode.NOAUTH:                     "请求头中的auth为空",
		ApiCode.AUTHFORMATERROR:            "请求头中的auth格式有错误",
		ApiCode.MissingLoginParameter:      "缺少登录参数",
		ApiCode.VerificationCodeHasExpired: "验证码已失效",
		ApiCode.CAPTCHANOTTRUE:             "验证码不正确，请重新输入",
		ApiCode.PASSWORDNOTTRUE:            "密码不正确",
		ApiCode.STATUSISENABLE:             "您的账号已被停用,请联系管理员",
		ApiCode.POSTALREADYEXISTS:          "岗位名称或岗位编码已存在，请重新输入",
	}
}

//供外部调用

func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
