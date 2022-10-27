package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RetryJobRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o RetryJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryJobRequest struct{}"
	}

	return strings.Join([]string{"RetryJobRequest", string(data)}, " ")
}
