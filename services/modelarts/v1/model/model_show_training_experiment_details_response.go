package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingExperimentDetailsResponse Response Object
type ShowTrainingExperimentDetailsResponse struct {
	Metadata *TrainingExperimentResponseMetadata `json:"metadata,omitempty"`

	Statistic      *TrainingExperimentStatistic `json:"statistic,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowTrainingExperimentDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingExperimentDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingExperimentDetailsResponse", string(data)}, " ")
}
