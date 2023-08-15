package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimSendDetailsRequest Request Object
type ListAimSendDetailsRequest struct {

	// 任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 签名。
	SmsSign *string `json:"sms_sign,omitempty"`

	// 创建解析任务时填写用户唯一标识。  > 手机号码或者任何的唯一标识，唯一标识不超过64个字符。发送智能信息时则必须填客户的手机号码。此处为手机号。样例为：130****0001。
	CustFlag *string `json:"cust_flag,omitempty"`

	//  发送开始时间。格式为：2019-10-12T07:20:50Z。  > 需同时传入end_time才能生效，单独传begin_time不会作为过滤条件。
	BeginTime *string `json:"begin_time,omitempty"`

	// 发送结束时间。格式为：2019-10-12T07:20:50Z。  > 需同时传入begin_time才能生效，单独传end_time不会作为过滤条件。
	EndTime *string `json:"end_time,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。  >为提高查询效率，offset+limit须小于等于10000，超出范围查询为空。
	Offset int32 `json:"offset"`

	// 每页显示的条目数量。
	Limit int32 `json:"limit"`
}

func (o ListAimSendDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimSendDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListAimSendDetailsRequest", string(data)}, " ")
}
