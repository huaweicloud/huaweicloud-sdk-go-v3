package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopTransferTaskRequest Request Object
type BatchStopTransferTaskRequest struct {

	// 需要查询的通道名称。
	StreamName string `json:"stream_name"`

	Body *BatchStopTransferTaskReq `json:"body,omitempty"`
}

func (o BatchStopTransferTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopTransferTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchStopTransferTaskRequest", string(data)}, " ")
}
