package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResolveRuleParam struct {

	// endpoint的ID，uuid形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// 规则名称。
	Name *string `json:"name,omitempty"`

	// 域名。
	DomainName *string `json:"domain_name,omitempty"`

	// 当前规则所属的endpoint_id。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 资源状态。 取值范围：PENDING_CREATE, ACTIVE, PENDING_DELETE, ERROR。
	Status *string `json:"status,omitempty"`

	// 规则类型。 预留字段，当前默认为FORWARD。
	RuleType *string `json:"rule_type,omitempty"`

	// 当前规则下的ip地址数量。
	IpaddressCount *int32 `json:"ipaddress_count,omitempty"`

	Routers *[]Router `json:"routers,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ResolveRuleParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResolveRuleParam struct{}"
	}

	return strings.Join([]string{"ResolveRuleParam", string(data)}, " ")
}
