package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimResolveDetail 解析详情。
type AimResolveDetail struct {

	// 解析详情唯一标识ID。
	ResolveId *string `json:"resolve_id,omitempty"`

	// 任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 发送的用户名。
	SendAccount *string `json:"send_account,omitempty"`

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 创建解析任务时填写用户唯一标识，手机号码或者任何的唯一标识，唯一标识不超过64个字符。发送智能信息时则必须填客户的手机号码。此处为手机号。样例为：130****0001。
	CustFlag *string `json:"cust_flag,omitempty"`

	// 智能信息短链，通过自己的短信渠道发送时，需要把该短链添加到短信模板中，并确保发送短信时的签名与创建短链时的签名保持一致。
	AimUrl *string `json:"aim_url,omitempty"`

	// 短链申请结果返回码。 - 0：成功 - 非0：失败，具体请参见错误码。
	ResultCode *string `json:"result_code,omitempty"`

	// 短链生成时间。样例为：2019-10-12T07:20:50Z。
	GenerateDate *string `json:"generate_date,omitempty"`

	// 短链到期时间。样例为：2019-10-12T07:20:50Z。
	ExpireDate *string `json:"expire_date,omitempty"`

	// 解析时间。样例为：2019-10-12T07:20:50Z。
	ResolvedTime *string `json:"resolved_time,omitempty"`

	// 解析状态。 - success：解析成功 - fail：解析失败 - unresolved：未解析
	ResolvedStatus *string `json:"resolved_status,omitempty"`
}

func (o AimResolveDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimResolveDetail struct{}"
	}

	return strings.Join([]string{"AimResolveDetail", string(data)}, " ")
}
