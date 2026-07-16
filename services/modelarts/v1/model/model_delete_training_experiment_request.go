package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrainingExperimentRequest Request Object
type DeleteTrainingExperimentRequest struct {

	// **参数解释**：实验ID。创建训练实验时自动生成返回。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ExperimentId string `json:"experiment_id"`
}

func (o DeleteTrainingExperimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrainingExperimentRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrainingExperimentRequest", string(data)}, " ")
}
