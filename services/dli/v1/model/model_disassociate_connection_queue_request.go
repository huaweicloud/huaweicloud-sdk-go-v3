package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateConnectionQueueRequest Request Object
type DisassociateConnectionQueueRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Body *DisassociateConnectionQueueReq `json:"body,omitempty"`
}

func (o DisassociateConnectionQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateConnectionQueueRequest struct{}"
	}

	return strings.Join([]string{"DisassociateConnectionQueueRequest", string(data)}, " ")
}
