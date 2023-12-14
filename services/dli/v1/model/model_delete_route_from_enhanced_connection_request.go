package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRouteFromEnhancedConnectionRequest Request Object
type DeleteRouteFromEnhancedConnectionRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Name string `json:"name"`
}

func (o DeleteRouteFromEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRouteFromEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteRouteFromEnhancedConnectionRequest", string(data)}, " ")
}
