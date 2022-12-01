package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RetryBatchJobRequest struct {

	// 批量处理作业ID
	JobId string `json:"job_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o RetryBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryBatchJobRequest struct{}"
	}

	return strings.Join([]string{"RetryBatchJobRequest", string(data)}, " ")
}
