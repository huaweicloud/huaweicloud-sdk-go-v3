package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrainingExperimentRequestMetadata 创建训练实验的数据。
type TrainingExperimentRequestMetadata struct {

	// **参数解释**：实验名称。 **约束限制**：最大长度64，不支持特殊字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：描述信息。 **约束限制**：最大长度256，不支持特殊字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：工作空间ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：默认为0。
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o TrainingExperimentRequestMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingExperimentRequestMetadata struct{}"
	}

	return strings.Join([]string{"TrainingExperimentRequestMetadata", string(data)}, " ")
}
