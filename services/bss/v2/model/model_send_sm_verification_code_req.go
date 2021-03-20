package model

import (
	"encoding/json"

	"strings"
)

type SendSmVerificationCodeReq struct {
	// 接受短信验证码的手机号码。

	MobilePhone string `json:"mobile_phone"`
	// 超时时间，不指定时默认为10分钟。 单位：分钟

	Timeout *int32 `json:"timeout,omitempty"`
	// 发送的短信的语言。 zh-cn: 中文en-us: 英语 不设置默认为偏好设置的默认语言。

	Language *string `json:"language,omitempty"`
	// 短信发送模板中的变量，具体参见表1。

	SmTemplateArgs *[]TemplateArgs `json:"sm_template_args,omitempty"`
}

func (o SendSmVerificationCodeReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SendSmVerificationCodeReq struct{}"
	}

	return strings.Join([]string{"SendSmVerificationCodeReq", string(data)}, " ")
}
