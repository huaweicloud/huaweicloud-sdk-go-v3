package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResolverRuleRequestBody struct {

	// 规则名称。 取值范围：1-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）。
	Name string `json:"name"`

	// 域名。
	DomainName string `json:"domain_name"`

	// 当前规则所属的终端节点ID。
	EndpointId string `json:"endpoint_id"`

	// 规则的目标IP地址。
	Ipaddresses []IpInfo `json:"ipaddresses"`
}

func (o CreateResolverRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResolverRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResolverRuleRequestBody", string(data)}, " ")
}
