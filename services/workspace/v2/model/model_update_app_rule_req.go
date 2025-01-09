package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppRuleReq 创建规则。
type UpdateAppRuleReq struct {

	// 规则名称： 1. 名称允许可见字符或空格，不可为全空格。 2. 长度1~64个字符。
	Name *string `json:"name,omitempty"`

	// 规则描述： 1. 名称允许可见字符或空格，不可为全空格。 2. 长度0~128个字符。
	Description *string `json:"description,omitempty"`

	Rule *Rule `json:"rule,omitempty"`
}

func (o UpdateAppRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppRuleReq struct{}"
	}

	return strings.Join([]string{"UpdateAppRuleReq", string(data)}, " ")
}
