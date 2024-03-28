package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnhancedConnectionRequest Request Object
type UpdateEnhancedConnectionRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Body *UpdateEnhancedConnectionRequestBody `json:"body,omitempty"`
}

func (o UpdateEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnhancedConnectionRequest", string(data)}, " ")
}
