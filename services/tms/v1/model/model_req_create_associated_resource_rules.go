package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqCreateAssociatedResourceRules 启用规则的请求参数
type ReqCreateAssociatedResourceRules struct {

	// 批量启用的规则信息
	Rules *[]ReqAssociatedResourceRule `json:"rules,omitempty"`
}

func (o ReqCreateAssociatedResourceRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqCreateAssociatedResourceRules struct{}"
	}

	return strings.Join([]string{"ReqCreateAssociatedResourceRules", string(data)}, " ")
}
