package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateBatchJobRequest struct {
	// 数据开发任务ID。

	JobId string `json:"job_id"`

	Body *Job `json:"body,omitempty"`
}

func (o UpdateBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBatchJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateBatchJobRequest", string(data)}, " ")
}
