package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePriority **参数解释：** 节点优先级批量配置，通过节点标签选择器匹配节点并为匹配的节点设置升级优先级。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type NodePriority struct {
	NodeSelector *NodeSelector `json:"nodeSelector"`

	// **参数解释：** 该批次节点的优先级，数值越大优先级越高 **约束限制：** 不涉及 **取值范围：** 非负整数 **默认取值：** 0
	Priority int32 `json:"priority"`
}

func (o NodePriority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePriority struct{}"
	}

	return strings.Join([]string{"NodePriority", string(data)}, " ")
}
