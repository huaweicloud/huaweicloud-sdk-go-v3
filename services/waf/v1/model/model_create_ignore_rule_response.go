package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateIgnoreRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则的时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 屏蔽的内置规则id（内置规则id通常可以在Web应用防火墙控制台的防护策略->策略名称->Web基础防护->防护规则中查询，也可以从防护事件的事件详情中查询内置规则id）
	Rule *string `json:"rule,omitempty"`

	// 版本号固定值为1,代表v2版本误报屏蔽规则，v1版本仅支持兼容旧版本，不支持创建
	Mode *int32 `json:"mode,omitempty"`

	// 条件列表
	Conditions *[]Condition `json:"conditions,omitempty"`

	// 防护域名或防护网站
	Domain         *[]string `json:"domain,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateIgnoreRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIgnoreRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateIgnoreRuleResponse", string(data)}, " ")
}
