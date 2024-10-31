package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpRuleActionResponse 防护规则返回页面
type HttpRuleActionResponse struct {

	// 内容类型
	ContentType *string `json:"content_type,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`
}

func (o HttpRuleActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpRuleActionResponse struct{}"
	}

	return strings.Join([]string{"HttpRuleActionResponse", string(data)}, " ")
}
