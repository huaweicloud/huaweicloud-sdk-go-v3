package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubStage 作业流程阶段信息列表的子阶段元信息。
type SubStage struct {

	// **参数解释**：子阶段名称。  **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：子阶段英文描述信息。  **取值范围**：不涉及。
	EnMessage *string `json:"en_message,omitempty"`

	// **参数解释**：子阶段中文描述信息。  **取值范围**：不涉及。
	ZhMessage *string `json:"zh_message,omitempty"`

	// **参数解释**：子阶段开始时间。  **取值范围**：不涉及。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o SubStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubStage struct{}"
	}

	return strings.Join([]string{"SubStage", string(data)}, " ")
}
