package model

import (
	"encoding/json"

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
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTransferTaskV3Request struct{}"
	}

	return strings.Join([]string{"DeleteTransferTaskV3Request", string(data)}, " ")
}
