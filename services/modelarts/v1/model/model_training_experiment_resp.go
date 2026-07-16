package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrainingExperimentResp 训练实验参数。
type TrainingExperimentResp struct {

	// **参数解释**：实验名称。 **约束限制**：最大长度64，不支持特殊字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：实验ID。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：当前训练作业在所属的训练实验中的序号。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：默认为0。
	SerialNumber *string `json:"serial_number,omitempty"`
}

func (o TrainingExperimentResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingExperimentResp struct{}"
	}

	return strings.Join([]string{"TrainingExperimentResp", string(data)}, " ")
}
