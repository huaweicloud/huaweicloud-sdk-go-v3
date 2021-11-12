package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteTransferTaskV3Request struct {
	// 已创建的通道的名称。

	StreamName string `json:"stream_name"`
	// 待删除的转储任务名称。

	TaskName string `json:"task_name"`
}

func (o DeleteTransferTaskV3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransferTaskV3Request struct{}"
	}

	return strings.Join([]string{"DeleteTransferTaskV3Request", string(data)}, " ")
}
