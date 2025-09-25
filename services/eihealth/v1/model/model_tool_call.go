package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ToolCall struct {

	// **参数解释**： 工具类型。 **约束限制**： 不涉及 **取值范围**： * drug_job：药物作业 * workflow：流程作业 **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 工具名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 工具操作。 **约束限制**： 不涉及 **取值范围**： * create：新增 * delete：删除 * update：修改 * query：查询 **默认取值**： 不涉及
	Operation *string `json:"operation,omitempty"`

	// **参数解释**： 工具调用所需参数列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

func (o ToolCall) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ToolCall struct{}"
	}

	return strings.Join([]string{"ToolCall", string(data)}, " ")
}
