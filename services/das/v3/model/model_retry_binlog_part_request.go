package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryBinlogPartRequest Request Object
type RetryBinlogPartRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *RetryBinlogPartRequestBody `json:"body,omitempty"`
}

func (o RetryBinlogPartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryBinlogPartRequest struct{}"
	}

	return strings.Join([]string{"RetryBinlogPartRequest", string(data)}, " ")
}
