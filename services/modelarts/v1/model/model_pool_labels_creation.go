package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolLabelsCreation 资源池标签信息。
type PoolLabelsCreation struct {

	// **参数解释**：用户指定的资源池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsName string `json:"os.modelarts/name"`

	// **参数解释**：工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc) **约束限制**：不涉及。 **取值范围**：未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **默认取值**：不涉及。
	OsModelartsWorkspaceId *string `json:"os.modelarts/workspace.id,omitempty"`

	// **参数解释**：自定义节点名称前缀。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsNodePrefix *string `json:"os.modelarts/node.prefix,omitempty"`
}

func (o PoolLabelsCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolLabelsCreation struct{}"
	}

	return strings.Join([]string{"PoolLabelsCreation", string(data)}, " ")
}
