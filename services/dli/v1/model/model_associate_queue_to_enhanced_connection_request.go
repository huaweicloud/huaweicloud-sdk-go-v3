package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateQueueToEnhancedConnectionRequest Request Object
type AssociateQueueToEnhancedConnectionRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Body *AssociateQueueToEnhancedConnectionRequestBody `json:"body,omitempty"`
}

func (o AssociateQueueToEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateQueueToEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"AssociateQueueToEnhancedConnectionRequest", string(data)}, " ")
}
