package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeSpecUpdateNodeNicSpecUpdatePrimaryNic **参数解释**： 主网卡的描述信息。 **约束限制**： 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type NodeSpecUpdateNodeNicSpecUpdatePrimaryNic struct {

	// **参数解释**： 网卡所在子网的网络ID。 **约束限制**： 主网卡创建时若未指定subnetId,将使用集群子网。若节点池同时配置了subnetList，则节点池扩容子网以subnetList字段为准。扩展网卡创建时必须指定subnetId。 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubnetId *string `json:"subnetId,omitempty"`

	// **参数解释**： 网卡所在子网的网络ID列表，支持节点池配置多个子网。 **约束限制**： 最多支持配置20个子网。 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubnetList *[]string `json:"subnetList,omitempty"`
}

func (o NodeSpecUpdateNodeNicSpecUpdatePrimaryNic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSpecUpdateNodeNicSpecUpdatePrimaryNic struct{}"
	}

	return strings.Join([]string{"NodeSpecUpdateNodeNicSpecUpdatePrimaryNic", string(data)}, " ")
}
