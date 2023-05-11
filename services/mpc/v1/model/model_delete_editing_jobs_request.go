package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEditingJobsRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o DeleteEditingJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEditingJobsRequest struct{}"
	}

	return strings.Join([]string{"DeleteEditingJobsRequest", string(data)}, " ")
}
