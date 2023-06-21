package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAimResolveDetailsRequest struct {

	// 解析任务ID或者发送任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 创建解析任务时填写用户唯一标识，手机号码或者任何的唯一标识，唯一标识不超过64个字符。发送智能信息时则必须填客户的手机号码。此处为手机号。样例为：130****0001。
	CustFlag *string `json:"cust_flag,omitempty"`

	// 签名。
	SmsSign *string `json:"sms_sign,omitempty"`

	// 智能信息短链，通过自己的短信渠道发送时，需要把该短链添加到短信模板中，并确保发送短信时的签名与创建短链时的签名保持一致。
	AimUrl *string `json:"aim_url,omitempty"`

	// 解析状态。 - success：解析成功  - fail：解析失败  - unresolved：未解析
	ResolvedStatus *string `json:"resolved_status,omitempty"`

	// 短链创建开始时间。格式为：2019-10-12T07:20:50Z。  > 需同时传入end_time才能生效，单独传begin_time不会作为过滤条件。缺省：查询最近二十四小时数据。
	BeginTime *string `json:"begin_time,omitempty"`

	// 短链创建结束时间。格式为：2019-10-12T07:20:50Z。  > 需同时传入begin_time才能生效，单独传end_time不会作为过滤条件。缺省：查询最近二十四小时数据。
	EndTime *string `json:"end_time,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。  >为提高查询效率，offset+limit须小于等于10000，超出范围查询为空。
	Offset int32 `json:"offset"`

	// 每页显示的条目数量。
	Limit int32 `json:"limit"`
}

func (o ListAimResolveDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimResolveDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListAimResolveDetailsRequest", string(data)}, " ")
}
