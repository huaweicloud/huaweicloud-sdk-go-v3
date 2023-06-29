package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmsChannel 短信发送通道参数。
type SmsChannel struct {

	// 短信通道号。  > 必须与另外三个字段sms_tpl_id、sms_sign、sms_app_name相匹配，这些字段信息可以从“云消息服务KooMessage-管理控制台-短信配置-短信签名管理-通道号”中获取。
	ChannelNumber string `json:"channel_number"`

	// 短信模板ID。  > 必须与另外三个字段channel_number、sms_sign、sms_app_name相匹配，这些字段信息可以从“云消息服务KooMessage-管理控制台-短信配置-短信模板管理-模板ID”中获取。
	SmsTplId string `json:"sms_tpl_id"`

	// 短信签名。  > 必须与另外三个字段channel_number、sms_tpl_id、sms_app_name相匹配，这些字段信息可以从“云消息服务KooMessage-管理控制台-短信配置-短信模板管理-所属签名”中获取。
	SmsSign string `json:"sms_sign"`

	// 短信应用名称。  > 必须与另外三个字段channel_number、sms_sign、sms_tpl_id相匹配，这些字段信息可以从“云消息服务KooMessage-管理控制台-短信配置-短信模板管理-所属应用”中获取。
	SmsAppName string `json:"sms_app_name"`
}

func (o SmsChannel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsChannel struct{}"
	}

	return strings.Join([]string{"SmsChannel", string(data)}, " ")
}
