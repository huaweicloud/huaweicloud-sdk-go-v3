package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBinlogTaskRequest Request Object
type CreateBinlogTaskRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *CreateBinlogTaskRequestBody `json:"body,omitempty"`
}

func (o CreateBinlogTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBinlogTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateBinlogTaskRequest", string(data)}, " ")
}
