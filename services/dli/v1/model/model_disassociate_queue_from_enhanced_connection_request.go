package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateQueueFromEnhancedConnectionRequest Request Object
type DisassociateQueueFromEnhancedConnectionRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Body *DisassociateQueueFromEnhancedConnectionRequestBody `json:"body,omitempty"`
}

func (o DisassociateQueueFromEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateQueueFromEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"DisassociateQueueFromEnhancedConnectionRequest", string(data)}, " ")
}
