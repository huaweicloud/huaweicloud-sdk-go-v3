package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRouteToEnhancedConnectionRequest Request Object
type CreateRouteToEnhancedConnectionRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Body *CreateRouteToEnhancedConnectionRequestBody `json:"body,omitempty"`
}

func (o CreateRouteToEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteToEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateRouteToEnhancedConnectionRequest", string(data)}, " ")
}
