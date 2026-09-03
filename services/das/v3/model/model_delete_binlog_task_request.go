package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBinlogTaskRequest Request Object
type DeleteBinlogTaskRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteBinlogTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBinlogTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteBinlogTaskRequest", string(data)}, " ")
}
