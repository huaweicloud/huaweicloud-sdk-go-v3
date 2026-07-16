package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainingJobRequest Request Object
type CreateTrainingJobRequest struct {
	Body *Job `json:"body,omitempty"`
}

func (o CreateTrainingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainingJobRequest struct{}"
	}

	return strings.Join([]string{"CreateTrainingJobRequest", string(data)}, " ")
}
