package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartTransferTaskRequest Request Object
type BatchStartTransferTaskRequest struct {

	// 需要查询的通道名称。
	StreamName string `json:"stream_name"`

	Body *BatchStartTransferTaskReq `json:"body,omitempty"`
}

func (o BatchStartTransferTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartTransferTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchStartTransferTaskRequest", string(data)}, " ")
}
