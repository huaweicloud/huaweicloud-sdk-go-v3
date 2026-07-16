package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrainingExperimentResponse 训练实验。
type TrainingExperimentResponse struct {
	Metadata *TrainingExperimentResponseMetadata `json:"metadata,omitempty"`

	Statistic *TrainingExperimentStatistic `json:"statistic,omitempty"`
}

func (o TrainingExperimentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingExperimentResponse struct{}"
	}

	return strings.Join([]string{"TrainingExperimentResponse", string(data)}, " ")
}
