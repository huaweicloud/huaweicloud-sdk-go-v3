package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnhancedConnectionRequest Request Object
type DeleteEnhancedConnectionRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`
}

func (o DeleteEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteEnhancedConnectionRequest", string(data)}, " ")
}
