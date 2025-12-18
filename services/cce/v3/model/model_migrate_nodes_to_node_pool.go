package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateNodesToNodePool 节点迁移到自定义节点池参数。
type MigrateNodesToNodePool struct {

	// **参数解释**： 节点ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	NodeID string `json:"nodeID"`
}

func (o MigrateNodesToNodePool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateNodesToNodePool struct{}"
	}

	return strings.Join([]string{"MigrateNodesToNodePool", string(data)}, " ")
}
