package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeSpecUpdateNodeNicSpecUpdate **参数解释**： 更新节点的网卡信息。 **约束限制**： 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type NodeSpecUpdateNodeNicSpecUpdate struct {
	PrimaryNic *NodeSpecUpdateNodeNicSpecUpdatePrimaryNic `json:"primaryNic,omitempty"`
}

func (o NodeSpecUpdateNodeNicSpecUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSpecUpdateNodeNicSpecUpdate struct{}"
	}

	return strings.Join([]string{"NodeSpecUpdateNodeNicSpecUpdate", string(data)}, " ")
}
