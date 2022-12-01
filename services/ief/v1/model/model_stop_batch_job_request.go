package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopBatchJobRequest struct {

	// 批量处理作业ID
	JobId string `json:"job_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o StopBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBatchJobRequest struct{}"
	}

	return strings.Join([]string{"StopBatchJobRequest", string(data)}, " ")
}
