package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBinlogTaskInfoRequest Request Object
type ShowBinlogTaskInfoRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 任务ID
	TaskId int64 `json:"task_id"`
}

func (o ShowBinlogTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowBinlogTaskInfoRequest", string(data)}, " ")
}
