package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateConnectionQueueRequest Request Object
type AssociateConnectionQueueRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Body *AssociateConnectionQueueReq `json:"body,omitempty"`
}

func (o AssociateConnectionQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateConnectionQueueRequest struct{}"
	}

	return strings.Join([]string{"AssociateConnectionQueueRequest", string(data)}, " ")
}
