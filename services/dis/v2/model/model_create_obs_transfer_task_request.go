package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObsTransferTaskRequest Request Object
type CreateObsTransferTaskRequest struct {

	// 已创建的通道名称。
	StreamName string `json:"stream_name"`

	Body *CreateTransferTaskReq `json:"body,omitempty"`
}

func (o CreateObsTransferTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObsTransferTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateObsTransferTaskRequest", string(data)}, " ")
}
