package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateHttpIgnoreRuleRequestBody struct {

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则描述，最长512字符
	Description *string `json:"description,omitempty"`

	// 误报路径
	Url *string `json:"url,omitempty"`

	// 规则编号
	Rule string `json:"rule"`

	// 误报屏蔽模式，默认为0即旧模式
	Mode int32 `json:"mode"`

	// 域名列表
	Domains []string `json:"domains"`

	// 屏蔽规则url类型（前缀：prefix，等于：equal）
	UrlLogic *string `json:"url_logic,omitempty"`

	Advanced *HttpIgnoreRuleCondition `json:"advanced,omitempty"`

	// 命中条件
	Conditions []HttpIgnoreRuleCondition `json:"conditions"`
}

func (o CreateHttpIgnoreRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpIgnoreRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHttpIgnoreRuleRequestBody", string(data)}, " ")
}
