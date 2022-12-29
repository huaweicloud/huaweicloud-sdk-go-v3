package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointResp struct {

	// endpoint的ID，uuid形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// endpoint的名称。
	Name *string `json:"name,omitempty"`

	// 终端节点方向。 取值： inbound，表示入站规则。 outbound，表示出站规则。
	Direction *string `json:"direction,omitempty"`

	// 资源状态。 取值范围：PENDING_CREATE, ACTIVE, PENDING_DELETE, ERROR。
	Status *string `json:"status,omitempty"`

	// endpoint所属的vpc_id。
	VpcId *string `json:"vpc_id,omitempty"`

	// 该endpoint下的ip地址数量。
	IpaddressCount *int32 `json:"ipaddress_count,omitempty"`

	// 该endpoint下的解析规则数量。 仅创建出站终端节点时返回该参数。
	ResolverRuleCount *int32 `json:"resolver_rule_count,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o EndpointResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointResp struct{}"
	}

	return strings.Join([]string{"EndpointResp", string(data)}, " ")
}
