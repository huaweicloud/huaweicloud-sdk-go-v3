package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateNodesToNodePoolList 节点迁移到自定义节点池参数。
type MigrateNodesToNodePoolList struct {

	// **参数解释**： API版本 **约束限制**： 固定值，不允许修改。 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion string `json:"apiVersion"`

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改。 **取值范围**： 不涉及 **默认取值**： List
	Kind string `json:"kind"`

	// **参数解释**： 迁移节点列表，当前最多支持同迁移50个节点。 **约束限制**： 不涉及
	NodeList []MigrateNodesToNodePool `json:"nodeList"`
}

func (o MigrateNodesToNodePoolList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodesToNodePoolList struct{}"
	}

	return strings.Join([]string{"MigrateNodesToNodePoolList", string(data)}, " ")
}
