package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpRuleActionDetail 防护规则动作
type HttpRuleActionDetail struct {

	// 返回页面重定向的url
	RedirectUrl *string `json:"redirect_url,omitempty"`

	Response *HttpRuleActionResponse `json:"response,omitempty"`
}

func (o HttpRuleActionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpRuleActionDetail struct{}"
	}

	return strings.Join([]string{"HttpRuleActionDetail", string(data)}, " ")
}
