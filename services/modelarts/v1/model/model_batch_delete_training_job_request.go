package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTrainingJobRequest Request Object
type BatchDeleteTrainingJobRequest struct {
	Body *BatchDeleteJobsReq `json:"body,omitempty"`
}

func (o BatchDeleteTrainingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTrainingJobRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTrainingJobRequest", string(data)}, " ")
}
