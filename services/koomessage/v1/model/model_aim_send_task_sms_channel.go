package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询智能信息发送任务响应体。
type AimSendTaskSmsChannel struct {

	// 短信通道号。  > 预留字段，暂时为空。
	ChannelNumber *string `json:"channel_number,omitempty"`

	// 短信模板ID。
	SmsTplId *string `json:"sms_tpl_id,omitempty"`

	// 短信签名。
	SmsSign *string `json:"sms_sign,omitempty"`

	// 短信应用名称。  > 预留字段，暂时为空。
	SmsAppName *string `json:"sms_app_name,omitempty"`
}

func (o AimSendTaskSmsChannel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimSendTaskSmsChannel struct{}"
	}

	return strings.Join([]string{"AimSendTaskSmsChannel", string(data)}, " ")
}
