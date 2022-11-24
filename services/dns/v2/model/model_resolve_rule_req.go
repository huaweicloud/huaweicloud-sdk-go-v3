package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResolveRuleReq struct {

	// 规则名称。 取值范围：1-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）。
	Name string `json:"name"`

	// 域名。
	DomainName string `json:"domain_name"`

	// 当前规则所属的endpoint_id。
	EndpointId string `json:"endpoint_id"`

	// 当前规则所在的region。
	Region string `json:"region"`

	// 规则关联的目标ip地址。
	Ipaddresses []IpInfo `json:"ipaddresses"`

	// 规则关联的目标ip地址。
	Routers []Router `json:"routers"`
}

func (o ResolveRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResolveRuleReq struct{}"
	}

	return strings.Join([]string{"ResolveRuleReq", string(data)}, " ")
}
