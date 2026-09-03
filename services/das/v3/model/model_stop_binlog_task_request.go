package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopBinlogTaskRequest Request Object
type StopBinlogTaskRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *StopBinlogTaskRequestBody `json:"body,omitempty"`
}

func (o StopBinlogTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBinlogTaskRequest struct{}"
	}

	return strings.Join([]string{"StopBinlogTaskRequest", string(data)}, " ")
}
