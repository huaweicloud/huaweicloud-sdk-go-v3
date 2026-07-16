package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingExperimentDetailsRequest Request Object
type ShowTrainingExperimentDetailsRequest struct {

	// **参数解释**：实验ID。创建训练实验时自动生成返回。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ExperimentId string `json:"experiment_id"`
}

func (o ShowTrainingExperimentDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingExperimentDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainingExperimentDetailsRequest", string(data)}, " ")
}
