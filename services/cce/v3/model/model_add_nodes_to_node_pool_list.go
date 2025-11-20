package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddNodesToNodePoolList 自定义节点池纳管节点参数，纳管过程将清理节点上系统盘、数据盘数据，并作为新节点接入Kuberntes集群，请提前备份迁移关键数据。
type AddNodesToNodePoolList struct {

	// **参数解释**： API版本 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion string `json:"apiVersion"`

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： List
	Kind string `json:"kind"`

	// **参数解释**： 纳管节点列表，当前最多支持同时纳管200个节点。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeList []AddNodesToNodePool `json:"nodeList"`
}

func (o AddNodesToNodePoolList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNodesToNodePoolList struct{}"
	}

	return strings.Join([]string{"AddNodesToNodePoolList", string(data)}, " ")
}
