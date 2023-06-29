package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreBatchJobRequest Request Object
type RestoreBatchJobRequest struct {

	// 批量处理作业ID
	JobId string `json:"job_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o RestoreBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreBatchJobRequest struct{}"
	}

	return strings.Join([]string{"RestoreBatchJobRequest", string(data)}, " ")
}
