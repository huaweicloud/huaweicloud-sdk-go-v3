package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointResp struct {

	// 终端节点的ID，UUID形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// 终端节点的名称。
	Name *string `json:"name,omitempty"`

	// 终端节点方向。 取值： inbound，表示入站终端节点。 outbound，表示出站终端节点。
	Direction *string `json:"direction,omitempty"`

	// **参数解释：** 资源状态。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_DELETE：删除中 - ERROR：失败
	Status *string `json:"status,omitempty"`

	// 终端节点所属的VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 该终端节点下的IP地址数量。
	IpaddressCount *int32 `json:"ipaddress_count,omitempty"`

	// 该终端节点下的转发规则数量。 仅创建出站终端节点时返回该参数。
	ResolverRuleCount *int32 `json:"resolver_rule_count,omitempty"`

	// 创建时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间。 格式：yyyy-MM-dd'T'HH:mm:ss.SSS。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o EndpointResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointResp struct{}"
	}

	return strings.Join([]string{"EndpointResp", string(data)}, " ")
}
