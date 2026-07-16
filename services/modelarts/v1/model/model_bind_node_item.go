package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindNodeItem struct {

	// **参数解释**：换绑的节点的名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：节点绑定的逻辑子池的ID。值为空则节点不绑定任何逻辑子池。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	QuotaName *string `json:"quotaName,omitempty"`
}

func (o BindNodeItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindNodeItem struct{}"
	}

	return strings.Join([]string{"BindNodeItem", string(data)}, " ")
}
