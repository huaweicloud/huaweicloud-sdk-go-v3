package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowResolverRuleRespParam struct {

	// 转发规则的ID，UUID形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// 规则名称。
	Name *string `json:"name,omitempty"`

	// 域名。
	DomainName *string `json:"domain_name,omitempty"`

	// 当前规则所属的终端节点ID。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 资源状态。 取值范围：PENDING_CREATE, ACTIVE, PENDING_DELETE, ERROR。
	Status *string `json:"status,omitempty"`

	// 规则类型。 预留字段，当前默认为FORWARD。
	RuleType *string `json:"rule_type,omitempty"`

	// 当前规则下的IP地址数量。
	IpaddressCount *int32 `json:"ipaddress_count,omitempty"`

	// 规则的目标IP地址。
	Ipaddresses *[]IpValue `json:"ipaddresses,omitempty"`

	// 关联的VPC信息。
	Routers *[]Router `json:"routers,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ShowResolverRuleRespParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResolverRuleRespParam struct{}"
	}

	return strings.Join([]string{"ShowResolverRuleRespParam", string(data)}, " ")
}
