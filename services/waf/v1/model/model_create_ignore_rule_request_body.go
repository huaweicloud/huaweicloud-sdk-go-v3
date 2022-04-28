package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建误报屏蔽规则请求体
type CreateIgnoreRuleRequestBody struct {

	// 防护域名或防护网站，数组长度为0时，代表规则对全部域名或防护网站生效
	Domain []string `json:"domain"`

	// 条件列表
	Conditions []CreateCondition `json:"conditions"`

	// 固定值为1,代表v2版本误报屏蔽规则，v1版本仅用于兼容旧版本，不支持创建
	Mode int32 `json:"mode"`

	// 屏蔽的内置规则id（内置规则id通常可以在Web应用防火墙控制台的防护策略->策略名称->Web基础防护->防护规则中查询，也可以在防护事件的事件详情中看到查询规则id）
	Rule string `json:"rule"`

	// 屏蔽规则描述
	Description *string `json:"description,omitempty"`
}

func (o CreateIgnoreRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIgnoreRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIgnoreRuleRequestBody", string(data)}, " ")
}
