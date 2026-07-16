package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrainingExperimentResponseMetadata 实验的响应数据。
type TrainingExperimentResponseMetadata struct {

	// **参数解释**：实验名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：描述信息。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：工作空间ID。 **取值范围**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**：更新时间。 **取值范围**：不涉及。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**：实验ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`
}

func (o TrainingExperimentResponseMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingExperimentResponseMetadata struct{}"
	}

	return strings.Join([]string{"TrainingExperimentResponseMetadata", string(data)}, " ")
}
