package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEditingJobsRequest struct {

	// 任务ID。一次最多10个
	JobId []string `json:"job_id"`
}

func (o ListEditingJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEditingJobsRequest struct{}"
	}

	return strings.Join([]string{"ListEditingJobsRequest", string(data)}, " ")
}
