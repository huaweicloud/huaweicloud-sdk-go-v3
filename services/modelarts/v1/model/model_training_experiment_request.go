package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrainingExperimentRequest 训练实验参数。
type TrainingExperimentRequest struct {

	// **参数解释**：实验ID，填写实验ID时，此训练作业将会纳入该已有实验分组，填写前请确保该实验ID真实存在。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：实验名称，只填写实验名称时，将会创建该实验分组，并将此训练作业纳入该分组。 **约束限制**：最大长度64，不支持特殊字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：描述信息。 **约束限制**：最大长度256，不支持特殊字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`
}

func (o TrainingExperimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingExperimentRequest struct{}"
	}

	return strings.Join([]string{"TrainingExperimentRequest", string(data)}, " ")
}
