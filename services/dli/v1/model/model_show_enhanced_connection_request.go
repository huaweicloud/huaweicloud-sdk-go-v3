package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnhancedConnectionRequest Request Object
type ShowEnhancedConnectionRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`
}

func (o ShowEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowEnhancedConnectionRequest", string(data)}, " ")
}
