package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreviewSqlJobResultRequest Request Object
type PreviewSqlJobResultRequest struct {

	// 作业ID。
	JobId string `json:"job_id"`

	// 指定获取作业结果的执行队列名称。若不指定则使用默认的系统队列
	QueueName *string `json:"queue-name,omitempty"`
}

func (o PreviewSqlJobResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewSqlJobResultRequest struct{}"
	}

	return strings.Join([]string{"PreviewSqlJobResultRequest", string(data)}, " ")
}
