package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransferTaskRequest Request Object
type ShowTransferTaskRequest struct {

	// 已创建的通道的名称。
	StreamName string `json:"stream_name"`

	// 待删除的转储任务名称。
	TaskName string `json:"task_name"`
}

func (o ShowTransferTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransferTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowTransferTaskRequest", string(data)}, " ")
}
