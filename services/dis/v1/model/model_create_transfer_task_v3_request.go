package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTransferTaskV3Request struct {
	// 已创建的通道名称。

	StreamName string `json:"stream_name"`

	Body *CreateTransferTaskRequest `json:"body,omitempty"`
}

func (o CreateTransferTaskV3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferTaskV3Request struct{}"
	}

	return strings.Join([]string{"CreateTransferTaskV3Request", string(data)}, " ")
}
