package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingExperimentRequestBody 训练实验创建请求体。
type CreateTrainingExperimentRequestBody struct {
	Metadata *TrainingExperimentRequestMetadata `json:"metadata"`
}

func (o CreateTrainingExperimentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingExperimentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTrainingExperimentRequestBody", string(data)}, " ")
}
