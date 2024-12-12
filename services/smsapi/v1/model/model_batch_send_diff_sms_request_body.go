package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSendDiffSmsRequestBody 批量发送短信的请求body数据结构
type BatchSendDiffSmsRequestBody struct {

	// 短信发送方的号码
	From *string `json:"from,omitempty"`

	// SP的回调地址，用于单条接收短信状态报告
	StatusCallback *string `json:"statusCallback,omitempty"`

	// 短信内容
	SmsContent *[]SmsContent `json:"smsContent,omitempty"`

	// 扩展参数
	Extend *string `json:"extend,omitempty"`
}

func (o BatchSendDiffSmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSendDiffSmsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSendDiffSmsRequestBody", string(data)}, " ")
}
