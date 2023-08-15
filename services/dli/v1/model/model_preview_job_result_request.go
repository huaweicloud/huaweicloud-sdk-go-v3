package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreviewJobResultRequest Request Object
type PreviewJobResultRequest struct {

	// 作业ID。
	JobId string `json:"job_id"`

	// 指定获取作业结果的执行队列名称。若不指定则使用默认的系统队列
	QueueName *string `json:"queue-name,omitempty"`
}

func (o PreviewJobResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewJobResultRequest struct{}"
	}

	return strings.Join([]string{"PreviewJobResultRequest", string(data)}, " ")
}
