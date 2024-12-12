package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsContent struct {

	// 群发短信接收方的号码
	To *[]string `json:"to,omitempty"`

	// 短信模板ID
	TemplateId *string `json:"templateId,omitempty"`

	// 短信模板的变量值列表
	TemplateParas *[]string `json:"templateParas,omitempty"`

	// 短信签名
	Signature *string `json:"signature,omitempty"`
}

func (o SmsContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsContent struct{}"
	}

	return strings.Join([]string{"SmsContent", string(data)}, " ")
}
