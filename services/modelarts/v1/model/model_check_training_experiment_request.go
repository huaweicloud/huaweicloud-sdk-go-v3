package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTrainingExperimentRequest Request Object
type CheckTrainingExperimentRequest struct {

	// **参数解释**：工作空间ID。获取方法请参见[查询工作空间列表](ListWorkspace.xml)。 **约束限制**：存在并使用的工作空间。 **取值范围**：不涉及。 **默认取值**：“0”。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：实验名称。 **约束限制**：最大长度64，不支持特殊字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ExperimentName string `json:"experiment_name"`
}

func (o CheckTrainingExperimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTrainingExperimentRequest struct{}"
	}

	return strings.Join([]string{"CheckTrainingExperimentRequest", string(data)}, " ")
}
