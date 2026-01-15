package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListResolverRuleParam struct {

	// 转发规则的ID。
	Id *string `json:"id,omitempty"`

	// 规则名称。
	Name *string `json:"name,omitempty"`

	// 域名。
	DomainName *string `json:"domain_name,omitempty"`

	// 当前规则所属的终端节点ID。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释：** 资源状态。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_DELETE：删除中 - ERROR：失败
	Status *string `json:"status,omitempty"`

	// 规则类型。 预留字段，当前默认为FORWARD。
	RuleType *string `json:"rule_type,omitempty"`

	// 当前规则下的IP地址数量。
	IpaddressCount *int32 `json:"ipaddress_count,omitempty"`

	// 关联的VPC信息。
	Routers *[]Router `json:"routers,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ListResolverRuleParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResolverRuleParam struct{}"
	}

	return strings.Join([]string{"ListResolverRuleParam", string(data)}, " ")
}
