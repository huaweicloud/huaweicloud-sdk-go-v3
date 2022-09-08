package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendVerificationCodeV2Req struct {

	// 发送验证码的类型： 2：发送邮件验证码
	ReceiverType int32 `json:"receiver_type"`

	// 发送验证码的超时时间。 如果不填的话，采用系统默认超时时间5分钟。 单位：分钟
	Timeout *int32 `json:"timeout,omitempty"`

	// 指定发送验证码的邮箱地址。
	Email string `json:"email"`

	// 根据该参数的取值选择发送邮件验证码的语言。 zh-cn：中文en-us：英文
	Lang *string `json:"lang,omitempty"`
}

func (o SendVerificationCodeV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendVerificationCodeV2Req struct{}"
	}

	return strings.Join([]string{"SendVerificationCodeV2Req", string(data)}, " ")
}
