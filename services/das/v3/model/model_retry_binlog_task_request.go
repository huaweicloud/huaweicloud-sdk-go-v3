package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryBinlogTaskRequest Request Object
type RetryBinlogTaskRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 任务ID
	TaskId int64 `json:"task_id"`
}

func (o RetryBinlogTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryBinlogTaskRequest struct{}"
	}

	return strings.Join([]string{"RetryBinlogTaskRequest", string(data)}, " ")
}
