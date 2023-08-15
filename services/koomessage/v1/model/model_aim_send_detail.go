package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimSendDetail 短信发送明细。
type AimSendDetail struct {

	// 发送明细的唯一标识ID。
	MsgId *string `json:"msg_id,omitempty"`

	// 任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 创建解析任务时填写用户唯一标识，手机号码或者任何的唯一标识，唯一标识不超过64个字符。发送智能信息时则必须填客户的手机号码。此处为手机号。样例为：130****0001。
	CustFlag *string `json:"cust_flag,omitempty"`

	// 发送的目标手机号。
	SendAccount *string `json:"send_account,omitempty"`

	// 发送状态。 - success：发送成功  - fail：发送失败  - sending：发送中
	SendStatus *string `json:"send_status,omitempty"`

	// 发送时间。
	SendTime *string `json:"send_time,omitempty"`

	// 接收时间。
	ReceiveTime *string `json:"receive_time,omitempty"`

	// 发送错误码。
	ResultCode *string `json:"result_code,omitempty"`
}

func (o AimSendDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimSendDetail struct{}"
	}

	return strings.Join([]string{"AimSendDetail", string(data)}, " ")
}
