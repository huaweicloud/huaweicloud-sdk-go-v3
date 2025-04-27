package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateResolverRuleRequestBody struct {

	// 规则名称。 取值范围：1-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）。
	Name *string `json:"name,omitempty"`

	// 规则的目标IP地址。
	Ipaddresses *[]IpInfo `json:"ipaddresses,omitempty"`
}

func (o UpdateResolverRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResolverRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResolverRuleRequestBody", string(data)}, " ")
}
