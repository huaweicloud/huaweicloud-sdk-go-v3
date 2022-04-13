package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteFailureJobRequest struct {
	// 失败任务ID。

	FailureJobId string `json:"failure_job_id"`
}

func (o DeleteFailureJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFailureJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteFailureJobRequest", string(data)}, " ")
}
