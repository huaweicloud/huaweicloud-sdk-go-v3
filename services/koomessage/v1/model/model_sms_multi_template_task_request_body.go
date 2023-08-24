package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsMultiTemplateTaskRequestBody struct {

	// 短信内容。
	SmsContent []MsgContent `json:"sms_content"`

	// 短信通道号。  > 模板所属签名的通道号，可以从“云消息服务KooMessage-管理控制台-短消息配置（国内）-短消息签名管理-通道号”中获取。 > 签名和模板为对应关系，模板所属签名可在“短消息模板管理”查看。未填写时默认取sms_content第一条数据模板所属签名的通道号。
	ChannelNum string `json:"channel_num"`

	// 扩展参数。  在状态报告中会原样返回。  不允许赋空值，不允许携带以下字符：“{”，“}”（即大括号）。
	Extend *string `json:"extend,omitempty"`

	// 发送任务名称。  > 不能为空白字符串，允许重复，为空时默认为Task_拼接当前时间值。
	TaskName *string `json:"task_name,omitempty"`
}

func (o SmsMultiTemplateTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsMultiTemplateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"SmsMultiTemplateTaskRequestBody", string(data)}, " ")
}
