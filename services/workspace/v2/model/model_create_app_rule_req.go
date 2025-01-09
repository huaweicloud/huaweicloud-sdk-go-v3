package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppRuleReq 创建规则。
type CreateAppRuleReq struct {

	// 规则名称： 1. 名称允许可见字符或空格，不可为全空格。 2. 长度1~64个字符。
	Name string `json:"name"`

	// 规则描述： 1. 名称允许可见字符或空格，不可为全空格。 2. 长度0~128个字符。
	Description *string `json:"description,omitempty"`

	Rule *Rule `json:"rule"`
}

func (o CreateAppRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppRuleReq struct{}"
	}

	return strings.Join([]string{"CreateAppRuleReq", string(data)}, " ")
}
