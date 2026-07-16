package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingExperimentResponse Response Object
type CreateTrainingExperimentResponse struct {
	Metadata *TrainingExperimentResponseMetadata `json:"metadata,omitempty"`

	Statistic      *TrainingExperimentStatistic `json:"statistic,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o CreateTrainingExperimentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingExperimentResponse struct{}"
	}

	return strings.Join([]string{"CreateTrainingExperimentResponse", string(data)}, " ")
}
