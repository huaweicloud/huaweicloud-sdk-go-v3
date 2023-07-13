package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VerifyCodeSendDtov1 struct {

	// 用户身份信息（手机号码或邮箱帐号或用户真实帐号）。 > 必须和发送滑块验证码时带的用户身份信息相同。
	User string `json:"user"`

	// 验证码发送方式。 user类型是用户真实帐号时必选；如果没有选择的话，优先短信方式。 * sms：短信方式 * email：邮件方式
	SendMethod *string `json:"sendMethod,omitempty"`

	// 访问Token字符串。通过[[校验滑块验证码](https://support.huaweicloud.com/api-meeting/meeting_21_0101.html)](tag:hws)[[校验滑块验证码](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0101.html)](tag:hk) 接口获取。
	Token *string `json:"token,omitempty"`
}

func (o VerifyCodeSendDtov1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyCodeSendDtov1 struct{}"
	}

	return strings.Join([]string{"VerifyCodeSendDtov1", string(data)}, " ")
}
