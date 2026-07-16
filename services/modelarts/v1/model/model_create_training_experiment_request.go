package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingExperimentRequest Request Object
type CreateTrainingExperimentRequest struct {
	Body *CreateTrainingExperimentRequestBody `json:"body,omitempty"`
}

func (o CreateTrainingExperimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingExperimentRequest struct{}"
	}

	return strings.Join([]string{"CreateTrainingExperimentRequest", string(data)}, " ")
}
